package ast

import "github.com/dragonfax/java_converter/trans/node"

type Hierarchy struct {
	Packages map[string]*Package
}

func NewHierarchy() *Hierarchy {
	return &Hierarchy{
		Packages: make(map[string]*Package),
	}
}

func (h *Hierarchy) String() string {
	return "hierarchy"
}

func (h *Hierarchy) ClassCount() int64 {
	n := 0
	for _, pkg := range h.Packages {
		n += len(pkg.Classes)
	}
	return int64(n)
}

func (h *Hierarchy) Children() []node.Node {
	list := make([]node.Node, 0)
	for _, pkg := range h.Packages {
		for _, class := range pkg.Classes {
			list = append(list, class)
		}
	}

	return list
}

func (h *Hierarchy) GetPackage(packageName string) *Package {
	pkg, ok := h.Packages[packageName]
	if !ok {
		pkg = NewPackage(packageName)
		h.Packages[packageName] = pkg
	}
	return pkg
}

// AddClass called before AST crawling. to prepopulate the classes we already have from the parse tree.
func (h *Hierarchy) AddClass(class *Class) {
	packageName := class.PackageName
	pkg := h.GetPackage(packageName)
	class.Package = pkg
	pkg.AddClass(class)
}

type Package struct {
	Name    string
	Classes map[string]*Class
}

func NewPackage(name string) *Package {
	return &Package{
		Name:    name,
		Classes: make(map[string]*Class),
	}
}

func (pkg *Package) AddClass(class *Class) {
	pkg.Classes[class.Name] = class
}
