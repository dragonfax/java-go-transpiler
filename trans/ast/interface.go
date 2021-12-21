package ast

import (
	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/ast/exp"
	"github.com/dragonfax/java_converter/trans/node"
)

type InterfaceMember struct {
	Name       string
	Arguments  []node.Node
	ReturnType node.Node
}

func (im *InterfaceMember) Children() []node.Node {
	return node.AppendNodeLists(im.Arguments, im.ReturnType)
}

func NewInterface(ctx *parser.InterfaceDeclarationContext) *Class {
	this := &Class{
		Name:      ctx.IDENTIFIER().GetText(),
		Members:   make([]node.Node, 0),
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
