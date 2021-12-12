package exp

import "github.com/dragonfax/delver_converter/parser"

type Operator string

const (
	Equals Operator = "="
)

type OperatorNode interface {
	String() string
}

type BaseOperatorNode struct {
	Operator Operator
}

var _ OperatorNode = &BinaryOperatorNode{}

type BinaryOperatorNode struct {
	BaseOperatorNode

	Left  OperatorNode
	Right OperatorNode
}

func (bo *BinaryOperatorNode) String() string {
	return bo.Left.String() + string(bo.Operator) + bo.Right.String()
}

var _ OperatorNode = &UnaryOperatorNode{}

type UnaryOperatorNode struct {
	BaseOperatorNode

	Left OperatorNode
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

// deal with the recursive expression tree.
func expressionProcessor(expression *parser.ExpressionContext) OperatorNode {

	var operator Operator
	if expression.ASSIGN() != nil {
		operator = Equals
	}

	subExpressions := expression.AllExpression()
	if len(subExpressions) == 1 {
		node := &UnaryOperatorNode{
			Left: expressionProcessor(subExpressions[0].(*parser.ExpressionContext)),
		}
		node.Operator = operator
		return node
	} else if len(subExpressions) == 2 {
		node := &BinaryOperatorNode{
			Left:  expressionProcessor(subExpressions[0].(*parser.ExpressionContext)),
			Right: expressionProcessor(subExpressions[1].(*parser.ExpressionContext)),
		}
		node.Operator = operator
		return node
	}

	// not a simple unary or binary operator

	if expression.Primary() != nil {
		primary := expression.Primary().(*parser.PrimaryContext)
		if primary.IDENTIFIER() != nil {
			node := &VariableNode{
				Name: primary.IDENTIFIER().GetText(),
			}
			return node
		} else if primary.Literal() != nil {
			literal := primary.Literal().(*parser.LiteralContext)
			return &LiteralNode{
				Value: literal.GetText(),
			}
		}
	}

	// TODO

	return nil
}
