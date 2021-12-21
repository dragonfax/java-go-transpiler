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

/* this first attempt of an AST pass just defineds the interconnectedness between classes and packages */
type ASTVisitor[T comparable] struct {
	Hierarchy *ast.Hierarchy
	zero      T

	// Context
	CurrentPackage *ast.Package
	CurrentClass   *ast.Class
	CurrentMethod  node.Node

	ProgressBar *progressbar.ProgressBar
}

func NewASTVisitor[T comparable](h *ast.Hierarchy) *ASTVisitor[T] {

	return &ASTVisitor[T]{
		Hierarchy:   h,
		ProgressBar: progressbar.Default(h.ClassCount()),
	}
}

func (av *ASTVisitor[T]) VisitNode(tree node.Node) T {
	if tree == nil {
		fmt.Printf("someone gave us a nil node to visit\n")
		return av.zero
	}

	if scope, ok := tree.(ast.Scope); av.CurrentMethod != nil && ok {
		scope.SetScope(av.CurrentMethod)
	}

	if class, ok := tree.(*ast.Class); ok {
		return av.VisitClass(class)
	} else if te, ok := tree.(*ast.TypeElementNode); ok {
		return av.VisitTypeElement(te)
	} else if pkg, ok := tree.(*ast.Package); ok {
		return av.VisitPackage(pkg)
	} else if method, ok := tree.(*ast.Method); ok {
		return av.VisitMethod(method)
	} else if con, ok := tree.(*ast.Constructor); ok {
		return av.VisitConstructor(con)
	} else {
		return av.VisitChildren(tree)
	}
}

func (av *ASTVisitor[T]) VisitChildren(tree node.Node) T {
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
			// TODO aggregate function
		}
	}
	return av.zero
}
