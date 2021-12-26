package ast

import (
	"strings"

	"github.com/dragonfax/java_converter/trans/node"
)

type Package struct {
	*node.Base

	QualifiedName    string
	Classes          map[string]*Class
	ImportReferences []*Import
}

func NewPackage(name string) *Package {
	return &Package{
		Base: node.New(),

		/* a package name never has a class name in it (like an import name might have) */
		QualifiedName:    name,
		Classes:          make(map[string]*Class),
		ImportReferences: make([]*Import, 0),
	}
}

func (pkg *Package) Basename() string {
	parts := strings.Split(pkg.QualifiedName, ".")
	return parts[len(parts)-1]
}

func (pkg *Package) Dir() string {
	return strings.Join(strings.Split(pkg.QualifiedName, "."), "/")
}

func (pkg *Package) RootPackage() string {
	return pkg.GetParent().(*Hierarchy).RootPackage
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

func (pkg *Package) HasClass(className string) *Class {
	return pkg.Classes[className]
}

func (pkg *Package) String() string {
	return "package " + pkg.Basename()
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
