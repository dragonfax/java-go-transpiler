package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dragonfax/java_converter/parser"
)

func main() {

	input, err := antlr.NewFileStream("DrawableMesh.java")
	if err != nil {
		panic(err)
	}
	lexer := parser.NewJavaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJavaParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.CompilationUnit()

	structVisitor := &StructVisitor{}
	visitor := NewJavaVisitor(structVisitor)
	ast := visitor.Visit(tree)

	fmt.Println(ast)
}
