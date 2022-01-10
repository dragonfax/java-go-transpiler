package visitor

import (
	"github.com/dragonfax/java-go-transpiler/trans/ast"
	"github.com/dragonfax/java-go-transpiler/trans/node"
)

type ScopePass struct {
	*BaseASTVisitor[int] // throwaway return value

	// Context
	CurrentPackage *ast.Package
	CurrentClass   *ast.Class
	CurrentMethod  *ast.Method
}

func NewScopePass(h *ast.Hierarchy) ASTVisitor[int] {
	this := &ScopePass{}
	this.BaseASTVisitor = NewASTVisitor[int](h, this)
	return this
}

func (av *ScopePass) VisitChildren(tree node.Node) int {

	/* set the method and class scope */
	if scope, ok := tree.(ast.MethodScope); av.CurrentMethod != nil && ok {
		scope.SetMethodScope(av.CurrentMethod)
	}

	if scope, ok := tree.(ast.ClassScope); av.CurrentClass != nil && ok {
		scope.SetClassScope(av.CurrentClass)
	}

	// resume normal operation
	return av.BaseASTVisitor.VisitChildren(tree)
}

func (av *ScopePass) VisitPackage(pkg *ast.Package) int {
	av.CurrentPackage = pkg

	return av.VisitChildren(pkg)
}

func (av *ScopePass) VisitClass(class *ast.Class) int {
	defer av.ProgressBar.Add(1)

	av.CurrentClass = class
	// already has its package

	// TODO resolve import references, as those are't children for some reason.

	return av.VisitChildren(class)
}

func (av *ScopePass) VisitImport(imp *ast.Import) int {

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

	return av.VisitChildren(imp)
}

func (av *ScopePass) VisitMethod(method *ast.Method) int {

	av.CurrentMethod = method
	method.ClassScope = av.CurrentClass

	return av.VisitChildren(method)

}

func (av *ScopePass) VisitField(field *ast.Field) int {
	field.ClassScope = av.CurrentClass

	return av.VisitChildren(field)
}

func (cv *ScopePass) VisitLocalVarDecl(localVarDecl *ast.LocalVarDecl) int {
	localVarDecl.MethodScope = cv.CurrentMethod

	if localVarDecl.MethodScope == nil {
		panic("no method found around local var decl")
	}

	return cv.VisitChildren(localVarDecl)
}
