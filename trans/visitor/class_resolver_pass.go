package visitor

import (
	"github.com/dragonfax/java-go-transpiler/trans/ast"
)

/* Resolve what various class references refer to within the entire package system
 * i.e. baseclass, static method reference.
 * Not to be confused with type resolution, which is for the value in a field or results of an expression.
 */
type ClassResolver struct {
	*BaseASTVisitor[int] // throwaway return value
}

func NewClassResolver(h *ast.Hierarchy) ASTVisitor[int] {
	this := &ClassResolver{}
	this.BaseASTVisitor = NewASTVisitor[int](h, this)
	return this
}

func (rp *ClassResolver) VisitClassRef(classRef *ast.ClassRef) int {

	classRef.Class = classRef.ClassScope.ResolveClassName(classRef.ClassName)

	return 0
}

func (rp *ClassResolver) VisitClass(class *ast.Class) int {

	// base class
	if class.BaseClassName != "" {
		class.BaseClass = class.ResolveClassName(class.BaseClassName)
	}

	return rp.VisitChildren(class)
}

func (rp *ClassResolver) VisitMethodCall(methodCall *ast.MethodCall) int {

	// straight up MethodCall nodes don't yet even know their class name yet. its in a expression, seperated by a DOT binary operator or a chain.
	// but constructor calls do.

	if methodCall.Constructor {
		methodCall.Class = methodCall.MethodScope.ClassScope.ResolveClassName(methodCall.MethodName)
	}

	return rp.VisitChildren(methodCall)
}

func (rp *ClassResolver) VisitTypeElement(typeElement *ast.TypeElement) int {

	typeElement.Class = typeElement.ClassScope.ResolveClassName(typeElement.ClassName)

	return rp.VisitChildren(typeElement)
}

func (rp *ClassResolver) VisitTypePath(typePath *ast.TypePath) int {

	// visit children first, because we use their resolved classes
	rp.VisitChildren(typePath)

	// use the last element in the path as the path's real class.
	typePath.Class = typePath.Elements[len(typePath.Elements)-1].Class

	return 0
}
