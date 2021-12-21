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
	"github.com/dragonfax/java_converter/trans/hier"
	"github.com/dragonfax/java_converter/trans/node"
)

func walkFunc(h *hier.Hierarchy) fs.WalkDirFunc {
	return func(filename string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() || !strings.HasSuffix(filename, ".java") {
			return nil
		}

		parseFile(h, filename)
		return nil
	}
}

func Translate(path string) error {

	h := hier.New()

	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	dir := path
	if info.IsDir() {
		err = filepath.WalkDir(path, walkFunc(h))
	} else {
		dir = filepath.Dir(path)
		parseFile(h, path)
	}
	if err != nil {
		return err
	}

	// process the global AST

	outputRoot := generateOutputRoot(dir)

	outputStructures(h, outputRoot)

	return nil
}

// input filename to go-code string
func parseFile(h *hier.Hierarchy, filename string) {
	fmt.Println(filename)
	tree := input.ParseToTree(filename)

	trans.BuildAST(h, tree)
}

func outputStructures(h *hier.Hierarchy, outputRoot string) {

	for _, class := range h.Classes {
		outputStructure(class, outputRoot)
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
