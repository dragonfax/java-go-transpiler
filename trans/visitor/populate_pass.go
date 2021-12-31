package visitor

import (
	"github.com/dragonfax/java_converter/trans/ast"
)

/* populate some maps and lists in nodes.
 * list of var's ina method, list of fields in a class
 */
type PopulatePass struct {
	*BaseASTVisitor[int] // throwaway return value
}

func NewPopulatePass(h *ast.Hierarchy) ASTVisitor[int] {
	this := &PopulatePass{}
	this.BaseASTVisitor = NewASTVisitor[int](h, this)
	return this
}

func (cv *PopulatePass) VisitLocalVarDecl(localVarDecl *ast.LocalVarDecl) int {

	if localVarDecl.MethodScope == nil {
		ast.DebugPrint(localVarDecl)
		panic("local var never got a method")
	}

	// this ends up covering both local var declarations and formal parameters in methods/constructors
	localVarDecl.MethodScope.AddLocalVar(localVarDecl)

	return 0 // no children that need processesing.
}

func (cv *PopulatePass) VisitField(field *ast.Field) int {

	field.ClassScope.AddField(field)

	return 0 // no children that need processesing.
}
