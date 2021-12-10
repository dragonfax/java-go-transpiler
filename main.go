package main

import (
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dragonfax/delver_converter/parser"
)

type listener struct {
	*parser.BaseJavaParserListener
}

func main() {

	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewJavaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewJavaParser(stream)

	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	p.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(&listener{}, p.CompilationUnit())

}
