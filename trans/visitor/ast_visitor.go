package visitor

import (
	"github.com/dragonfax/java_converter/trans/ast"
	"github.com/dragonfax/java_converter/trans/node"
)

type ASTVisitor[T comparable] interface {
	// pattern methods
	VisitNode(tree node.Node) T
	VisitChildren(tree node.Node) T
	AggregateResult(result, nextResult T) T

	// specific nodes
	VisitField(tree *ast.Field) T
	VisitPackage(tree *ast.Package) T
	VisitClass(tree *ast.Class) T
	VisitImport(tree *ast.Import) T
	VisitMember(tree *ast.Member) T
	VisitVarRef(ctx *ast.VarRef) T
}
