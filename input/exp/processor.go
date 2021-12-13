package exp

import "github.com/dragonfax/java_converter/input/parser"

// deal with the recursive expression tree.
func expressionProcessor(expression *parser.ExpressionContext) ExpressionNode {
	if expression == nil {
		return nil
	}

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

func StatementProcessor(statementCtx *parser.StatementContext) ExpressionNode {
	// TODO only one expression per block? no this isn't complicated enough.
	// but okay for a first of expression parsing

	if statementCtx.GetBlockLabel() != nil {
		return NewBlockNode(statementCtx.GetBlockLabel().(*parser.BlockContext))
	}

	if statementCtx.IF() != nil {
		return &IfNode{
			Condition: expressionProcessor(statementCtx.ParExpression().(*parser.ParExpressionContext).Expression().(*parser.ExpressionContext)),
			Body:      StatementProcessor(statementCtx.Statement(0).(*parser.StatementContext)),
			Else:      StatementProcessor(statementCtx.Statement(1).(*parser.StatementContext)),
		}
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
		return &ReturnNode{
			Expression: expressionProcessor(statementCtx.Expression(0).(*parser.ExpressionContext)),
		}
	}

	if statementCtx.THROW() != nil {
		return &ThrowNode{
			Expression: expressionProcessor(statementCtx.Expression(0).(*parser.ExpressionContext)),
		}
	}

	if statementCtx.BREAK() != nil {
		return &BreakNode{
			Label: statementCtx.IDENTIFIER().GetText(),
		}
	}

	if statementCtx.CONTINUE() != nil {
		return &ContinueNode{
			Label: statementCtx.IDENTIFIER().GetText(),
		}
	}

	/*
	   | TRY block (catchClause+ finallyBlock? | finallyBlock)
	   | TRY resourceSpecification block catchClause* finallyBlock?
	   | SWITCH parExpression '{' switchBlockStatementGroup* switchLabel* '}'

	   | statementExpression=expression ';'
	*/

	if statementCtx.GetIdentifierLabel() != nil {
		// must be a statement, with a label
		return &LabelNode{
			Label:      statementCtx.GetIdentifierLabel().GetText(),
			Expression: StatementProcessor(statementCtx.Statement(0).(*parser.StatementContext)),
		}
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
