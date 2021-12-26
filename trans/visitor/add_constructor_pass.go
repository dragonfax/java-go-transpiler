package visitor

import (
	"github.com/dragonfax/java_converter/trans/ast"
	"github.com/dragonfax/java_converter/trans/node"
)

/* for each class, if it doesn't have a 0-args constructor, add one */
func AddConstructorPass(h *ast.Hierarchy) {
	for _, class := range h.GetClasses() {
		needs := true
		for _, item := range class.Members {
			if method, ok := item.(*ast.Member); ok {
				if method.Constructor && len(method.Arguments) == 0 {
					needs = false
				}
			}
		}

		if needs {
			// prepend the new zero-arg constructor
			class.Members = append([]node.Node{ast.NewEmptyConstructorFromClass(class)}, class.Members...)
		}
	}
}
