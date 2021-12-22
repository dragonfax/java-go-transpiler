package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/trans/node"
)

/* argument in a method delcaration.
 * not used in a method call, thats just an expression
 */
type ArgumentNode struct {
	*node.BaseNode

	Type     *TypeNode
	Name     string
	Ellipses bool
}

func (an *ArgumentNode) Children() []node.Node {
	return []node.Node{an.Type}
}

func (an *ArgumentNode) String() string {
	if an.Ellipses {
		return fmt.Sprintf("%s %s...", an.Name, an.Type)
	}
	return fmt.Sprintf("%s %s", an.Name, an.Type)
}

func NewArgument(typ *TypeNode, name string, ellipses bool) *ArgumentNode {
	if typ == nil {
		panic(" no variable type")
	}
	if name == "" {
		panic("no variable name")
	}
	return &ArgumentNode{BaseNode: node.NewNode(), Type: typ, Name: name, Ellipses: ellipses}
}
