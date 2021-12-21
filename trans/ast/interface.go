package ast

import (
	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/ast/exp"
)

type InterfaceMember struct {
	Name       string
	Arguments  []exp.ExpressionNode
	ReturnType exp.ExpressionNode
}

func NewInterface(ctx *parser.InterfaceDeclarationContext) *Class {
	this := &Class{
		Name:      ctx.IDENTIFIER().GetText(),
		Members:   make([]Member, 0),
		Interface: true,
	}

	if ctx.EXTENDS() != nil {
		panic("interface extending another interface")
	}

	for _, decl := range ctx.InterfaceBody().AllInterfaceBodyDeclaration() {
		// NOTE: all interface members are assumed public

		declCtx := decl.InterfaceMemberDeclaration().InterfaceMethodDeclaration()
		if declCtx == nil {
			panic("some unsupported type of member inside interface declaration")
		}

		member := &InterfaceMember{
			Name:       declCtx.IDENTIFIER().GetText(),
			Arguments:  exp.FormalParameterListProcessor(declCtx.FormalParameters().FormalParameterList()),
			ReturnType: exp.NewTypeOrVoidNode(declCtx.TypeTypeOrVoid()),
		}

		this.Members = append(this.Members, member)

	}

	return this
}

func (im *InterfaceMember) String() string {
	return im.Name
}

func (im *InterfaceMember) ArgumentsString() string {
	return exp.ArgumentListToString(im.Arguments)
}
