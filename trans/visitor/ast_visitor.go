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
	"github.com/dragonfax/java_converter/trans/ast/exp"
	"github.com/dragonfax/java_converter/trans/node"
)

/* this first attempt of an AST pass just defineds the interconnectedness between classes and packages */
type ASTVisitor[T comparable] struct {
	Hierarchy *ast.Hierarchy
	zero      T

	// Context
	CurrentPackage *ast.Package
	CurrentClass   *ast.Class
	CurrentMethod  *node.Node
}

func NewASTVisitor[T comparable](h *ast.Hierarchy) *ASTVisitor[T] {
	return &ASTVisitor[T]{Hierarchy: h}
}

func (av *ASTVisitor[T]) VisitNode(tree node.Node) T {
	if tree == nil {
		fmt.Printf("someone gave us a nil node to visit\n")
		return av.zero
	}
	if class, ok := tree.(*ast.Class); ok {
		return av.VisitClass(class)
	} else if te, ok := tree.(*exp.TypeElementNode); ok {
		return av.VisitTypeElement(te)
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

func (av *ASTVisitor[T]) VisitClass(class *ast.Class) T {

	pkg := av.Hierarchy.GetPackage(class.PackageName)
	class.Package = pkg
	av.Hierarchy.AddClass(class)

	// above also adds imports and instantiates their references, all class references. created new classes as needed.

	return av.VisitChildren(class)
}

func (av *ASTVisitor[T]) VisitTypeElement(ctx *exp.TypeElementNode) T {

	// connect the type to its class,
	// its type arguments will get connected as children later.
	// will have to figure out hierarchy of known types in this class.
	// startign with imports, then the local package, then types in the same file.
	// ? how does the TypeElement know what class its currently in and the rest of its context?
	// I'm confused how this will work now.

	return av.VisitChildren(ctx)
}
