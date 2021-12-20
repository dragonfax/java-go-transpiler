package input

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dragonfax/java_converter/input/parser"
)

var lexer = parser.NewJavaLexer(nil)
var p = parser.NewJavaParser(nil)

func ParseToTree(filename string) antlr.RuleContext {

	input, _ := antlr.NewFileStream(filename)
	lexer.SetInputStream(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p.SetInputStream(stream)
	p.BuildParseTrees = true

	return p.CompilationUnit()
}
