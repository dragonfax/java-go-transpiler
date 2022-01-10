package visitor

import (
	"github.com/dragonfax/java-go-transpiler/trans/node"
)

type ASTVisitor[T comparable] interface {
	GenASTVisitor[T]

	// pattern methods
	VisitNode(tree node.Node) T
	VisitChildren(tree node.Node) T
	AggregateResult(result, nextResult T) T
}
