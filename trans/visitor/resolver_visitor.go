package visitor

import (
	"fmt"

	"github.com/dragonfax/java_converter/trans/ast"
)

type ResolverVisitor struct {
	*BaseASTVisitor[int] // throwaway return value
}

func NewResolverVisitor(h *ast.Hierarchy) *ResolverVisitor {
	this := &ResolverVisitor{}
	this.BaseASTVisitor = NewASTVisitor[int](h, this)
	return this
}

func (cv *ResolverVisitor) VisitVarRef(varRef *ast.VarRef) int {

	if varRef.Super {
		return 0
	}

	if varRef.This {
		return 0
	}

	name := varRef.VariableName

	// check in the method arguments and other method variables
	localVar, ok := varRef.MemberScope.LocalVars[name]
	if ok {
		varRef.VariableDecl = localVar
		return 0
	}

	// check in the class fields
	field, ok := varRef.MemberScope.ClassScope.FieldsByName[name]
	if ok {
		varRef.VariableDecl = field
		return 0
	}

	fmt.Printf("didn't find source of var ref.\n")

	return 0 // these have no chuildren
}
