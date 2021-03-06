package output

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/dragonfax/java_converter/output/trans"
)

func walkFunc(filename string, entry fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if entry.IsDir() || !strings.HasSuffix(filename, ".java") {
		return nil
	}

	TranslateOneFile(filename)
	return nil
}

func RemoveFileIfExists(filename string) {
	// remove target file if its already there.
	_, err := os.Stat(filename)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}
	} else {
		// file exists. thats a problem.
		err = os.Remove(filename)
		if err != nil {
			panic(err)
		}
	}

}

func TranslateOneFile(filename string) {
	targetFilename := GenerateTargetFilename(filename)

	RemoveFileIfExists(targetFilename)
	RemoveFileIfExists(targetFilename + ".json")
	RemoveFileIfExists(targetFilename + ".err.txt")

	targetDirectory := filepath.Dir(targetFilename)
	err := os.MkdirAll(targetDirectory, 0775)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s -> %s\n", filename, targetFilename)
	trans.TranslateFiles(filename, targetFilename)

}

func CrawlDir(path string) error {
	return filepath.WalkDir(path, walkFunc)
}

func GenerateTargetFilename(filename string) string {

	// change suffix
	baseFilename := strings.TrimSuffix(filename, ".java")
	goFilename := baseFilename + ".go"

	//take first real component, and change its name.
	components := strings.Split(goFilename, "/")
	for i, component := range components {
		if component != "." && component != ".." {
			components[i] = components[i] + "_converted"
			break
		}
	}
	targetFilename := strings.Join(components, "/")

	if filename == targetFilename {
		panic("didn't generate a new filename")
	}

	return targetFilename
}
