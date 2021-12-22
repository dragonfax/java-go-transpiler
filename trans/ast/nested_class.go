package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/trans/node"
)

type NestedClass struct {
	*node.BaseNode

	Name string
}

func (sc *NestedClass) Children() []node.Node {
	return nil
}

func NewNestedClass(name string) *NestedClass {
	return &NestedClass{BaseNode: node.NewNode(), Name: name}
}

func (sc *NestedClass) String() string {
	return fmt.Sprintf("\n// TODO elevate subclass %s (pre-translation)\n\n", sc.Name)
}
