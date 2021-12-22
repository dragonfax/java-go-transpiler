package ast

import "github.com/dragonfax/java_converter/trans/node"

type UnimplementedNode struct {
	*node.BaseNode
	Msg string
}

func (un *UnimplementedNode) Children() []node.Node {
	return nil
}

func (un *UnimplementedNode) String() string {
	return un.Msg
}

func NewUnimplementedNode(msg string) *UnimplementedNode {
	return &UnimplementedNode{BaseNode: node.NewNode(), Msg: msg}
}
