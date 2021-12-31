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

func Translate(sourceDir, targetDir, targetPackage, targetStubDir string) error {

	info, err := os.Stat(sourceDir)
	if err != nil {
		panic(err)
	}

	// parsing files

	fmt.Println("parsing")
	dir := sourceDir
	var classes []*ast.Class
	if info.IsDir() {
		filenames := gatherFilenames(sourceDir)
		classes = parseFiles(filenames)
	} else {
		dir = filepath.Dir(sourceDir)
		class := parseFile(sourceDir)
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
	visitor.RunGroup(h)

	// output
	fmt.Println("writing files")
	if targetDir == "" {
		targetDir = generateOutputRoot(dir)
	}
	h.RootGoPackage = targetPackage
	outputStructures(h.GetClasses(), targetDir)
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

	targetFilename := outputRoot + "/" + class.Filename()
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
