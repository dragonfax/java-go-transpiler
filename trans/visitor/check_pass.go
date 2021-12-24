package visitor

import (
	"fmt"
	"os"

	"github.com/dragonfax/java_converter/trans/ast"
)

/* check that all var refs are resolved */

type CheckPass struct {
	*BaseASTVisitor[int] // throwaway return value
}

func NewCheckPass(h *ast.Hierarchy) ASTVisitor[int] {
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
		fmt.Printf("Var Ref Unresolved: %s\n", ctx.VariableName)
		if ctx.GetParent() == nil {
			fmt.Println("varref has no parent")
			ast.DebugPrint(ctx)
			os.Exit(1)
		}
	}

	return 0 // these have no chuildren
}
