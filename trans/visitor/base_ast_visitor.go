/* visitor to process the AST, not the parse tree .
 *
 * makings writing optimization and translation passes easier.
 *
 * a rudimentar copy of the visitor system from antlr.
 */
package visitor

import (
	"fmt"

	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/ast"
	"github.com/dragonfax/java_converter/trans/node"

	"github.com/schollz/progressbar/v3"
)

var _ ASTVisitor[int] = &BaseASTVisitor[int]{}

/* this first attempt of an AST pass just defineds the interconnectedness between classes and packages */
type BaseASTVisitor[T comparable] struct {
	Hierarchy *ast.Hierarchy
	zero      T
	root      ASTVisitor[T]

	ProgressBar *progressbar.ProgressBar
}

func NewASTVisitor[T comparable](h *ast.Hierarchy, root ASTVisitor[T]) *BaseASTVisitor[T] {

	return &BaseASTVisitor[T]{
		Hierarchy:   h,
		root:        root,
		ProgressBar: progressbar.Default(h.ClassCount()),
	}
}

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

func (av *BaseASTVisitor[T]) VisitChildren(tree node.Node) T {
	var result T
	for _, child := range tree.Children() {
		if child == nil {
			fmt.Printf("someone delivered a nil child\n")
			continue
		}
		if tool.IsNilInterface(child) {
			fmt.Printf("someone delivered a typed nil child\n")
			continue
		}
		nextResult := av.VisitNode(child)
		if nextResult == av.zero && result == av.zero {
			// nothing
		} else if nextResult == av.zero && result != av.zero {
			// nothing
		} else if nextResult != av.zero && result == av.zero {
			result = nextResult
		} else {
			// both are not nil.
			result = av.root.AggregateResult(result, nextResult)
		}
	}
	return result
}

func (av *BaseASTVisitor[T]) AggregateResult(result, nextResult T) T {
	// just throws away siblings, and returns the last one in the list.
	return nextResult
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
