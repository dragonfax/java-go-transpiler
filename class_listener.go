package main

import "github.com/dragonfax/delver_converter/parser"

type ClassListener struct {
	*parser.BaseJavaParserListener

	file *File
}

func NewClassListener(file *File) *ClassListener {
	s := &ClassListener{}
	s.file = file
	return s
}

func (s *ClassListener) EnterTypeType(ctx *parser.TypeTypeContext) {
	if s.file.BaseClass == "" {
		if ctx.ClassOrInterfaceType() != nil {
			s.file.BaseClass = ctx.ClassOrInterfaceType().GetText()
		}
	}
}

func (s *ClassListener) ExitTypeDeclaration(ctx *parser.TypeDeclarationContext) {
	stackListener.Pop()
}
