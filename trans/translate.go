package trans

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dragonfax/java_converter/trans/ast"
	"github.com/dragonfax/java_converter/trans/visitor"
)

func BuildAST(tree antlr.RuleContext) *ast.File {

	goVisitor := visitor.NewGoVisitor()

	file := goVisitor.Visit(tree)

	return file.(*ast.File)
}
