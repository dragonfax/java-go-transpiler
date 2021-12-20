package trans

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dragonfax/java_converter/trans/ast"
	"github.com/dragonfax/java_converter/trans/hier"
	"github.com/dragonfax/java_converter/trans/visitor"
)

func BuildAST(h *hier.Hierarchy, tree antlr.RuleContext) ast.Node {

	TreeVisitor := visitor.NewTreeVisitor(h)

	file := TreeVisitor.Visit(tree)

	return file
}
