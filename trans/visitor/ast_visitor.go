package visitor

import (
	"github.com/dragonfax/java_converter/trans/node"
)

type ASTVisitor[T comparable] interface {
	GenASTVisitor[T]

	// pattern methods
	VisitNode(tree node.Node) T
	VisitChildren(tree node.Node) T
	AggregateResult(result, nextResult T) T
}
