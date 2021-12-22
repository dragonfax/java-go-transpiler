package ast

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/trans/node"
)

type Import struct {
	*node.BaseNode
	*BaseClassScope

	ImportString      string
	Star              bool
	ImportPackageName string
	ImportPackage     *Package
	ImportClass       *Class
}

func NewImport(s string) *Import {
	return &Import{BaseNode: node.NewNode(), BaseClassScope: NewClassScope(), ImportString: s}
}

func (i *Import) String() string {
	return fmt.Sprintf("import \"%s\"", i.ImportString)
}

func (i *Import) Children() []node.Node {
	return nil
}

func SplitPackageName(importString string) (string, string) {
	i := strings.LastIndex(importString, ".")
	packageName := importString[0:i]
	className := importString[i+1:]
	return packageName, className
}
