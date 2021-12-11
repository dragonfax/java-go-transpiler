package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/aymerick/raymond"
	"github.com/dragonfax/delver_converter/parser"
)

const sourceDir = "../delverengine/Dungeoneer/src/com/interrupt/"
const golangTemplateFilename = "golang.tmpl"

var lexer = parser.NewJavaLexer(nil)
var p = parser.NewJavaParser(nil)

var golangTemplate = mustByteListErr(ioutil.ReadFile(golangTemplateFilename))
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

func outputFile(file *File) {
	result, err := raymond.Render(string(golangTemplate), file)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(targetPath+"/"+file.QualifiedPackageName+".go", []byte(result), 0664)
	if err != nil {
		panic(err)
	}
}

func parse(path string) {

	fmt.Println(path)

	stackListener := NewStackListener()

	fileListener := NewFileListener(stackListener, path)
	stackListener.Push(fileListener)

	input, _ := antlr.NewFileStream(path)
	lexer.SetInputStream(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p.SetInputStream(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(stackListener, p.CompilationUnit())

	if stackListener.Len() != 1 {
		if stackListener.Len() == 2 {
			panic(fmt.Sprintf("left over listener %T", stackListener.stack[1]))
		}
		panic("wrong number of listeners left over, " + fmt.Sprintf("%d", stackListener.Len()))
	}

	/*
		js, err := json.MarshalIndent(p.CompilationUnit(), "", "  ")
		if err != nil {
			panic(err)
		}
	*/
	// fmt.Println(p.CompilationUnit().ToStringTree())

	js, err := json.MarshalIndent(fileListener.File, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(js))

	outputFile(fileListener.File)
}
