package exp

import (
	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/tool"
)

func StatementProcessor(statementCtxI parser.IStatementContext) ExpressionNode {
	// TODO only one expression per block? no this isn't complicated enough.
	// but okay for a first of expression parsing

	if tool.IsNilInterface(statementCtxI) {
		return nil
	}

	statementCtx := statementCtxI.(*parser.StatementContext)

	if statementCtx.GetBlockLabel() != nil {
		return NewBlockNode(statementCtx.GetBlockLabel().(*parser.BlockContext))
	}

	if statementCtx.IF() != nil {
		return NewIfNode(
			ExpressionProcessor(statementCtx.ParExpression().(*parser.ParExpressionContext).Expression()),
			StatementProcessor(statementCtx.Statement(0)),
			StatementProcessor(statementCtx.Statement(1)),
		)
	}

	if statementCtx.FOR() != nil {
		return NewForNode(statementCtx)
	}

	if statementCtx.WHILE() != nil {
		return &ClassicForNode{
			Condition: ExpressionProcessor(statementCtx.ParExpression().(*parser.ParExpressionContext).Expression().(*parser.ExpressionContext)),
			Body:      StatementProcessor(statementCtx.Statement(0).(*parser.StatementContext)),
		}
	}

	if statementCtx.DO() != nil {
		return &ClassicForNode{
			Condition:     ExpressionProcessor(statementCtx.ParExpression().(*parser.ParExpressionContext).Expression().(*parser.ExpressionContext)),
			Body:          StatementProcessor(statementCtx.Statement(0).(*parser.StatementContext)),
			ConditionLast: true,
		}
	}

	if statementCtx.RETURN() != nil {
		return NewReturnNode(ExpressionProcessor(statementCtx.Expression(0)))
	}

	if statementCtx.THROW() != nil {
		return NewThrowNode(ExpressionProcessor(statementCtx.Expression(0).(*parser.ExpressionContext)))
	}

	if statementCtx.BREAK() != nil {
		if statementCtx.IDENTIFIER() != nil {
			return NewBreakNode(statementCtx.IDENTIFIER().GetText())
		}
		return NewBreakNode("")
	}

	if statementCtx.CONTINUE() != nil {
		if statementCtx.IDENTIFIER() != nil {
			return NewContinueNode(statementCtx.IDENTIFIER().GetText())
		}
		return NewContinueNode("")
	}

	if statementCtx.TRY() != nil {
		return NewTryCatchNode(statementCtx)
	}

	// TODO | SWITCH parExpression '{' switchBlockStatementGroup* switchLabel* '}'

	if statementCtx.GetIdentifierLabel() != nil {
		// must be a statement, with a label
		return NewLabelNode(
			statementCtx.GetIdentifierLabel().GetText(),
			StatementProcessor(statementCtx.Statement(0).(*parser.StatementContext)),
		)
	}

	// we dont' expect a lone statement
	// but we might get a lone expression.
	// check for both.

	statementCount := len(statementCtx.AllStatement())
	if statementCount >= 1 {

		if statementCount > 1 {
			// TODO log warning, really didn't expect this.
		}

		// TODO log warning, didn't expect this. missing grammar element?
		return StatementProcessor(statementCtx.Statement(0).(*parser.StatementContext))
	}

	// multiple expressions are possible in a single statement,
	// joined by commas, such as multi assignment.
	expressionCount := len(statementCtx.AllExpression())
	if expressionCount == 0 {
		// unimplemented
	}
	if expressionCount >= 1 {

		if expressionCount > 1 {
			// TODO log this, should handle this scenario. its not uncommon.
		}

		// common scenario.
		return ExpressionProcessor(statementCtx.Expression(0).(*parser.ExpressionContext))
	}

	// ignore unknown structures.
	// TODO log them
	return NewUnimplementedNode("unimplemented code: " + statementCtxI.GetText() + "\n\n" + statementCtxI.ToStringTree(parser.RuleNames, nil))
}
