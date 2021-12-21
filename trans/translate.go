package trans

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dragonfax/java_converter/trans/hier"
	"github.com/dragonfax/java_converter/trans/node"
	"github.com/dragonfax/java_converter/trans/visitor"
)

func BuildAST(h *hier.Hierarchy, tree antlr.RuleContext) node.Node {

	TreeVisitor := visitor.NewTreeVisitor(h)

	file := TreeVisitor.Visit(tree)

	return file
}
