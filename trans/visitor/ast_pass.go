package visitor

import (
	"github.com/dragonfax/java_converter/trans/ast"
)

func (av *ASTVisitor[T]) VisitPackage(pkg *ast.Package) T {
	av.CurrentPackage = pkg

	return av.VisitChildren(pkg)
}

func (av *ASTVisitor[T]) VisitClass(class *ast.Class) T {
	defer av.ProgressBar.Add(1)

	av.CurrentClass = class
	// already has its package

	// TODO resolve import references, as those are't children for some reason.

	return av.VisitChildren(class)
}

func (av *ASTVisitor[T]) VisitTypeElement(node *ast.TypeElementNode) T {

	// connect the type to its class,
	// its type arguments will get connected as children later.
	// will have to figure out hierarchy of known types in this class.
	// startign with imports, then the local package, then types in the same file.
	// ? how does the TypeElement know what class its currently in and the rest of its context?
	// I'm confused how this will work now.

	return av.VisitChildren(node)
}

func (av *ASTVisitor[T]) VisitMethod(method *ast.Method) T {

	av.CurrentMethod = method
	method.Class = av.CurrentClass

	return av.VisitChildren(method)

}

func (av *ASTVisitor[T]) VisitConstructor(constructor *ast.Constructor) T {
	av.CurrentMethod = constructor
	constructor.Class = av.CurrentClass

	return av.VisitChildren(constructor)
}
