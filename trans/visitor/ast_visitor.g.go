package visitor

import (
	"github.com/dragonfax/java_converter/trans/ast"
)

type GenASTVisitor[T comparable] interface {
	VisitField(tree *ast.Field) T
	VisitPackage(tree *ast.Package) T
	VisitClass(tree *ast.Class) T
	VisitImport(tree *ast.Import) T
	VisitMember(tree *ast.Member) T
	VisitVarRef(ctx *ast.VarRef) T
	VisitLocalVarDecl(ctx *ast.LocalVarDecl) T
}
