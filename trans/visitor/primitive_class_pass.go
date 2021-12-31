package visitor

import (
	"github.com/dragonfax/java_converter/trans/ast"
)

/* Tell primitives what their runtime class is */
type PrimitiveClassPass struct {
	*BaseASTVisitor[int] // throwaway return value
}

func NewPrimitiveClassPass(h *ast.Hierarchy) *PrimitiveClassPass {
	this := &PrimitiveClassPass{}
	this.BaseASTVisitor = NewASTVisitor[int](h, this)

	return this
}

func (av *PrimitiveClassPass) VisitLiteral(literal *ast.Literal) int {

	literal.Type = av.Hierarchy.GetPackage("runtime").GetClass(string(literal.LiteralType))

	// have no children
	return 0
}
