package visitor

import (
	"github.com/dragonfax/java_converter/trans/ast"
	"github.com/dragonfax/java_converter/trans/node"
)

type ASTVisitor[T comparable] interface {
	VisitNode(tree node.Node) T
	VisitChildren(tree node.Node) T
	VisitField(tree *ast.Field) T
	VisitConstructor(tree *ast.Constructor) T
	VisitPackage(tree *ast.Package) T
	VisitClass(tree *ast.Class) T
	VisitImport(tree *ast.Import) T
	VisitMethod(tree *ast.Method) T
	AggregateResult(result, nextResult T) T
}
