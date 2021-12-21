package visitor

import (
	"github.com/dragonfax/java_converter/trans/ast"
	"github.com/dragonfax/java_converter/trans/node"
)

type ScopeVisitor struct {
	*BaseASTVisitor[int] // throwaway return value

	// Context
	CurrentPackage *ast.Package
	CurrentClass   *ast.Class
	CurrentMethod  node.Node
}

func NewScopeVisitor(h *ast.Hierarchy) *ScopeVisitor {
	this := &ScopeVisitor{}
	this.BaseASTVisitor = NewASTVisitor[int](h, this)
	return this
}

func (av *ScopeVisitor) VisitPackage(pkg *ast.Package) int {
	av.CurrentPackage = pkg

	return av.VisitChildren(pkg)
}

func (av *ScopeVisitor) VisitClass(class *ast.Class) int {
	defer av.ProgressBar.Add(1)

	av.CurrentClass = class
	// already has its package

	// TODO resolve import references, as those are't children for some reason.

	return av.VisitChildren(class)
}

func (av *ScopeVisitor) VisitImport(imp *ast.Import) int {

	imp.ClassScope = av.CurrentClass

	impPackageName, impClassName := ast.SplitPackageName(imp.ImportString)

	impPkg := av.Hierarchy.GetPackage(impPackageName)
	imp.ImportPackage = impPkg

	if impClassName == "*" {
		imp.Star = true
		impPkg.AddImportReference(imp)
	} else {
		impClass := impPkg.GetClass(impClassName)
		imp.ImportClass = impClass
		impPkg.AddImportReference(imp)
	}

	return av.zero // no children
}

func (av *ScopeVisitor) VisitMethod(method *ast.Method) int {

	av.CurrentMethod = method
	method.ClassScope = av.CurrentClass

	return av.VisitChildren(method)

}

func (av *ScopeVisitor) VisitConstructor(constructor *ast.Constructor) int {
	av.CurrentMethod = constructor
	constructor.ClassScope = av.CurrentClass

	return av.VisitChildren(constructor)
}

func (av *ScopeVisitor) VisitField(field *ast.Field) int {
	field.ClassScope = av.CurrentClass

	return av.VisitChildren(field)
}

func (av *ScopeVisitor) VisitChildren(tree node.Node) int {
	if scope, ok := tree.(ast.MethodScope); av.CurrentMethod != nil && ok {
		scope.SetMethodScope(av.CurrentMethod)
	}

	if scope, ok := tree.(ast.ClassScope); av.CurrentClass != nil && ok {
		scope.SetClassScope(av.CurrentClass)
	}

	return av.BaseASTVisitor.VisitChildren(tree)
}
