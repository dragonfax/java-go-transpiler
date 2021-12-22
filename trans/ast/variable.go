package ast

import "github.com/dragonfax/java_converter/trans/node"

type VariableNode struct {
	Name string
}

func NewVariableNode(name string) *VariableNode {
	if name == "" {
		panic("missing name")
	}
	return &VariableNode{
		Name: name,
	}
}

func (vn *VariableNode) Children() []node.Node {
	return nil
}

func (vn *VariableNode) String() string {
	return vn.Name
}
