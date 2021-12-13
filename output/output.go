package output

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/dragonfax/delver_converter/output/trans"
)

func snakeCase(s string) string {
	if s == "" {
		panic("no filename given")
	}
	s = strings.TrimPrefix(s, "../")
	s = strings.TrimSuffix(s, ".java")
	return strings.ReplaceAll(s, "/", "_")
}

func walkFunc(filename string, entry fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if entry.IsDir() || !strings.HasSuffix(filename, ".java") {
		return nil
	}

	/* this isn't good enough, we might have ../
	write the go files to the same place?
	no that'll get ugly.
	*/
	targetFilename := generateTargetFilename(filename)

	// remove target file if its already there.
	_, err = os.Stat(targetFilename)
	if err == nil {
		err = os.Remove(targetFilename)
		if err != nil {
			panic(err)
		}
	} else if err != os.ErrNotExist {
		panic(err)
	}

	fmt.Printf("%s -> %s\n", filename, targetFilename)
	return trans.TranslateFiles(filename, targetFilename)
}

func CrawlDir(path string) error {
	return filepath.WalkDir(path, walkFunc)
}

func generateTargetFilename(filename string) string {

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
