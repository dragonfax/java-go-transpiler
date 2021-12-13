package exp

import (
	"fmt"
)

type ExpressionNode interface {
	String() string
}

func expressionListToString(list []ExpressionNode) string {
	s := ""
	for _, node := range list {
		s += node.String() + "\n"
	}
	return s
}

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

func (bo *BinaryOperatorNode) String() string {
	return bo.Left.String() + string(bo.Operator) + bo.Right.String()
}

var _ ExpressionNode = &UnaryOperatorNode{}

type UnaryOperatorNode struct {
	BaseOperatorNode

	Left ExpressionNode
}

func (uo *UnaryOperatorNode) String() string {
	return string(uo.Operator) + uo.Left.String()
}

type LiteralNode struct {
	Value string
}

func (ln *LiteralNode) String() string {
	return ln.Value
}

type VariableNode struct {
	Name string
}

func (vn *VariableNode) String() string {
	return vn.Name
}

type IfNode struct {
	Condition ExpressionNode
	Body      ExpressionNode
	Else      ExpressionNode
}

func (in *IfNode) String() string {
	if in.Else == nil {
		return fmt.Sprintf("if %s {\n%s}\n", in.Condition, in.Body.String())
	}
	return fmt.Sprintf("if %s {\n%s} else {\n%s}\n", in.Condition, in.Body.String(), in.Else.String())
}

type ReturnNode struct {
	Expression ExpressionNode
}

func (rn *ReturnNode) String() string {
	return fmt.Sprintf("return %s\n", rn.Expression.String())
}

type ThrowNode struct {
	Expression ExpressionNode
}

func (tn *ThrowNode) String() string {
	return fmt.Sprintf("panic(%s)\n", tn.Expression.String())
}

type BreakNode struct {
	Label string
}

func (bn *BreakNode) String() string {
	return fmt.Sprintf("break %s\n", bn.Label)
}

type ContinueNode struct {
	Label string
}

func (cn *ContinueNode) String() string {
	return fmt.Sprintf("continue %s\n", cn.Label)
}

type LabelNode struct {
	Label      string
	Expression ExpressionNode
}

func (ln *LabelNode) String() string {
	return fmt.Sprintf("%s: %s\n", ln.Label, ln.Expression.String())
}
