package ast

import "github.com/dragonfax/java_converter/trans/node"

type Hierarchy struct {
	*node.Base

	Packages map[string]*Package

	RootPackage string
}

func NewHierarchy() *Hierarchy {
	return &Hierarchy{
		Base:     node.New(),
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
		list = append(list, pkg)
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
	class.PackageScope = pkg
	pkg.AddClass(class)
}

/* GetClasses, list of all classes,
 * used at the end of processesing to output all the classes as files in the target directory
 */
func (h *Hierarchy) GetClasses() []*Class {
	list := make([]*Class, 0)
	for _, pkg := range h.Packages {
		for _, class := range pkg.Classes {
			list = append(list, class)
		}
	}
	return list
}
