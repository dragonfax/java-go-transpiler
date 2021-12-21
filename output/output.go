package output

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/dragonfax/java_converter/input"
	"github.com/dragonfax/java_converter/trans"
	"github.com/dragonfax/java_converter/trans/ast"
	"github.com/dragonfax/java_converter/trans/node"
)

func walkFunc(classes *[]node.Node) fs.WalkDirFunc {
	*classes = make([]node.Node, 0)
	return func(filename string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() || !strings.HasSuffix(filename, ".java") {
			return nil
		}

		class := parseFile(filename)
		*classes = append(*classes, class)
		return nil
	}
}

func Translate(path string) error {

	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	dir := path
	var classes []node.Node
	if info.IsDir() {
		err = filepath.WalkDir(path, walkFunc(&classes))
	} else {
		dir = filepath.Dir(path)
		class := parseFile(path)
		classes = []node.Node{class}
	}
	if err != nil {
		return err
	}

	// process the global AST

	outputRoot := generateOutputRoot(dir)

	outputStructures(classes, outputRoot)

	return nil
}

// input filename to go-code string
func parseFile(filename string) node.Node {
	fmt.Println(filename)
	tree := input.ParseToTree(filename)

	return trans.BuildAST(tree)
}

func outputStructures(classes []node.Node, outputRoot string) {

	for _, class := range classes {
		outputStructure(class.(*ast.Class), outputRoot)
	}
}

func outputStructure(class *ast.Class, outputRoot string) {

	targetFilename := outputRoot + "/" + class.OutputFilename()
	targetJSONFilename := targetFilename + ".json"
	targetErrorFilename := targetFilename + ".err.txt"

	targetDir := filepath.Dir(targetFilename)
	err := os.MkdirAll(targetDir, 0775)

	removeFileIfExists(targetFilename)
	removeFileIfExists(targetJSONFilename)
	removeFileIfExists(targetErrorFilename)

	goCode := class.AsFile()

	goCode2, err := goFmt(goCode)
	if err != nil {
		outputFile(targetErrorFilename, goCode2)
	} else {
		goCode = goCode2
	}

	astJSON, err := dumpAST(class)
	if err != nil {
		outputFile(targetJSONFilename, astJSON)
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

func dumpAST(node node.Node) (string, error) {
	js, err := json.MarshalIndent(node, "", "  ")
	return string(js), err
}
