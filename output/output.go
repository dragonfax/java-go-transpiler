package output

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/dragonfax/java_converter/input"
	"github.com/dragonfax/java_converter/trans"
	"github.com/dragonfax/java_converter/trans/ast"
	"github.com/dragonfax/java_converter/trans/node"
	"github.com/dragonfax/java_converter/trans/visitor"
	"github.com/schollz/progressbar/v3"
)

func walkFunc(filenames *[]string) fs.WalkDirFunc {
	return func(filename string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() || !strings.HasSuffix(filename, ".java") {
			return nil
		}

		*filenames = append(*filenames, filename)

		return nil
	}
}

func gatherFilenames(path string) []string {
	filenames := make([]string, 0)
	filepath.WalkDir(path, walkFunc(&filenames))
	return filenames
}

func parseFiles(filenames []string) []*ast.Class {

	bar := progressbar.Default(100)

	classes := make([]*ast.Class, 0)
	for _, filename := range filenames {
		class := parseFile(filename)
		if class != nil {
			classes = append(classes, class)
		}
		bar.Add(1)
	}
	return classes
}

// input filename to go-code string
func parseFile(filename string) *ast.Class {
	// fmt.Println(filename)
	tree := input.ParseToTree(filename)

	return trans.BuildAST(tree)
}

func Translate(path string) error {

	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	// parsing files

	fmt.Println("parsing")
	dir := path
	var classes []*ast.Class
	if info.IsDir() {
		filenames := gatherFilenames(path)
		classes = parseFiles(filenames)
	} else {
		dir = filepath.Dir(path)
		class := parseFile(path)
		classes = []*ast.Class{class}
	}
	if err != nil {
		return err
	}
	fmt.Println("parsing complete")

	// process the global AST

	// build the root hierary
	fmt.Println("ast walking")
	h := ast.NewHierarchy()
	for _, class := range classes {
		h.AddClass(class)
	}
	// process the hierarchy (all classes and packages) at once
	fmt.Println("scope walk")
	scopeVisitor := visitor.NewScopeVisitor(h)
	scopeVisitor.VisitNode(h)
	//fmt.Println("resolve walk")
	//resolver := visitor.NewResolverVisitor(h)
	//resolver.VisitNode(h)
	fmt.Println("ast walking complete")

	// output
	fmt.Println("writing files")
	outputRoot := generateOutputRoot(dir)
	outputStructures(h.GetClasses(), outputRoot)
	fmt.Println("writing files complete")

	return nil
}

func outputStructures(classes []*ast.Class, outputRoot string) {

	bar := progressbar.Default(int64(len(classes)))

	for _, class := range classes {
		outputStructure(class, outputRoot)
		bar.Add(1)
	}
}

func outputStructure(class *ast.Class, outputRoot string) {

	targetFilename := outputRoot + "/" + class.OutputFilename()
	targetErrorFilename := targetFilename + ".err.txt"

	targetDir := filepath.Dir(targetFilename)
	err := os.MkdirAll(targetDir, 0775)

	jsonFilename := targetFilename + ".json"
	removeFileIfExists(jsonFilename)
	outputFile(jsonFilename, node.JSONMarshalNode(class))

	removeFileIfExists(targetFilename)
	removeFileIfExists(targetErrorFilename)

	goCode := class.AsFile()

	goCode2, err := goFmt(goCode)
	if err != nil {
		outputFile(targetErrorFilename, goCode2)
	} else {
		goCode = goCode2
	}

	outputFile(targetFilename, goCode)
}

func generateOutputRoot(dir string) string {
	/* take the first path thats not ../
	 * and add a "_converted" to it
	 * then add the origin rest of the path back on.
	 * so the new file structure starts at the base of the old one.
	 */

	//take first real component, and change its name.
	components := strings.Split(dir, "/")
	for i, component := range components {
		if component != "." && component != ".." {
			components[i] = components[i] + "_converted"
			break
		}
	}
	targetDir := strings.Join(components, "/")

	if dir == targetDir {
		panic("didn't generate a new filename")
	}

	return targetDir
}
