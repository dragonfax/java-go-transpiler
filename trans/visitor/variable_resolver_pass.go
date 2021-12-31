package visitor

import (
	"fmt"
	"os"

	"github.com/dragonfax/java_converter/trans/ast"
)

/* Determine what each variable reference refers to, a field, on this class
 * or another class, or is it the name of a class itself.
 */
type VarResolver struct {
	*BaseASTVisitor[int] // throwaway return value
}

func NewVarResolver(h *ast.Hierarchy) ASTVisitor[int] {
	this := &VarResolver{}
	this.BaseASTVisitor = NewASTVisitor[int](h, this)

	return this
}

func (cv *VarResolver) VisitIdentRef(identRef *ast.IdentRef) int {

	if identRef.Super {
		return 0
	}

	if identRef.This {
		return 0
	}

	name := identRef.Name

	// check in the method arguments and other method variables
	if identRef.MethodScope == nil {
		ast.DebugPrint(identRef.GetParent())
		fmt.Println("no method scope for var ref")
		os.Exit(1)
	}

	localVar, ok := identRef.MethodScope.LocalVars[name]
	if ok {
		// fmt.Println("found var ref in method")
		identRef.LocalVariableDecl = localVar
		return 0
	}

	// check in the class fields
	class := identRef.MethodScope.ClassScope
	if class == nil {
		// fmt.Println("didn't find class for var ref.")
		os.Exit(1)
	}

	field, ok := class.FieldsByName[name]
	if ok {
		// fmt.Println("found var ref in class")
		identRef.Field = field
		return 0
	}

	// check base class
	baseClass := class.BaseClass
	if baseClass != nil {
		field, ok := baseClass.FieldsByName[name]
		if ok {
			// fmt.Println("found var ref in base class")
			identRef.Field = field
			return 0
		}

	}

	refClass := class.ResolveClassName(name)
	if refClass != nil {
		// fmt.Println("found var ref as a class")
		identRef.Class = refClass
	}

	return 0 // these have no chuildren
}
