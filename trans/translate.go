package trans

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dragonfax/java-go-transpiler/trans/ast"
	"github.com/dragonfax/java-go-transpiler/trans/visitor"
)

func BuildAST(tree antlr.RuleContext) *ast.Class {

	TreeVisitor := visitor.NewTreeVisitor()

	class := TreeVisitor.Visit(tree)

	if class == nil {
		return nil
	}

	return class.(*ast.Class)
}
