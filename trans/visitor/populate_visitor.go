package visitor

import (
	"github.com/dragonfax/java_converter/trans/ast"
)

/* populate some maps and lists in nodes. */

type PopulateVisitor struct {
	*BaseASTVisitor[int] // throwaway return value
}

func NewPopulateVisitor(h *ast.Hierarchy) *PopulateVisitor {
	this := &PopulateVisitor{}
	this.BaseASTVisitor = NewASTVisitor[int](h, this)
	return this
}

func (cv *PopulateVisitor) VisitLocalVarDecl(localVarDecl *ast.LocalVarDecl) int {

	// this ends up covering both local var declarations and formal parameters in methods/constructors
	localVarDecl.MemberScope.AddLocalVar(localVarDecl)

	return 0 // no children that need processesing.
}

func (cv *PopulateVisitor) VisitField(field *ast.Field) int {

	field.ClassScope.AddField(field)

	return 0 // no children that need processesing.
}
