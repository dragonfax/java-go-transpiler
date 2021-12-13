package exp

import (
	"fmt"

	"github.com/dragonfax/java_converter/tool"
)

type ExpressionNode interface {
	String() string
}

func expressionListToString(list []ExpressionNode) string {
	if list == nil {
		panic("list expression list")
	}
	s := ""
	for _, node := range list {
		if tool.IsNilInterface(node) {
			panic("nil node in expression list")
		}
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
	if bo == nil {
		panic("nil binary operator")
	}
	left := "<NIL>"
	if bo.Left != nil {
		left = bo.Left.String()
	}
	right := "<NIL>"
	if bo.Right != nil {
		right = bo.Right.String()
	}
	return left + string(bo.Operator) + right
}

var _ ExpressionNode = &UnaryOperatorNode{}

type UnaryOperatorNode struct {
	BaseOperatorNode

	Left ExpressionNode
}

func (uo *UnaryOperatorNode) String() string {
	if uo == nil {
		panic("nil unary operator")
	}
	left := "<NIL>"
	if uo.Left != nil {
		left = uo.Left.String()
	}
	return string(uo.Operator) + left
}

type LiteralNode struct {
	Value string
}

func (ln *LiteralNode) String() string {
	if ln == nil {
		panic("nil literal node")
	}
	return ln.Value
}

type VariableNode struct {
	Name string
}

func (vn *VariableNode) String() string {
	if vn == nil {
		panic("nil variable node")
	}
	return vn.Name
}

type IfNode struct {
	Condition ExpressionNode
	Body      ExpressionNode
	Else      ExpressionNode
}

func (in *IfNode) String() string {
	if in == nil {
		panic("nil if node")
	}
	if in.Else == nil {
		return fmt.Sprintf("if %s {\n%s}\n", in.Condition, in.Body.String())
	}
	return fmt.Sprintf("if %s {\n%s} else {\n%s}\n", in.Condition, in.Body.String(), in.Else.String())
}

type ReturnNode struct {
	Expression ExpressionNode
}

func (rn *ReturnNode) String() string {
	if rn == nil {
		panic("nil return node")
	}
	return fmt.Sprintf("return %s\n", rn.Expression.String())
}

type ThrowNode struct {
	Expression ExpressionNode
}

func (tn *ThrowNode) String() string {
	if tn == nil {
		panic("nil throw node")
	}
	return fmt.Sprintf("panic(%s)\n", tn.Expression.String())
}

type BreakNode struct {
	Label string
}

func (bn *BreakNode) String() string {
	if bn == nil {
		panic("nil break node")
	}
	return fmt.Sprintf("break %s\n", bn.Label)
}

type ContinueNode struct {
	Label string
}

func (cn *ContinueNode) String() string {
	if cn == nil {
		panic("nil continue node")
	}
	return fmt.Sprintf("continue %s\n", cn.Label)
}

type LabelNode struct {
	Label      string
	Expression ExpressionNode
}

func (ln *LabelNode) String() string {
	if ln == nil {
		panic("nil label node")
	}
	return fmt.Sprintf("%s: %s\n", ln.Label, ln.Expression.String())
}
