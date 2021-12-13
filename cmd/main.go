package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/dragonfax/delver_converter/ast"

	"github.com/aymerick/raymond"
)

const sourceDir = "../delverengine/Dungeoneer/src/com/interrupt/"

var targetPath = os.Args[1]

func main() {

	if len(os.Args) > 2 {
		parse(os.Args[2])
	} else {

		walkFunc := func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".java") {
				parse(path)
			}
			return nil
		}
		err := filepath.WalkDir(sourceDir, walkFunc)
		if err != nil {
			panic(err)
		}
	}

}

func snakeCase(s string) string {
	if s == "" {
		panic("no filename given")
	}
	s = strings.TrimPrefix(s, "../")
	s = strings.TrimSuffix(s, ".java")
	return strings.ReplaceAll(s, "/", "_")
}

func outputFile(file *ast.File) {
	result, err := raymond.Render(string(golangTemplate), file)
	if err != nil {
		panic(err)
	}

	snakeFilename := snakeCase(file.Filename)
	fmt.Printf("writing '%s'\n", snakeFilename)
	err = ioutil.WriteFile(targetPath+"/"+snakeFilename+".go", []byte(result), 0664)
	if err != nil {
		panic(err)
	}
}
