package exp

import (
	"fmt"

	"github.com/dragonfax/java_converter/tool"
)

type BaseOperatorNode struct {
	Operator string
}

var _ ExpressionNode = &BinaryOperatorNode{}

type BinaryOperatorNode struct {
	BaseOperatorNode

	Left  ExpressionNode
	Right ExpressionNode
}

func NewBinaryOperatorNode(operator string, left ExpressionNode, right ExpressionNode) *BinaryOperatorNode {
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
	Prefix bool

	Left ExpressionNode
}

func NewUnaryOperatorNode(prefix bool, operator string, left ExpressionNode) *UnaryOperatorNode {
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

type TernaryOperatorNode struct {
	BaseOperatorNode

	Left   ExpressionNode
	Middle ExpressionNode
	Right  ExpressionNode
}

func NewTernaryOperatorNode(operator string, left ExpressionNode, middle ExpressionNode, right ExpressionNode) *TernaryOperatorNode {
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
	this.Operator = operator
	return this
}

func (bo *TernaryOperatorNode) String() string {
	return fmt.Sprintf("%s%s%s:%s", bo.Left, bo.Operator, bo.Middle, bo.Right)
}
