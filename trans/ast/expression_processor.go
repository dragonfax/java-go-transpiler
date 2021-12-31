package ast

import (
	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/tool"
)

// deal with the recursive expression tree.
func ExpressionProcessor(expressionI *parser.ExpressionContext) Expression {
	if tool.IsNilInterface(expressionI) {
		return nil
	}

	expression := expressionI

	if expression.LambdaExpression() != nil {
		return NewLambda(expression.LambdaExpression())
	}

	if expression.Primary() != nil {
		return expressionFromPrimary(expression.Primary())
	}

	if expression.COLONCOLON() != nil {
		// method reference
		return NewMethodRef(expression)
	}

	if expression.NEW() != nil {
		return NewConstructorCall(expression.Creator())
	}

	if expression.MethodCall() != nil {
		methodCall := expression.MethodCall()
		return NewMethodCall(methodCall, nil)
	}

	if expression.GetPrefix() != nil {
		// prefix operator
		return NewUnaryOperator(true, expression.GetPrefix().GetText(), ExpressionProcessor(expression.Expression(0)))
	}

	if expression.GetPostfix() != nil {
		return NewUnaryOperator(false, expression.GetPostfix().GetText(), ExpressionProcessor(expression.Expression(0)))
	}

	if expression.GetBop() != nil {
		if expression.GetBop().GetText() == "." {
			/* There is no regular binary operator for DOT
			 * Its never between 2 simople expressions.
			 * Its always between and expression and something else more specific.
			 */
			firstExpression := ExpressionProcessor(expression.Expression(0))
			if expression.THIS() != nil {
				panic("qualified 'this'. For referencing outer class from inner class")
			}
			if expression.IDENTIFIER() != nil {
				return NewFieldRef(expression.IDENTIFIER().GetText(), firstExpression)
			}
			if expression.MethodCall() != nil {
				return NewMethodCall(expression.MethodCall(), firstExpression)
			}
			if expression.NEW() != nil {
				panic("qualified constructor, for constructing inner class from outer class")
			}
			if expression.SUPER() != nil {
				panic("qualified super")
			}
			if expression.ExplicitGenericInvocation() != nil {
				panic("explicit generic invocation")
			}
			panic("unknown DOT binary operator usage")
		}
		if expression.COLON() != nil {
			left := ExpressionProcessor(expression.Expression(0))
			middle := ExpressionProcessor(expression.Expression(1))
			right := ExpressionProcessor(expression.Expression(2))
			return NewTernaryOperator(expression.GetBop().GetText(), left, middle, right)
		}
		if expression.INSTANCEOF() != nil {
			left := ExpressionProcessor(expression.Expression(0))
			right := NewTypeNodeFromContext(expression.TypeType(0))
			return NewBinaryOperator("instanceof", left, right)
		}

		// just some regular binary operator, between 2 expressions.
		left := ExpressionProcessor(expression.Expression(0))
		right := ExpressionProcessor(expression.Expression(1))
		if right == nil {
			panic("missing right for binary: " + expression.GetText() + "\n\n" + expression.ToStringTree(parser.RuleNames, nil))
		}
		return NewBinaryOperator(expression.GetBop().GetText(), left, right)
	}

	if expression.LBRACK() != nil {
		// array index
		left := ExpressionProcessor(expression.Expression(0))
		right := ExpressionProcessor(expression.Expression(1))
		return NewBinaryOperator("[", left, right)
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

		return NewBinaryOperator(operator, left, right)
	}

	if expression.LPAREN() != nil {
		// cast
		left := NewTypeNodeFromContext(expression.TypeType(0))
		right := ExpressionProcessor(expression.Expression(0))
		return NewBinaryOperator("(", left, right)
	}

	panic("unknown expression: " + expression.GetText() + "\n" + expression.ToStringTree(parser.RuleNames, nil))

	// return nil
}

func expressionFromPrimary(primary *parser.PrimaryContext) Expression {
	/* Turns out "primary" just means a single item expression .
	 * Like a variable reference, or a literal, or a parenthesized expression.
	 * A single atom, if you will. But not necessarily a terminal atom.
	 */
	primaryCtx := primary

	if primaryCtx.IDENTIFIER() != nil {
		// variable reference
		return NewIdentRef(primaryCtx.IDENTIFIER().GetText())
	}

	if primaryCtx.THIS() != nil {
		v := NewIdentRef("this")
		v.This = true
		return v
	}

	if primaryCtx.SUPER() != nil {
		v := NewIdentRef("super")
		v.Super = true
		return v
	}

	if primaryCtx.Expression() != nil {
		// removes parenthesized expressions
		// TODO may want to keep these in the future so I don't have to
		//      recalculate where these are needed.
		return ExpressionProcessor(primaryCtx.Expression())
	}

	if primaryCtx.Literal() != nil {
		literal := primaryCtx.Literal()
		return NewLiteralFromContext(literal)
	}

	if primaryCtx.CLASS() != nil {
		return NewClassRef(primaryCtx.TypeTypeOrVoid().GetText())
	}

	panic("unknown primary type: " + primary.GetText() + "\n\n" + primary.ToStringTree(parser.RuleNames, nil))
}

func FormalParameterListProcessor(formal *parser.FormalParameterListContext) []*LocalVarDecl {
	if tool.IsNilInterface(formal) {
		return nil
	}
	ctx := formal

	parameters := make([]*LocalVarDecl, 0)
	for _, formalParam := range ctx.AllFormalParameter() {
		formalParamCtx := formalParam
		t := NewTypeNodeFromContext(formalParamCtx.TypeType())
		name := formalParamCtx.VariableDeclaratorId().GetText()
		parameters = append(parameters, NewArgument(t, name, false))
	}

	if ctx.LastFormalParameter() != nil {
		lastParameterCtx := ctx.LastFormalParameter()
		t := NewTypeNodeFromContext(lastParameterCtx.TypeType())
		name := lastParameterCtx.VariableDeclaratorId().GetText()
		parameters = append(parameters, NewArgument(t, name, true))
	}

	return parameters
}
