package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/trans/node"
)

/* argument in a method delcaration.
 * not used in a method call, thats just an expression
 */
type Argument struct {
	*node.Base

	Type     *Type
	Name     string
	Ellipses bool
}

func (an *Argument) Children() []node.Node {
	return []node.Node{an.Type}
}

func (an *Argument) String() string {
	if an.Ellipses {
		return fmt.Sprintf("%s %s...", an.Name, an.Type)
	}
	return fmt.Sprintf("%s %s", an.Name, an.Type)
}

func NewArgument(typ *Type, name string, ellipses bool) *Argument {
	if typ == nil {
		panic(" no variable type")
	}
	if name == "" {
		panic("no variable name")
	}
	return &Argument{Base: node.New(), Type: typ, Name: name, Ellipses: ellipses}
}
