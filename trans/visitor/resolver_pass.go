package visitor

import (
	"fmt"
	"os"

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

	parent := varRef.GetParent()
	if chainParent, ok := parent.(*ast.Chain); ok {
		siblings := chainParent.Children()

		if siblings[0] != varRef {
			// must the first element in a chain.
			// or its not going to be resolvable by itself.
			return 0
		}
	}

	name := varRef.VariableName

	// check in the method arguments and other method variables
	if varRef.MemberScope == nil {
		ast.DebugPrint(varRef.GetParent())
		fmt.Println("no member scope for var ref")
		os.Exit(1)
		return 0
	}

	localVar, ok := varRef.MemberScope.LocalVars[name]
	if ok {
		fmt.Println("found var ref in member")
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
		fmt.Println("found var ref in class")
		varRef.VariableDecl = field
		return 0
	}

	// TODO check through the list of imported clases.

	fmt.Println("didn't find source of var ref.")
	ast.DebugPrint(varRef.GetParent())
	os.Exit(1)

	return 0 // these have no chuildren
}
