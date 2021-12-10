package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dragonfax/delver_converter/parser"
)

const sourceDir = "../delverengine/Dungeoneer/src/com/interrupt/"

var lexer *parser.JavaLexer
var p *parser.JavaParser

func main() {

	lexer = parser.NewJavaLexer(nil)
	p = parser.NewJavaParser(nil)

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

func parse(path string) {
	input, _ := antlr.NewFileStream(path)
	lexer.SetInputStream(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p.SetInputStream(stream)

	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	p.BuildParseTrees = true
	listener := &listener{Filename: path}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.CompilationUnit())

	js, err := json.MarshalIndent(listener, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(js))
}
