package visitor

import "github.com/dragonfax/java_converter/trans/ast"

type FetchTypesPass struct {
	*BaseASTVisitor[int] // throwaway return value
}

func NewFetchTypesPass(h *ast.Hierarchy) ASTVisitor[int] {
	this := &FetchTypesPass{}
	this.BaseASTVisitor = NewASTVisitor[int](h, this)

	return this
}

func (ft *FetchTypesPass) VisitField(field *ast.Field) int {

	field.GetType()

	return ft.VisitChildren(field)
}
