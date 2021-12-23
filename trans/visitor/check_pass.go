package visitor

import (
	"fmt"

	"github.com/dragonfax/java_converter/trans/ast"
)

/* check that all var refs are resolved */

type CheckPass struct {
	*BaseASTVisitor[int] // throwaway return value
}

func NewCheckPass(h *ast.Hierarchy) *CheckPass {
	this := &CheckPass{}
	this.BaseASTVisitor = NewASTVisitor[int](h, this)
	return this
}

func (cv *CheckPass) VisitVarRef(ctx *ast.VarRef) int {

	if ctx.Super {
		return 0
	}

	if ctx.This {
		return 0
	}

	if ctx.VariableDecl == nil {
		fmt.Printf("Var Ref Unresolved: %s from %s\n", ctx.VariableName, ctx.MemberScope.String())
	}

	return 0 // these have no chuildren
}
