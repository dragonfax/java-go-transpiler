package ast

import (
	"github.com/dragonfax/java_converter/input/parser"
)

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

		member := NewMethod(
			declCtx.IDENTIFIER().GetText(),
			FormalParameterListProcessor(declCtx.FormalParameters().FormalParameterList()),
			NewTypeOrVoid(declCtx.TypeTypeOrVoid()),
			nil,
		)

		this.Members = append(this.Members, member)

	}

	return this
}
