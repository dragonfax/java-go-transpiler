package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/node"
)

type BaseOperator struct {
	*BaseExpression

	Operator string
}

var _ node.Node = &BinaryOperator{}

func (bo *BaseOperator) NodeName() string {
	return bo.Operator
}

type BinaryOperator struct {
	*BaseOperator

	Left  Expression
	Right Expression
}

func (bo *BinaryOperator) Children() []node.Node {
	return []node.Node{bo.Left, bo.Right}
}

func NewOperator(operator string) *BaseOperator {
	return &BaseOperator{BaseExpression: NewExpression(), Operator: operator}
}

func NewBinaryOperator(operator string, left, right Expression) *BinaryOperator {
	if operator == "" {
		panic("no operator")
	}
	if tool.IsNilInterface(left) {
		panic("no left expression")
	}
	if tool.IsNilInterface(right) {
		panic("no right expression")
	}
	return &BinaryOperator{BaseOperator: NewOperator(operator), Left: left, Right: right}
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

type UnaryOperator struct {
	*BaseOperator
	Prefix bool

	Left Expression
}

func (uo *UnaryOperator) Children() []node.Node {
	return []node.Node{uo.Left}
}

func NewUnaryOperator(prefix bool, operator string, left Expression) node.Node {
	if operator == "" {
		panic("no operator")
	}
	if tool.IsNilInterface(left) {
		panic("no expression")
	}

	// a single dash is a negative. A negative number sometimes parses this way.
	if operator == "-" {
		if literal, ok := left.(*Literal); ok {
			switch literal.LiteralType {
			case Integer:
				return NewLiteral(Integer, "-"+literal.Value)
			case Float:
				return NewLiteral(Float, "-"+literal.Value)
			default:
				panic("use of a negative against non-numberic literal")
			}
		}
		// else we can make a multiplication against -1
		return NewBinaryOperator("*", NewLiteral(Integer, "-1"), left)
	}

	this := &UnaryOperator{BaseOperator: NewOperator(operator), Left: left, Prefix: prefix}
	return this
}

func (uo *UnaryOperator) String() string {
	if uo.Prefix {
		return fmt.Sprintf("%s%s", uo.Operator, uo.Left)
	}
	return fmt.Sprintf("%s%s", uo.Left, uo.Operator)
}

type TernaryOperator struct {
	*BaseOperator

	Left   Expression
	Middle Expression
	Right  Expression
}

func (to *TernaryOperator) Children() []node.Node {
	return []node.Node{to.Left, to.Middle, to.Right}
}

func NewTernaryOperator(operator string, left, middle, right Expression) *TernaryOperator {
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
	this := &TernaryOperator{BaseOperator: NewOperator(operator), Left: left, Middle: middle, Right: right}
	return this
}

func (bo *TernaryOperator) String() string {
	return fmt.Sprintf("ternary(%s,%s,%s)", bo.Left, bo.Middle, bo.Right)
}
