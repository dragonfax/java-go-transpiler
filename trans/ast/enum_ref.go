package ast

import "github.com/dragonfax/java_converter/trans/node"

/* Only used in the label of a switch statement
 * TODO probably replaceable with an IdentRef
 */
type EnumRef struct {
	*node.Base
	Name string
}

func (er *EnumRef) Children() []node.Node {
	return nil
}

func NewEnumRef(name string) *EnumRef {
	return &EnumRef{Base: node.New(), Name: name}
}

func (er *EnumRef) String() string {
	return er.Name + " /* Unresolved */"
}
