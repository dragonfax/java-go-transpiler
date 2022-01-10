package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dragonfax/java-go-transpiler/output"
)

func main() {

	var sourceDir string
	var targetDir string
	var targetPackage string
	var targetStubDir string

	flag.StringVar(&sourceDir, "source", "", "directory for java source, required")
	flag.StringVar(&targetDir, "target", "", "directory for go source, optional")
	flag.StringVar(&targetPackage, "package", "", "full go url of base package to generate, optional")
	flag.StringVar(&targetStubDir, "stubs", "", "directory for generated stub classes")

	flag.Parse()

	if sourceDir == "" {
		fmt.Println("missing source")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if targetStubDir == "" {
		targetStubDir = targetDir
	}

	err := output.Translate(sourceDir, targetDir, targetPackage, targetStubDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
