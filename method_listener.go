package main

import "github.com/dragonfax/delver_converter/parser"

type MethodListener struct {
	*parser.BaseJavaParserListener

	file *File
}

func NewMethodListener(file *File, ctx *parser.MethodDeclarationContext) *MethodListener {
	m := &MethodListener{}
	m.file = file

	method := &Member{}

	method.Type = ctx.TypeTypeOrVoid().GetText()
	method.Name = ctx.IDENTIFIER().GetText()

	file.Class.Members = append(file.Class.Members, method)

	return m
}
