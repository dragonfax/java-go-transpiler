package listen

import (
	"github.com/dragonfax/java_converter/input/ast"
	"github.com/dragonfax/java_converter/input/ast/exp"
	"github.com/dragonfax/java_converter/input/parser"
)

type ClassListener struct {
	*StackableListener
	*parser.BaseJavaParserListener

	lastModifier string

	file  *ast.File
	class *ast.Class
}

func NewClassListener(sl *StackListener, file *ast.File, ctx *parser.ClassDeclarationContext) *ClassListener {
	s := &ClassListener{StackableListener: NewStackableListener(sl)}
	s.file = file
	s.class = ast.NewClass()

	s.class.Name = ctx.IDENTIFIER().GetText()

	if ctx.TypeType() != nil {
		s.class.BaseClass = ctx.TypeType().GetText()
	}

	if ctx.TypeList() != nil {
		for _, typeType := range ctx.TypeList().(*parser.TypeListContext).AllTypeType() {
			typeTypeCtx := typeType.(*parser.TypeTypeContext)
			s.class.Interfaces = append(s.class.Interfaces, typeTypeCtx.GetText())
		}
	}

	s.file.Classes = append(s.file.Classes, s.class)

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

	node := exp.NewBlockNode(ctx.Block().(*parser.BlockContext))
	c.Body = node

	s.class.Members = append(s.class.Members, c)
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

	body := exp.NewBlockNode(ctx.MethodBody().(*parser.MethodBodyContext).Block())
	m := ast.NewMethod(s.lastModifier, name, s.class.Name, ctx.FormalParameters().GetText(), ctx.TypeTypeOrVoid().GetText(), body)

	s.class.Members = append(s.class.Members, m)
}

func (s *ClassListener) EnterFieldDeclaration(ctx *parser.FieldDeclarationContext) {

	s.class.Members = append(s.class.Members, ast.NewFields(ctx)...)

}

func (s *ClassListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	s.Push(s, NewClassListener(s.StackListener, s.file, ctx))
}
