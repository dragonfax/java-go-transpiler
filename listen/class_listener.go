package listen

import (
	"github.com/dragonfax/delver_converter/ast"
	"github.com/dragonfax/delver_converter/exp"
	"github.com/dragonfax/delver_converter/parser"
)

type ClassListener struct {
	*StackableListener
	*parser.BaseJavaParserListener

	lastModifier string

	File *ast.File
}

func NewClassListener(sl *StackListener, file *ast.File, ctx *parser.ClassDeclarationContext) *ClassListener {
	s := &ClassListener{StackableListener: NewStackableListener(sl)}
	s.File = file
	s.File.Class = ast.NewClass()

	s.File.Class.Name = ctx.IDENTIFIER().GetText()

	if ctx.TypeType() != nil {
		s.File.Class.BaseClass = ctx.TypeType().GetText()
	}

	if ctx.TypeList() != nil {
		for _, typeType := range ctx.TypeList().(*parser.TypeListContext).AllTypeType() {
			typeTypeCtx := typeType.(*parser.TypeTypeContext)
			s.File.Class.Interfaces = append(s.File.Class.Interfaces, typeTypeCtx.GetText())
		}
	}

	return s
}

func (s *ClassListener) ExitClassDeclaration(ctx *parser.ClassDeclarationContext) {

	s.Pop(s, s)
}

func (s *ClassListener) EnterConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) {

	name := ctx.IDENTIFIER().GetText()

	c := ast.NewConstructor()
	c.Modifier = s.lastModifier
	c.Name = name

	// ctx.FormalPameters()

	for _, blockChild := range ctx.Block().GetChildren() {
		blockStatementContext, ok := blockChild.(*parser.BlockStatementContext)
		if ok {
			statement := blockStatementContext.Statement().(*parser.StatementContext)

			// TODO only one expression per block? no this isn't complicated enough.
			// but okay for a first of expression parsing
			node := exp.ExpressionProcessor(statement.Expression(0).(*parser.ExpressionContext))
			c.Expressions = append(c.Expressions, node)
		}
	}
	s.File.Class.Members = append(s.File.Class.Members, c)
}

func (s *ClassListener) EnterClassBodyDeclaration(ctx *parser.ClassBodyDeclarationContext) {
	// capture the public/private modifer for each class body

	s.lastModifier = "private" // java default

	for _, modifier := range ctx.AllModifier() {
		modifierText := modifier.GetText()
		if modifierText == "public" || modifierText == "private" {
			// these are all we care about right now.
			// for each class member
			s.lastModifier = modifierText
		}
	}

}

func (s *ClassListener) EnterMethodDeclaration(ctx *parser.MethodDeclarationContext) {

	name := ctx.IDENTIFIER().GetText()

	m := ast.NewMethod(s.File.Class.Name)
	m.Modifier = s.lastModifier
	m.Name = name
	m.Arguments = ctx.FormalParameters().GetText()
	m.ReturnType = ctx.TypeTypeOrVoid().GetText()

	for _, blockChild := range ctx.MethodBody().(*parser.MethodBodyContext).Block().GetChildren() {
		blockStatementContext, ok := blockChild.(*parser.BlockStatementContext)
		if ok {
			statement := blockStatementContext.Statement().(*parser.StatementContext)

			node := exp.StatementProcessor(statement)
			m.Expressions = append(m.Expressions, node)

		}
	}
	s.File.Class.Members = append(s.File.Class.Members, m)
}
