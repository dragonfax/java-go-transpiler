package exp

import (
	"fmt"

	"github.com/dragonfax/java_converter/tool"
)

type Operator string

const (
	Equals Operator = "="
)

type BaseOperatorNode struct {
	Operator Operator
}

var _ ExpressionNode = &BinaryOperatorNode{}

type BinaryOperatorNode struct {
	BaseOperatorNode

	Left  ExpressionNode
	Right ExpressionNode
}

func NewBinaryOperatorNode(operator Operator, left ExpressionNode, right ExpressionNode) *BinaryOperatorNode {
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
	this.Operator = operator
	return this
}

func (bo *BinaryOperatorNode) String() string {
	return fmt.Sprintf("%s%s%s", bo.Left, bo.Operator, bo.Right)
}

var _ ExpressionNode = &UnaryOperatorNode{}

type UnaryOperatorNode struct {
	BaseOperatorNode

	Left ExpressionNode
}

func NewUnaryOperatorNode(operator Operator, left ExpressionNode) *UnaryOperatorNode {
	if operator == "" {
		panic("no operator")
	}
	if tool.IsNilInterface(left) {
		panic("no expression")
	}
	this := &UnaryOperatorNode{Left: left}
	this.Operator = operator
	return this
}

func (uo *UnaryOperatorNode) String() string {
	return fmt.Sprintf("%s%s", uo.Operator, uo.Left)
}
