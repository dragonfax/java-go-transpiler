package exp

import (
	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/tool"
)

// deal with the recursive expression tree.
func expressionProcessor(expressionI parser.IExpressionContext) ExpressionNode {
	if tool.IsNilInterface(expressionI) {
		return nil
	}

	expression := expressionI.(*parser.ExpressionContext)

	if expression.Primary() != nil {
		return expressionFromPrimary(expression.Primary())
	}

	if expression.NEW() != nil {
		return NewConstructorCall(expression.Creator())
	}

	if expression.DOT() != nil {
		if expression.IDENTIFIER() != nil {
			return NewInstanceAttributeReference(expression.IDENTIFIER().GetText(), expressionProcessor(expression.Expression(0)))
		}
	}

	if expression.MethodCall() != nil {
		if expression.DOT() != nil {

			// method call

			// (variable with instance, method call (Identifier =name, expressionList = parameters) )
			if tool.IsNilInterface(expression.Expression(0)) {
				panic("expression with dot and 1 expression, but no identifier (not method call?\n" + expression.GetText() + "\n")
			}

			instance := expressionProcessor(expression.Expression(0))

			methodCall := expression.MethodCall()
			return NewMethodCall(instance, methodCall)
		} else {
			panic("method call without a preceding DOT")
		}
	}

	if expression.GetPrefix() != nil {
		// prefix operator
		return NewUnaryOperatorNode(true, expression.GetPrefix().GetText(), expressionProcessor(expression.Expression(0)))
	}

	if expression.GetPostfix() != nil {
		return NewUnaryOperatorNode(false, expression.GetPostfix().GetText(), expressionProcessor(expression.Expression(0)))
	}

	if expression.GetBop() != nil {
		if expression.COLON() != nil {
			left := expressionProcessor(expression.Expression(0))
			middle := expressionProcessor(expression.Expression(1))
			right := expressionProcessor(expression.Expression(2))
			return NewTernaryOperatorNode(expression.GetBop().GetText(), left, middle, right)
		}
		return NewBinaryOperatorNode(expression.GetBop().GetText(), expressionProcessor(expression.Expression(0)), expressionProcessor(expression.Expression(1)))
	}

	if len(expression.AllGT())+len(expression.AllLT()) > 0 {
		// shifting binary operator

		left := expressionProcessor(expression.Expression(0))
		right := expressionProcessor(expression.Expression(1))

		operator := ""
		for _, t := range expression.AllGT() {
			operator += t.GetText()
		}
		for _, t := range expression.AllLT() {
			operator += t.GetText()
		}

		return NewBinaryOperatorNode(operator, left, right)
	}

	panic("unknown expression: " + expression.GetText() + "\n" + expression.ToStringTree(parser.RuleNames, nil))

	// return nil
}

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
			expressionProcessor(statementCtx.ParExpression().(*parser.ParExpressionContext).Expression()),
			StatementProcessor(statementCtx.Statement(0)),
			StatementProcessor(statementCtx.Statement(1)),
		)
	}

	if statementCtx.FOR() != nil {
		return NewForNode(statementCtx)
	}

	if statementCtx.WHILE() != nil {
		return &ForNode{
			Condition: expressionProcessor(statementCtx.ParExpression().(*parser.ParExpressionContext).Expression().(*parser.ExpressionContext)),
			Body:      StatementProcessor(statementCtx.Statement(0).(*parser.StatementContext)),
		}
	}

	if statementCtx.DO() != nil {
		return &ForNode{
			Condition:     expressionProcessor(statementCtx.ParExpression().(*parser.ParExpressionContext).Expression().(*parser.ExpressionContext)),
			Body:          StatementProcessor(statementCtx.Statement(0).(*parser.StatementContext)),
			ConditionLast: true,
		}
	}

	if statementCtx.RETURN() != nil {
		return NewReturnNode(expressionProcessor(statementCtx.Expression(0).(*parser.ExpressionContext)))
	}

	if statementCtx.THROW() != nil {
		return NewThrowNode(expressionProcessor(statementCtx.Expression(0).(*parser.ExpressionContext)))
	}

	if statementCtx.BREAK() != nil {
		return NewBreakNode(statementCtx.IDENTIFIER().GetText())
	}

	if statementCtx.CONTINUE() != nil {
		return NewContinueNode(statementCtx.IDENTIFIER().GetText())
	}

	/*
	   | TRY block (catchClause+ finallyBlock? | finallyBlock)
	   | TRY resourceSpecification block catchClause* finallyBlock?
	   | SWITCH parExpression '{' switchBlockStatementGroup* switchLabel* '}'

	   | statementExpression=expression ';'
	*/

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
		// TODO warn, I dont' expect this to happen.
		return nil
	}
	if expressionCount >= 1 {

		if expressionCount > 1 {
			// TODO log this, should handle this scenario. its not uncommon.
		}

		// common scenario.
		return expressionProcessor(statementCtx.Expression(0).(*parser.ExpressionContext))
	}

	// ignore unknown structures.
	// TODO log them
	return nil
}

func expressionFromPrimary(primary parser.IPrimaryContext) ExpressionNode {
	primaryCtx := primary.(*parser.PrimaryContext)

	if primaryCtx.IDENTIFIER() != nil {
		return NewIdentifierNode(primaryCtx.IDENTIFIER().GetText())
	}

	if primaryCtx.THIS() != nil {
		return NewIdentifierNode("this")
	}

	if primaryCtx.SUPER() != nil {
		return NewIdentifierNode("super")
	}

	if primaryCtx.Expression() != nil {
		return expressionProcessor(primaryCtx.Expression())
	}

	if primaryCtx.Literal() != nil {
		literal := primaryCtx.Literal().(*parser.LiteralContext)
		return NewLiteralNode(literal.GetText())
	}

	panic("unknown primary type")
}
