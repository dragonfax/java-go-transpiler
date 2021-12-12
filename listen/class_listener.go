package main

import "github.com/dragonfax/delver_converter/parser"

type ClassListener struct {
	*StackableListener
	*parser.BaseJavaParserListener

	File *File
}

func NewClassListener(sl *StackListener, file *File, ctx *parser.ClassDeclarationContext) *ClassListener {
	s := &ClassListener{StackableListener: NewStackableListener(sl)}
	s.File = file
	s.File.Class = NewClass()

	s.File.Class.Name = ctx.IDENTIFIER().GetText()

	if ctx.TypeType() != nil {
		s.File.Class.BaseClass = ctx.TypeType().GetText()
	}
	/*
		if ctx.TypeList() != nil {
			s.file.Class.Interfaces = ctx.TypeList().
		}
	*/

	return s
}

func (s *ClassListener) ExitClassDeclaration(ctx *parser.ClassDeclarationContext) {

	s.Pop(s, s)
}

func (s *ClassListener) EnterConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) {

	name := ctx.IDENTIFIER().GetText()

	c := &Constructor{
		Expressions: make([]OperatorNode, 0),
	}
	c.Name = name

	// ctx.FormalPameters()

	for _, blockChild := range ctx.Block().GetChildren() {
		blockStatementContext, ok := blockChild.(*parser.BlockStatementContext)
		if ok {
			statement := blockStatementContext.Statement().(*parser.StatementContext)

			// TODO only one expression per block? no this isn't complicated enough.
			// but okay for a first of expression parsing
			node := expressionProcessor(statement.Expression(0).(*parser.ExpressionContext))
			c.Expressions = append(c.Expressions, node)
		}
	}
	s.File.Class.Members = append(s.File.Class.Members, c)
}
