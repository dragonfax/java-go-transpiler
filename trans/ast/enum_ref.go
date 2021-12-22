package ast

import "github.com/dragonfax/java_converter/trans/node"

type EnumRef struct {
	*node.BaseNode
	Name string
}

func (er *EnumRef) Children() []node.Node {
	return nil
}

func NewEnumRef(name string) *EnumRef {
	return &EnumRef{BaseNode: node.NewNode(), Name: name}
}

func (er *EnumRef) String() string {
	return er.Name
}
