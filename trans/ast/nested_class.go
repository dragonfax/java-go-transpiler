package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/trans/node"
)

type NestedClass struct {
	*node.Base

	Name string
}

func (sc *NestedClass) Children() []node.Node {
	return nil
}

func NewNestedClass(name string) *NestedClass {
	return &NestedClass{Base: node.New(), Name: name}
}

func (sc *NestedClass) String() string {
	return fmt.Sprintf("\n// TODO elevate nested-class %s (pre-translation)\n\n", sc.Name)
}
