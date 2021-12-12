package main

import "github.com/dragonfax/delver_converter/parser"

type Operator string

const (
	Equals Operator = "="
)

type OperatorNode interface {
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

var _ OperatorNode = &UnaryOperatorNode{}

type UnaryOperatorNode struct {
	BaseOperatorNode

	Left OperatorNode
}

type LiteralNode struct {
	Value string
}

type VariableNode struct {
	Name string
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
