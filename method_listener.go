package main

import "github.com/dragonfax/delver_converter/parser"

type MethodListener struct {
	*parser.BaseJavaParserListener

	file   *File
	method *Member
}

func NewMethodListener(file *File, ctx *parser.MethodDeclarationContext) *MethodListener {
	m := &MethodListener{}
	m.file = file

	m.method = NewMember()
	m.method.Type = ctx.TypeTypeOrVoid().GetText()
	m.method.Name = ctx.IDENTIFIER().GetText()
	m.file.Class.Members = append(m.file.Class.Members, m.method)

	return m
}

func (m *MethodListener) ExitMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	stackListener.Pop()
}

func (m *MethodListener) ExitStatement(ctx *parser.StatementContext) {

	// add statement to body list.
	for _, expI := range ctx.AllExpression() {
		exp := expI.(*parser.ExpressionContext)
		m.method.Body = append(m.method.Body, &CodeLine{Body: exp.GetText()})
	}
}
