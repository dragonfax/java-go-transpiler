package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

type LocalVarDecl struct {
	*node.Base
	*BaseMethodScope

	Name       string
	Expression node.Node
	Type       *Type
}

func NewLocalVarDecl(typ *Type, name string, expression node.Node) *LocalVarDecl {
	return &LocalVarDecl{
		Base:            node.New(),
		BaseMethodScope: NewMethodScope(),
		Type:            typ,
		Name:            name,
		Expression:      expression,
	}
}

func (f *LocalVarDecl) Children() []node.Node {
	list := []node.Node{}
	if f.Type != nil {
		list = append(list, f.Type)
	}
	if f.Expression != nil {
		list = append(list, f.Expression)
	}
	return list
}

func (f *LocalVarDecl) String() string {
	if f.Expression != nil {
		return fmt.Sprintf("%s := %s", f.Name, f.Expression)
	}
	return fmt.Sprintf("var %s %s", f.Name, f.Type)
}

func NewLocalVarDeclNodeList(decl *parser.LocalVariableDeclarationContext) []*LocalVarDecl {

	typ := NewTypeNodeFromContext(decl.TypeType())

	list := make([]*LocalVarDecl, 0)
	for _, varDecl := range decl.VariableDeclarators().AllVariableDeclarator() {

		varDeclCtx := varDecl

		var exp node.Node
		if varDeclCtx.VariableInitializer() != nil {
			varInitCtx := varDeclCtx.VariableInitializer()
			exp = variableInitializerProcessor(varInitCtx)
		}

		name := varDeclCtx.VariableDeclaratorId().GetText()

		node := NewLocalVarDecl(typ, name, exp)

		list = append(list, node)
	}

	return list
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
