package ast

import (
	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

type InterfaceMember struct {
	*node.BaseNode

	Name       string
	Arguments  []node.Node
	ReturnType node.Node
}

func (im *InterfaceMember) Children() []node.Node {
	return node.AppendNodeLists(im.Arguments, im.ReturnType)
}

func NewInterface(ctx *parser.InterfaceDeclarationContext) *Class {
	this := NewClass()
	this.Name = ctx.IDENTIFIER().GetText()
	this.Interface = true

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
			BaseNode:   node.NewNode(),
			Name:       declCtx.IDENTIFIER().GetText(),
			Arguments:  FormalParameterListProcessor(declCtx.FormalParameters().FormalParameterList()),
			ReturnType: NewTypeOrVoidNode(declCtx.TypeTypeOrVoid()),
		}

		this.Members = append(this.Members, member)

	}

	return this
}

func (im *InterfaceMember) String() string {
	return im.Name
}

func (im *InterfaceMember) ArgumentsString() string {
	return ArgumentListToString(im.Arguments)
}
