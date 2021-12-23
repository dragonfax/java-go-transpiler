package visitor

import (
	"fmt"

	"github.com/dragonfax/java_converter/trans/ast"
	"github.com/dragonfax/java_converter/trans/node"
)

func (av *BaseASTVisitor[T]) VisitNode(tree node.Node) T {
	if tree == nil {
		fmt.Printf("someone gave us a nil node to visit\n")
		return av.zero
	}

	if class, ok := tree.(*ast.Class); ok {
		return av.root.VisitClass(class)
	} else if pkg, ok := tree.(*ast.Package); ok {
		return av.root.VisitPackage(pkg)
	} else if method, ok := tree.(*ast.Member); ok {
		return av.root.VisitMember(method)
	} else if imp, ok := tree.(*ast.Import); ok {
		return av.root.VisitImport(imp)
	} else if field, ok := tree.(*ast.Field); ok {
		return av.root.VisitField(field)
	} else if varRef, ok := tree.(*ast.VarRef); ok {
		return av.root.VisitVarRef(varRef)
	} else if localVarDecl, ok := tree.(*ast.LocalVarDecl); ok {
		return av.root.VisitLocalVarDecl(localVarDecl)
	} else {
		return av.root.VisitChildren(tree)
	}
}

func (av *BaseASTVisitor[T]) VisitField(tree *ast.Field) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitPackage(tree *ast.Package) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitClass(tree *ast.Class) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitImport(tree *ast.Import) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitMember(tree *ast.Member) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitVarRef(tree *ast.VarRef) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitLocalVarDecl(tree *ast.LocalVarDecl) T {
	return av.root.VisitChildren(tree)
}
