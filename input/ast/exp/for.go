package exp

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
)

type ForNode struct {
	Condition     ExpressionNode
	Init          []ExpressionNode
	Increment     []ExpressionNode
	Body          ExpressionNode
	ConditionLast bool // Do...While
}

func (fn *ForNode) String() string {
	// TODO ConditionLast
	// TODO remove unnecessary semicolons
	return fmt.Sprintf("for %s;%s;%s {\n%s}\n", fn.Init, fn.Condition, fn.Increment, fn.Body)
}

func NewForNode(statementCtx *parser.StatementContext) *ForNode {
	init, condition, increment := forControlProcessor(statementCtx.ForControl().(*parser.ForControlContext))
	return &ForNode{
		Condition: condition,
		Init:      init,
		Increment: increment,
		Body:      StatementProcessor(statementCtx.Statement(0).(*parser.StatementContext)),
	}
}

func forControlProcessor(forControlCtx *parser.ForControlContext) (init []ExpressionNode, condition ExpressionNode, increment []ExpressionNode) {
	if forControlCtx.EnhancedForControl() != nil {
		panic("didn't think we'd see these")
	}

	if forControlCtx.GetForUpdate() != nil {
		increment = make([]ExpressionNode, 0)
		for _, exp := range forControlCtx.GetForUpdate().(*parser.ExpressionListContext).AllExpression() {
			node := expressionProcessor(exp.(*parser.ExpressionContext))
			increment = append(increment, node)
		}
	}

	if forControlCtx.Expression() != nil {
		condition = expressionProcessor(forControlCtx.Expression().(*parser.ExpressionContext))
	}

	if forControlCtx.ForInit() != nil {
		initCtx := forControlCtx.ForInit().(*parser.ForInitContext)
		init = make([]ExpressionNode, 0)
		if initCtx.LocalVariableDeclaration() != nil {
			// variable declaractions
			declCtx := initCtx.LocalVariableDeclaration().(*parser.LocalVariableDeclarationContext)
			init = NewVariableDeclNodeList(declCtx)
		} else {
			// expression list
			for _, exp := range initCtx.ExpressionList().(*parser.ExpressionListContext).AllExpression() {
				node := expressionProcessor(exp.(*parser.ExpressionContext))
				init = append(init, node)
			}
		}
	}

	return
}
