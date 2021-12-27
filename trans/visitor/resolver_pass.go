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

func (rp *ResolvePass) VisitClassRef(classRef *ast.ClassRef) int {

	classRef.Class = classRef.ClassScope.ResolveClassName(classRef.ClassName)

	return 0
}

func (rp *ResolvePass) VisitClass(class *ast.Class) int {

	// base class
	if class.BaseClassName != "" {
		class.BaseClass = class.ResolveClassName(class.BaseClassName)
	}

	return rp.VisitChildren(class)
}

func (rp *ResolvePass) VisitMethodCall(methodCall *ast.MethodCall) int {

	// straight up MethodCall nodes don't yet even know their class name yet. its in a expression, seperated by a DOT binary operator or a chain.
	// but constructor calls do.

	if methodCall.Constructor {
		methodCall.Class = methodCall.MethodScope.ClassScope.ResolveClassName(methodCall.MethodName)
	}

	return rp.VisitChildren(methodCall)
}

func (rp *ResolvePass) VisitTypeElement(typeElement *ast.TypeElement) int {

	typeElement.Class = typeElement.ClassScope.ResolveClassName(typeElement.ClassName)

	return rp.VisitChildren(typeElement)
}

func (rp *ResolvePass) VisitTypePath(typePath *ast.TypePath) int {

	// visit children first, because we use their resolved classes
	rp.VisitChildren(typePath)

	// use the last element in the path as the path's real class.
	typePath.Class = typePath.Elements[len(typePath.Elements)-1].Class

	return 0
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
	if varRef.MethodScope == nil {
		ast.DebugPrint(varRef.GetParent())
		fmt.Println("no method scope for var ref")
		os.Exit(1)
	}

	localVar, ok := varRef.MethodScope.LocalVars[name]
	if ok {
		// fmt.Println("found var ref in method")
		varRef.VariableDecl = localVar
		return 0
	}

	// check in the class fields
	class := varRef.MethodScope.ClassScope
	if class == nil {
		// fmt.Println("didn't find class for var ref.")
		os.Exit(1)
	}

	field, ok := class.FieldsByName[name]
	if ok {
		// fmt.Println("found var ref in class")
		varRef.VariableDecl = field
		return 0
	}

	// check base class
	baseClass := class.BaseClass
	if baseClass != nil {
		field, ok := baseClass.FieldsByName[name]
		if ok {
			// fmt.Println("found var ref in base class")
			varRef.VariableDecl = field
			return 0
		}

	}

	refClass := class.ResolveClassName(name)
	if refClass != nil {
		// fmt.Println("found var ref as a class")
		varRef.VariableDecl = refClass
	}

	return 0 // these have no chuildren
}
