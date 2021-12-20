package hier

import "github.com/dragonfax/java_converter/trans/ast"

type Hierarchy struct {
	Classes  []*ast.Class
	Packages map[string]*Package
}

func New() *Hierarchy {
	return &Hierarchy{
		Classes:  make([]*ast.Class, 0),
		Packages: make(map[string]*Package),
	}
}

func (h *Hierarchy) AddClass(packageName string, class *ast.Class) {
	h.Classes = append(h.Classes, class)

	pack, ok := h.Packages[packageName]
	if !ok {
		pack = NewPackage(packageName)
		h.Packages[packageName] = pack
		pack.Classes[class.Name] = class
	}
}

type Package struct {
	Name    string
	Classes map[string]*ast.Class
}

func NewPackage(name string) *Package {
	return &Package{
		Name:    name,
		Classes: make(map[string]*ast.Class),
	}
}
