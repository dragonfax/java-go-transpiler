package visitor

import (
	"github.com/dragonfax/java_converter/trans/ast"
	"github.com/dragonfax/java_converter/trans/node"
)

type ScopePass struct {
	*BaseASTVisitor[int] // throwaway return value

	// Context
	CurrentPackage *ast.Package
	CurrentClass   *ast.Class
	CurrentMember  *ast.Member
	ParentNode     node.Node
}

func NewScopePass(h *ast.Hierarchy) ASTVisitor[int] {
	this := &ScopePass{}
	this.BaseASTVisitor = NewASTVisitor[int](h, this)
	return this
}

func (av *ScopePass) VisitNode(tree node.Node) int {

	// fill in parents for all nodes.
	tree.SetParent(av.ParentNode)

	return av.BaseASTVisitor.VisitNode(tree)
}

func (av *ScopePass) VisitChildren(tree node.Node) int {

	/* set the method and class scope */
	if scope, ok := tree.(ast.MemberScope); av.CurrentMember != nil && ok {
		scope.SetMemberScope(av.CurrentMember)
	}

	if scope, ok := tree.(ast.ClassScope); av.CurrentClass != nil && ok {
		scope.SetClassScope(av.CurrentClass)
	}

	av.ParentNode = tree

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

	return av.zero // no children
}

func (av *ScopePass) VisitMember(member *ast.Member) int {

	av.CurrentMember = member
	member.ClassScope = av.CurrentClass

	return av.VisitChildren(member)

}

func (av *ScopePass) VisitField(field *ast.Field) int {
	field.ClassScope = av.CurrentClass

	return av.VisitChildren(field)
}
