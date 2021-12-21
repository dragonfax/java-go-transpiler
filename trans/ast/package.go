package ast

import "github.com/dragonfax/java_converter/trans/node"

type Package struct {
	Name             string
	Classes          map[string]*Class
	ImportReferences []*Import
}

func NewPackage(name string) *Package {
	return &Package{
		Name:             name,
		Classes:          make(map[string]*Class),
		ImportReferences: make([]*Import, 0),
	}
}

func (pkg *Package) AddClass(class *Class) {
	pkg.Classes[class.Name] = class
}

func (pkg *Package) String() string {
	return "package " + pkg.Name
}

func (pkg *Package) Children() []node.Node {
	list := make([]node.Node, 0)
	for _, class := range pkg.Classes {
		list = append(list, class)
	}
	return list
}

func (pkg *Package) GetClass(className string) *Class {
	// creates the empty class if it doesn't exist.
	// by the time this is called we're already parsed all the files into ASTs and added them to the hierarchy
	// any classes created now are 3rd party, not part of the translated source.

	class := NewClass()

	class.Package = pkg

	return class
}

func (pkg *Package) AddImportReference(imp *Import) {
	// we need to track all import statements to this package, for later renaming things.

	pkg.ImportReferences = append(pkg.ImportReferences, imp)
}
