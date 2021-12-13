package main

import (
	"os"

	"github.com/dragonfax/delver_converter/output"
)

func main() {

	source := os.Args[1]

	info, err := os.Stat(source)
	if err != nil {
		panic(err)
	}

	if info.IsDir() {
		err = output.CrawlDir(source)
	} else {
		err = output.ParseFilename(source)
	}
	if err != nil {
		panic(err)
	}
}
