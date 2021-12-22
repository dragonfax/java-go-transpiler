package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/node"
)

type IfNode struct {
	*node.BaseNode

	Condition node.Node
	Body      node.Node
	Else      node.Node
}

func (in *IfNode) Children() []node.Node {
	list := []node.Node{in.Body}
	if in.Condition != nil {
		list = append(list, in.Condition)
	}
	if in.Else != nil {
		list = append(list, in.Else)
	}
	return list
}

func NewIfNode(condition, body, els node.Node) *IfNode {
	if tool.IsNilInterface(body) {
		panic("missing body")
	}
	if tool.IsNilInterface(condition) {
		panic("missing condition")
	}
	return &IfNode{
		BaseNode:  node.NewNode(),
		Condition: condition,
		Body:      body,
		Else:      els,
	}
}

func (in *IfNode) String() string {
	if tool.IsNilInterface(in.Else) {
		return fmt.Sprintf("if %s {\n%s}\n", in.Condition, in.Body)
	}
	return fmt.Sprintf("if %s {\n%s} else {\n%s}\n", in.Condition, in.Body, in.Else)
}
