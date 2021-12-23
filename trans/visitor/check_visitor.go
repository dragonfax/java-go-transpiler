package visitor

import (
	"fmt"

	"github.com/dragonfax/java_converter/trans/ast"
)

/* check that all var refs are resolved */

type CheckVisitor struct {
	*BaseASTVisitor[int] // throwaway return value
}

func NewCheckVisitor(h *ast.Hierarchy) *CheckVisitor {
	this := &CheckVisitor{}
	this.BaseASTVisitor = NewASTVisitor[int](h, this)
	return this
}

func (cv *CheckVisitor) VisitVarRef(ctx *ast.VarRef) int {

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
