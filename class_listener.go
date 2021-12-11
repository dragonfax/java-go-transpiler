package main

import "github.com/dragonfax/delver_converter/parser"

type ClassListener struct {
	*parser.BaseJavaParserListener

	file *File
}

func NewClassListener(file *File) *ClassListener {
	s := &ClassListener{}
	s.file = file
	s.file.Class = NewClass()
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

	var candidate interface {
		AnnotationTypeDeclaration() parser.IAnnotationTypeDeclarationContext
		ClassDeclaration() parser.IClassDeclarationContext
		EnumDeclaration() parser.IEnumDeclarationContext
		InterfaceDeclaration() parser.IInterfaceDeclarationContext
	} = ctx
	if candidate.AnnotationTypeDeclaration() != nil {
		// candidate = candidate.AnnotationTypeDeclaration().(*parser.AnnotationTypeDeclarationContext).AnnotationTypeBody().(*parser.AnnotationTypeBodyContext).AnnotationTypeElementDeclaration().(*parser.AnnotationTypeElementDeclarationContext).AnnotationTypeElementRest().(*parser.AnnotationTypeElementRestContext)
		a := candidate.AnnotationTypeDeclaration().(*parser.AnnotationTypeDeclarationContext)
		b := a.AnnotationTypeBody().(*parser.AnnotationTypeBodyContext)
		c := b.AnnotationTypeElementDeclaration(0).(*parser.AnnotationTypeElementDeclarationContext)
		d, _ := c.AnnotationTypeElementRest().(*parser.AnnotationTypeElementRestContext)
		candidate = d
	}

	if candidate.ClassDeclaration() != nil {
		s.file.Class.Name = candidate.ClassDeclaration().(*parser.ClassDeclarationContext).IDENTIFIER().GetText()
	} else if candidate.EnumDeclaration() != nil {
		s.file.Class.Name = candidate.EnumDeclaration().(*parser.EnumDeclarationContext).GetText()
	} else if candidate.InterfaceDeclaration() != nil {
		s.file.Class.Name = candidate.InterfaceDeclaration().(*parser.InterfaceDeclarationContext).GetText()
	}
	stackListener.Pop()
}

func (s *ClassListener) EnterMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	// stackListener.Push(NewMethodListener(s.file))
	stackListener.Push(NewMethodListener(s.file, ctx))
}
