package hier

import "github.com/dragonfax/java_converter/trans/ast"

type Hierarchy struct {
	Classes []*ast.Class
}

func New() *Hierarchy {
	return &Hierarchy{Classes: make([]*ast.Class, 0)}
}

func (h *Hierarchy) AddFile(file *ast.File) {
	if file.Class != nil {
		h.Classes = append(h.Classes, file.Class)
	}
}
