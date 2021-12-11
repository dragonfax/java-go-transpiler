package main

import "github.com/dragonfax/delver_converter/parser"

type ClassListener struct {
	*StackableListener
	*parser.BaseJavaParserListener

	file *File
}

func NewClassListener(sl *StackListener, file *File) *ClassListener {
	s := &ClassListener{StackableListener: NewStackableListener(sl)}
	s.file = file
	s.file.Class = NewClass()
	return s
}

func (s *ClassListener) ExitClassDeclaration(ctx *parser.ClassDeclarationContext) {

	s.file.Class.Name = ctx.IDENTIFIER().GetText()
	if ctx.TypeType() != nil {
		s.file.BaseClass = ctx.TypeType().GetText()
	}
	if ctx.TypeList() != nil {
		s.file.BaseClass = ctx.TypeList().GetText()
	}

	s.Pop(s, s)
}

func (s *ClassListener) EnterMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	s.Push(s, NewMethodListener(s.StackListener, s.file, ctx))
}
