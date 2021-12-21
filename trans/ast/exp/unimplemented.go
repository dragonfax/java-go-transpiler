package exp

import "github.com/dragonfax/java_converter/trans/node"

type UnimplementedNode struct {
	Msg string
}

func (un *UnimplementedNode) Children() []node.Node {
	return nil
}

func (un *UnimplementedNode) String() string {
	return un.Msg
}

func NewUnimplementedNode(msg string) *UnimplementedNode {
	return &UnimplementedNode{Msg: msg}
}
