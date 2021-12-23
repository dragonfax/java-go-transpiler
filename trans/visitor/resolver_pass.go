package visitor

import (
	"fmt"

	"github.com/dragonfax/java_converter/trans/ast"
)

type ResolvePass struct {
	*BaseASTVisitor[int] // throwaway return value
}

func NewResolvePass(h *ast.Hierarchy) ASTVisitor[int] {
	this := &ResolvePass{}
	this.BaseASTVisitor = NewASTVisitor[int](h, this)
	return this
}

func (cv *ResolvePass) VisitVarRef(varRef *ast.VarRef) int {

	if varRef.Super {
		return 0
	}

	if varRef.This {
		return 0
	}

	name := varRef.VariableName

	// check in the method arguments and other method variables
	if varRef.MemberScope == nil {
		fmt.Println("no member scope for var ref")
		return 0
	}

	localVar, ok := varRef.MemberScope.LocalVars[name]
	if ok {
		varRef.VariableDecl = localVar
		return 0
	}

	// check in the class fields
	if varRef.MemberScope.ClassScope == nil {
		fmt.Println("didn't find class for var ref.")
		return 0
	}

	field, ok := varRef.MemberScope.ClassScope.FieldsByName[name]
	if ok {
		varRef.VariableDecl = field
		return 0
	}

	fmt.Println("didn't find source of var ref.")

	return 0 // these have no chuildren
}
