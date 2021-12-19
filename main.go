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
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	tree := p.CompilationUnit()
	// fmt.Println(tree.ToStringTree(parser.RuleNames, nil))

	structVisitor := NewStructVisitor()

	ast := structVisitor.Visit(tree)
	fmt.Println(ast)
}
