package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

type VariableDecl struct {
	*node.Base
	*BaseMethodScope

	Type       *Type
	Name       string
	Expression node.Node // for now
}

func (vn *VariableDecl) Children() []node.Node {
	return []node.Node{vn.Type, vn.Expression}
}

func (vn *VariableDecl) String() string {
	if vn.Expression == nil {
		return fmt.Sprintf("var %s %s", vn.Name, vn.Type)
	}
	return fmt.Sprintf("%s := %s", vn.Name, vn.Expression) // we'll assume the type matches the expression.
}

func NewVariableDecl(typ *Type, name string, expression node.Node) *VariableDecl {
	if typ == nil {
		panic(" no variable type")
	}
	if name == "" {
		panic("no variable name")
	}
	return &VariableDecl{Base: node.New(), BaseMethodScope: NewMethodScope(), Type: typ, Name: name, Expression: expression}
}

func NewVariableDeclNodeList(decl *parser.LocalVariableDeclarationContext) []node.Node {

	l := make([]node.Node, 0)

	typ := NewTypeNodeFromContext(decl.TypeType())

	for _, varDecl := range decl.VariableDeclarators().AllVariableDeclarator() {

		varDeclCtx := varDecl

		var exp node.Node
		if varDeclCtx.VariableInitializer() != nil {
			varInitCtx := varDeclCtx.VariableInitializer()
			exp = variableInitializerProcessor(varInitCtx)
		}

		node := NewVariableDecl(typ, varDeclCtx.VariableDeclaratorId().GetText(), exp)

		l = append(l, node)
	}

	return l
}

func variableInitializerProcessor(ctx *parser.VariableInitializerContext) node.Node {
	var exp node.Node
	if ctx.Expression() != nil {
		exp = ExpressionProcessor(ctx.Expression())
	}
	if ctx.ArrayInitializer() != nil {
		exp = NewArrayLiteral(ctx.ArrayInitializer())
	}

	return exp
}
