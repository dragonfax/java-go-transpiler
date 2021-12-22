package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/node"
)

type BaseOperatorNode struct {
	*node.BaseNode

	Operator string
}

var _ node.Node = &BinaryOperatorNode{}

type BinaryOperatorNode struct {
	BaseOperatorNode

	Left  node.Node
	Right node.Node
}

func (bo *BinaryOperatorNode) Children() []node.Node {
	return []node.Node{bo.Left, bo.Right}
}

func NewBinaryOperatorNode(operator string, left node.Node, right node.Node) *BinaryOperatorNode {
	if operator == "" {
		panic("no operator")
	}
	if tool.IsNilInterface(left) {
		panic("no left expression")
	}
	if tool.IsNilInterface(right) {
		panic("no right expression")
	}
	this := &BinaryOperatorNode{Left: left, Right: right}
	this.BaseNode = node.NewNode()
	this.Operator = operator
	return this
}

func (bo *BinaryOperatorNode) String() string {
	if bo.Operator == "[" {
		return fmt.Sprintf("%s[%s]", bo.Left, bo.Right)
	}
	if bo.Operator == "(" {
		// cast operator
		return fmt.Sprintf("%s(%s)", bo.Left, bo.Right)
	}
	return fmt.Sprintf("%s%s%s", bo.Left, bo.Operator, bo.Right)
}

var _ node.Node = &UnaryOperatorNode{}

type UnaryOperatorNode struct {
	BaseOperatorNode
	Prefix bool

	Left node.Node
}

func (uo *UnaryOperatorNode) Children() []node.Node {
	return []node.Node{uo.Left}
}

func NewUnaryOperatorNode(prefix bool, operator string, left node.Node) *UnaryOperatorNode {
	if operator == "" {
		panic("no operator")
	}
	if tool.IsNilInterface(left) {
		panic("no expression")
	}
	this := &UnaryOperatorNode{Left: left, Prefix: prefix}
	this.BaseNode = node.NewNode()
	this.Operator = operator
	return this
}

func (uo *UnaryOperatorNode) String() string {
	if uo.Prefix {
		return fmt.Sprintf("%s%s", uo.Operator, uo.Left)
	}
	return fmt.Sprintf("%s%s", uo.Left, uo.Operator)
}

type TernaryOperatorNode struct {
	BaseOperatorNode

	Left   node.Node
	Middle node.Node
	Right  node.Node
}

func (to *TernaryOperatorNode) Children() []node.Node {
	return []node.Node{to.Left, to.Middle, to.Right}
}

func NewTernaryOperatorNode(operator string, left node.Node, middle node.Node, right node.Node) *TernaryOperatorNode {
	if operator == "" {
		panic("no operator")
	}
	if tool.IsNilInterface(left) {
		panic("no left expression")
	}
	if tool.IsNilInterface(middle) {
		panic("no middle")
	}
	if tool.IsNilInterface(right) {
		panic("no right expression")
	}
	this := &TernaryOperatorNode{Left: left, Middle: middle, Right: right}
	this.BaseNode = node.NewNode()
	this.Operator = operator
	return this
}

func (bo *TernaryOperatorNode) String() string {
	return fmt.Sprintf("ternary(%s,%s,%s)", bo.Left, bo.Middle, bo.Right)
}
