package ast

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/trans/node"
)

type Import struct {
	*node.Base
	*BaseClassScope

	ImportString      string // the original string from the java statement. package + (* or class)
	Star              bool
	ImportPackageName string
	ImportPackage     *Package
	ImportClass       *Class
}

func NewImport(s string) *Import {
	return &Import{Base: node.New(), BaseClassScope: NewClassScope(), ImportString: s}
}

func (i *Import) Name() string {
	return i.ImportString
}

func (i *Import) String() string {
	return fmt.Sprintf("import \"%s\"", i.ClassScope.PackageScope.RootPackage()+"/"+strings.ReplaceAll(i.ImportPackage.QualifiedName, ".", "/"))
}

func (i *Import) Children() []node.Node {
	return nil
}

func SplitPackageName(importString string) (string, string) {
	index := strings.LastIndex(importString, ".")
	if index == -1 {
		panic("unable to split imported package name")
	}
	packageName := importString[0:index]
	className := importString[index+1:]
	return packageName, className
}
