package exp

import (
	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/tool"
)

// deal with the recursive expression tree.
func ExpressionProcessor(expressionI parser.IExpressionContext) ExpressionNode {
	if tool.IsNilInterface(expressionI) {
		return nil
	}

	expression := expressionI.(*parser.ExpressionContext)

	if expression.LambdaExpression() != nil {
		return NewLambdaNode(expression.LambdaExpression())
	}

	if expression.Primary() != nil {
		return expressionFromPrimary(expression.Primary())
	}

	if expression.COLONCOLON() != nil {
		// method reference
		return NewMethodReference(expression)
	}

	if expression.NEW() != nil {
		return NewConstructorCall(expression.Creator())
	}

	if expression.DOT() != nil {
		if expression.IDENTIFIER() != nil {
			return NewInstanceAttributeReference(expression.IDENTIFIER().GetText(), ExpressionProcessor(expression.Expression(0)))
		}
	}

	if expression.MethodCall() != nil {
		methodCall := expression.MethodCall()
		if expression.DOT() != nil {

			// method call

			// (variable with instance, method call (Identifier =name, expressionList = parameters) )
			if tool.IsNilInterface(expression.Expression(0)) {
				panic("expression with dot and 1 expression, but no identifier (not method call?\n" + expression.GetText() + "\n")
			}

			instance := ExpressionProcessor(expression.Expression(0))

			return NewMethodCall(instance, methodCall)
		} else {
			return NewMethodCall(nil, methodCall)
		}
	}

	if expression.GetPrefix() != nil {
		// prefix operator
		return NewUnaryOperatorNode(true, expression.GetPrefix().GetText(), ExpressionProcessor(expression.Expression(0)))
	}

	if expression.GetPostfix() != nil {
		return NewUnaryOperatorNode(false, expression.GetPostfix().GetText(), ExpressionProcessor(expression.Expression(0)))
	}

	if expression.GetBop() != nil {
		if expression.COLON() != nil {
			left := ExpressionProcessor(expression.Expression(0))
			middle := ExpressionProcessor(expression.Expression(1))
			right := ExpressionProcessor(expression.Expression(2))
			return NewTernaryOperatorNode(expression.GetBop().GetText(), left, middle, right)
		}
		if expression.INSTANCEOF() != nil {
			left := ExpressionProcessor(expression.Expression(0))
			right := NewIdentifierNode(expression.TypeType(0).GetText())
			return NewBinaryOperatorNode("instanceof", left, right)
		}
		left := ExpressionProcessor(expression.Expression(0))
		right := ExpressionProcessor(expression.Expression(1))
		if right == nil {
			panic("missing right for binary: " + expression.GetText() + "\n\n" + expression.ToStringTree(parser.RuleNames, nil))
		}
		return NewBinaryOperatorNode(expression.GetBop().GetText(), left, right)
	}

	if expression.LBRACK() != nil {
		// array index
		left := ExpressionProcessor(expression.Expression(0))
		right := ExpressionProcessor(expression.Expression(1))
		return NewBinaryOperatorNode("[", left, right)
	}

	if len(expression.AllGT())+len(expression.AllLT()) > 0 {
		// shifting binary operator

		left := ExpressionProcessor(expression.Expression(0))
		right := ExpressionProcessor(expression.Expression(1))

		operator := ""
		for _, t := range expression.AllGT() {
			operator += t.GetText()
		}
		for _, t := range expression.AllLT() {
			operator += t.GetText()
		}

		return NewBinaryOperatorNode(operator, left, right)
	}

	if expression.LPAREN() != nil {
		// cast
		left := NewLiteralNode(expression.TypeType(0).GetText())
		right := ExpressionProcessor(expression.Expression(0))
		return NewBinaryOperatorNode("(", left, right)
	}

	panic("unknown expression: " + expression.GetText() + "\n" + expression.ToStringTree(parser.RuleNames, nil))

	// return nil
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
		return ExpressionProcessor(primaryCtx.Expression())
	}

	if primaryCtx.Literal() != nil {
		literal := primaryCtx.Literal().(*parser.LiteralContext)
		return NewLiteralNode(literal.GetText())
	}

	if primaryCtx.CLASS() != nil {
		return NewClassReference(primaryCtx.TypeTypeOrVoid().GetText())
	}

	panic("unknown primary type: " + primary.GetText() + "\n\n" + primary.ToStringTree(parser.RuleNames, nil))
}

func formalParameterListProcessor(formal parser.IFormalParameterListContext) []ExpressionNode {
	if tool.IsNilInterface(formal) {
		return nil
	}
	ctx := formal.(*parser.FormalParameterListContext)

	parameters := make([]ExpressionNode, 0)
	for _, formalParam := range ctx.AllFormalParameter() {
		formalParamCtx := formalParam.(*parser.FormalParameterContext)
		t := formalParamCtx.TypeType().GetText()
		name := formalParamCtx.VariableDeclaratorId().GetText()
		parameters = append(parameters, NewVariableDecl(t, name, nil, false))
	}

	if ctx.LastFormalParameter() != nil {
		lastParameterCtx := ctx.LastFormalParameter().(*parser.LastFormalParameterContext)
		t := lastParameterCtx.TypeType().GetText()
		name := lastParameterCtx.VariableDeclaratorId().GetText()
		parameters = append(parameters, NewVariableDecl(t, name, nil, true))
	}

	return parameters
}
