package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/node"
)

type BaseOperator struct {
	*node.Base

	Operator string
}

var _ node.Node = &BinaryOperator{}

func (bo *BaseOperator) Name() string {
	return bo.Operator
}

type BinaryOperator struct {
	BaseOperator

	Left  node.Node
	Right node.Node
}

func (bo *BinaryOperator) Children() []node.Node {
	return []node.Node{bo.Left, bo.Right}
}

func NewBinaryOperator(operator string, left node.Node, right node.Node) *BinaryOperator {
	if operator == "" {
		panic("no operator")
	}
	if tool.IsNilInterface(left) {
		panic("no left expression")
	}
	if tool.IsNilInterface(right) {
		panic("no right expression")
	}
	this := &BinaryOperator{Left: left, Right: right}
	this.Base = node.New()
	this.Operator = operator
	return this
}

func (bo *BinaryOperator) String() string {
	if bo.Operator == "[" {
		return fmt.Sprintf("%s[%s]", bo.Left, bo.Right)
	}
	if bo.Operator == "(" {
		// cast operator
		return fmt.Sprintf("%s(%s)", bo.Left, bo.Right)
	}
	return fmt.Sprintf("%s%s%s", bo.Left, bo.Operator, bo.Right)
}

var _ node.Node = &UnaryOperator{}

type UnaryOperator struct {
	BaseOperator
	Prefix bool

	Left node.Node
}

func (uo *UnaryOperator) Children() []node.Node {
	return []node.Node{uo.Left}
}

func NewUnaryOperator(prefix bool, operator string, left node.Node) *UnaryOperator {
	if operator == "" {
		panic("no operator")
	}
	if tool.IsNilInterface(left) {
		panic("no expression")
	}
	this := &UnaryOperator{Left: left, Prefix: prefix}
	this.Base = node.New()
	this.Operator = operator
	return this
}

func (uo *UnaryOperator) String() string {
	if uo.Prefix {
		return fmt.Sprintf("%s%s", uo.Operator, uo.Left)
	}
	return fmt.Sprintf("%s%s", uo.Left, uo.Operator)
}

type TernaryOperator struct {
	BaseOperator

	Left   node.Node
	Middle node.Node
	Right  node.Node
}

func (to *TernaryOperator) Children() []node.Node {
	return []node.Node{to.Left, to.Middle, to.Right}
}

func NewTernaryOperator(operator string, left node.Node, middle node.Node, right node.Node) *TernaryOperator {
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
	this := &TernaryOperator{Left: left, Middle: middle, Right: right}
	this.Base = node.New()
	this.Operator = operator
	return this
}

func (bo *TernaryOperator) String() string {
	return fmt.Sprintf("ternary(%s,%s,%s)", bo.Left, bo.Middle, bo.Right)
}
