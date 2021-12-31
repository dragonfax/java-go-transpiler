package visitor

import "github.com/dragonfax/java_converter/trans/ast"

func BaseClassPass(h *ast.Hierarchy) {
	for _, class := range h.GetClasses() {
		if class.BaseClassName != "" {
			runtimePkg := h.GetPackage("runtime")
			class.BaseClass = class.ResolveClassName(runtimePkg, class.BaseClassName)
		}
	}
}
