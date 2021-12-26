package visitor

import "github.com/dragonfax/java_converter/trans/ast"

func BaseClassPass(h *ast.Hierarchy) {
	for _, class := range h.GetClasses() {
		if class.BaseClassName != "" {
			class.BaseClass = class.ResolveClassName(class.BaseClassName)
		}
	}
}