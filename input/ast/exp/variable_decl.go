package exp

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
)

type VariableDeclNode struct {
	Type       string
	Name       string
	Expression ExpressionNode // for now
}

func (vn *VariableDeclNode) String() string {
	if vn.Expression == nil {
		return fmt.Sprintf("%s %s", vn.Name, vn.Type)
	}
	return fmt.Sprintf("%s := %s", vn.Name, vn.Expression) // we'll assume the type matches the expression.
}

func NewVariableDecl(typ string, name string, expression ExpressionNode) *VariableDeclNode {
	if typ == "" {
		panic(" no variable type")
	}
	if name == "" {
		panic("no variable name")
	}
	return &VariableDeclNode{Type: typ, Name: name, Expression: expression}
}

func NewVariableDeclNodeList(decl *parser.LocalVariableDeclarationContext) []ExpressionNode {

	l := make([]ExpressionNode, 0)

	typ := decl.TypeType().GetText()

	for _, varDecl := range decl.VariableDeclarators().(*parser.VariableDeclaratorsContext).AllVariableDeclarator() {

		varDeclCtx := varDecl.(*parser.VariableDeclaratorContext)

		var exp ExpressionNode
		if varDeclCtx.VariableInitializer() != nil {
			varInitCtx := varDeclCtx.VariableInitializer().(*parser.VariableInitializerContext)
			exp = variableInitializerProcessor(varInitCtx)
		}

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
		exp = NewArrayLiteral(ctx.ArrayInitializer().(*parser.ArrayInitializerContext))
	}

	return exp
}
