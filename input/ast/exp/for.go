package exp

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
)

func NewForNode(statementCtx *parser.StatementContext) ExpressionNode {

	if statementCtx.ForControl().(*parser.ForControlContext).EnhancedForControl() != nil {
		return NewEnhancedForNode(statementCtx)
	}
	return NewClassicForNode(statementCtx)
}

type EnhancedForNode struct {
	Variable ExpressionNode
	Instance ExpressionNode
	Body     ExpressionNode
}

func NewEnhancedForNode(statementCtx *parser.StatementContext) *EnhancedForNode {

	forControlCtx := statementCtx.ForControl().(*parser.ForControlContext)
	enhancedCtx := forControlCtx.EnhancedForControl().(*parser.EnhancedForControlContext)

	instance := expressionProcessor(enhancedCtx.Expression())
	variable := NewVariableDecl(enhancedCtx.TypeType().GetText(), enhancedCtx.VariableDeclaratorId().GetText(), nil, false)

	body := StatementProcessor(statementCtx.Statement(0).(*parser.StatementContext))
	return &EnhancedForNode{
		Variable: variable,
		Instance: instance,
		Body:     body,
	}
}

func (ef *EnhancedForNode) String() string {
	return fmt.Sprintf("for ( %s : %s ) %s", ef.Variable, ef.Instance, ef.Body)
}

type ClassicForNode struct {
	Condition     ExpressionNode
	Init          []ExpressionNode
	Increment     []ExpressionNode
	Body          ExpressionNode
	ConditionLast bool // Do...While
}

func (fn *ClassicForNode) String() string {
	// TODO ConditionLast
	// TODO remove unnecessary semicolons
	return fmt.Sprintf("for %s;%s;%s {\n%s}\n", fn.Init, fn.Condition, fn.Increment, fn.Body)
}

func NewClassicForNode(statementCtx *parser.StatementContext) *ClassicForNode {
	init, condition, increment := classicForControlProcessor(statementCtx.ForControl().(*parser.ForControlContext))
	return &ClassicForNode{
		Condition: condition,
		Init:      init,
		Increment: increment,
		Body:      StatementProcessor(statementCtx.Statement(0).(*parser.StatementContext)),
	}
}

func classicForControlProcessor(forControlCtx *parser.ForControlContext) (init []ExpressionNode, condition ExpressionNode, increment []ExpressionNode) {
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
