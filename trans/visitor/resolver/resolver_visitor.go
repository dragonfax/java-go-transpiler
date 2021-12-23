package resolver

import (
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

func (cv *ResolverVisitor) VisitVarRef(ctx *ast.VarRef) int {

	if ctx.Super {
		return 0
	}

	if ctx.This {
		return 0
	}

	// check in the method arguments

	ctx.MemberScope.Arguments()[varName]

	// check in the list of method local vars.

	ctx.MemberScope.LocalVars()[varName]

	// check in the class fields

	ctx.MemberScope.ClassScope().Fields[varName]

	// if we make mistakes we'll find them later, nd try to fix up the oce.

	// ctx.VariableDecl = ??

	return 0 // these have no chuildren
}
