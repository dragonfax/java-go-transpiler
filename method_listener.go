package main

import "github.com/dragonfax/delver_converter/parser"

type MethodListener struct {
	*StackableListener
	*parser.BaseJavaParserListener

	file   *File
	method *Member
}

func NewMethodListener(sl *StackListener, file *File, ctx *parser.MethodDeclarationContext) *MethodListener {
	m := &MethodListener{StackableListener: NewStackableListener(sl)}
	m.file = file

	m.method = NewMember()
	m.method.Type = ctx.TypeTypeOrVoid().GetText()
	m.method.Name = ctx.IDENTIFIER().GetText()
	m.file.Class.Members = append(m.file.Class.Members, m.method)

	// add statement to body list.
	if ctx.MethodBody().(*parser.MethodBodyContext).Block() != nil {
		for _, blockStatement := range ctx.MethodBody().(*parser.MethodBodyContext).Block().(*parser.BlockContext).AllBlockStatement() {
			blockStatementCtx := blockStatement.(*parser.BlockStatementContext)
			statement := blockStatementCtx.Statement()
			if statement != nil {
				m.method.Body = append(m.method.Body, &CodeLine{Body: statement.GetText()})
			}
		}
	}

	return m
}

func (m *MethodListener) ExitMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	m.Pop(m, m)
}
