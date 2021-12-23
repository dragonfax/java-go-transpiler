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

func (av *BaseASTVisitor[T]) VisitChildren(tree node.Node) T {
	var result T
	for _, child := range tree.Children() {
		if child == nil {
			ast.DebugPrint(tree)
			panic("someone delivered a nil child")
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
