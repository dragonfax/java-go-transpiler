package trans

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dragonfax/delver_converter/input/parser"
	"github.com/dragonfax/delver_converter/listen"
	"github.com/dragonfax/delver_converter/tool"
)

const golangTemplateFilename = "golang.tmpl"

var lexer = parser.NewJavaLexer(nil)
var p = parser.NewJavaParser(nil)

var golangTemplate = tool.MustByteListErr(ioutil.ReadFile(golangTemplateFilename))

func Translate(javaCode string) (goCode string) {

	fmt.Println(path)

	stackListener := listen.NewStackListener()

	fileListener := listen.NewFileListener(stackListener, path)
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
