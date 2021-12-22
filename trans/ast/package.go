package ast

import "github.com/dragonfax/java_converter/trans/node"

type Package struct {
	*node.Base

	Name             string
	Classes          map[string]*Class
	ImportReferences []*Import
}

func NewPackage(name string) *Package {
	return &Package{
		Base:             node.New(),
		Name:             name,
		Classes:          make(map[string]*Class),
		ImportReferences: make([]*Import, 0),
	}
}

/* AddClass, only for use before the AST walking phases begin
 * After that, the package is responsible for creating new classes,
 * since they will be empty reference classes.
 */
func (pkg *Package) AddClass(class *Class) {
	if _, ok := pkg.Classes[class.Name]; ok {
		panic("already have this class")
	}
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

	if class, ok := pkg.Classes[className]; ok {
		return class
	}

	class := NewClass()
	class.Name = className
	class.PackageScope = pkg
	class.Parent = pkg
	class.Generated = true

	pkg.Classes[className] = class

	return class
}

func (pkg *Package) AddImportReference(imp *Import) {
	// we need to track all import statements to this package, for later renaming things.

	/* at this point, the class thats doing the importing (the importer)
	 * has already asked for (and possibly created)
	 * the class being imported from this package. (the importee)
	 */

	pkg.ImportReferences = append(pkg.ImportReferences, imp)

}
