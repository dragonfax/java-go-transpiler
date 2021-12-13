package exp

import "github.com/dragonfax/delver_converter/parser"

// deal with the recursive expression tree.
func expressionProcessor(expression *parser.ExpressionContext) ExpressionNode {
	if expression == nil {
		return nil
	}

	var operator Operator
	if expression.ASSIGN() != nil {
		operator = Equals
	} else if expression.RETURN() {

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

	if statementCtx.IF() != nil {
		return &IfNode{
			Condition: expressionProcessor(statementCtx.ParExpression().(*parser.ParExpressionContext).Expression().(*parser.ExpressionContext)),
			Body:      StatementProcessor(statementCtx.Statement(0).(*parser.StatementContext)),
			Else:      StatementProcessor(statementCtx.Statement(1).(*parser.StatementContext)),
		}
	}

	if statementCtx.FOR() != nil {
		init, condition, increment := forControlProcessor(statementCtx.ForControl())
		return &ForNode{
			Condition: condition,
			Init:      init,
			Increment: increment,
			Body:      StatementProcessor(statementCtx.Statement(0).(*parser.StatementContext)),
		}
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
			init = localVariableDeclarationProcessor(declCtx)
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

func localVariableDeclarationProcessor(decl *parser.LocalVariableDeclarationContext) []ExpressionNode {

	l := make([]ExpressionNode, 0)

	typ := decl.TypeType().GetText()

	for _, varDecl := range decl.VariableDeclarators().(*parser.VariableDeclaratorsContext).AllVariableDeclarator() {

		varDeclCtx := varDecl.(*parser.VariableDeclaratorContext)

		varInitCtx := varDeclCtx.VariableInitializer().(*parser.VariableInitializerContext)

		exp := variableInitializerProcessor(varInitCtx)

		node := &VariableDeclNode{
			Type:       typ,
			Name:       varDeclCtx.VariableDeclaratorId().GetText(),
			Expression: exp,
		}

		l = append(l, node)
	}

	return l
}

func variableInitializerProcessor(ctx *parser.VariableInitializerContext) ExpressionNode {
	var exp ExpressionNode
	if ctx.Expression() != nil {
		exp = expressionProcessor(ctx.Expression().(*parser.ExpressionContext))
	}
	if ctx.ArrayInitializer() != nil {
		exp = arrayInitializerProcessor(ctx.ArrayInitializer().(*parser.ArrayInitializerContext))
	}

	return exp
}

func arrayInitializerProcessor(ctx *parser.ArrayInitializerContext) *ArrayLiteral {

	if len(ctx.AllVariableInitializer()) == 0 {
		return &ArrayLiteral{}
	}

	l := make([]ExpressionNode, 0)
	for _, varInit := range ctx.AllVariableInitializer() {
		varInitCtx := varInit.(*parser.VariableInitializerContext)

		exp := variableInitializerProcessor(varInitCtx)

		l = append(l, exp)
	}

	return &ArrayLiteral{
		Elements: l,
	}
}
