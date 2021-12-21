package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

type Constructor struct {
	Name       string
	Body       node.Node
	Parameters []node.Node
	Throws     string
	Class      *Class

	Public bool
}

func (c *Constructor) Children() []node.Node {
	return node.AppendNodeLists(c.Parameters, c.Body)
}

func (c *Constructor) SetPublic(public bool) {
	c.Public = public
}

func NewConstructor(ctx *parser.ConstructorDeclarationContext) *Constructor {

	c := &Constructor{Name: ctx.IDENTIFIER().GetText()}

	if ctx.GetConstructorBody() != nil {
		c.Body = NewBlockNode(ctx.GetConstructorBody())
	}

	if ctx.FormalParameters().FormalParameterList() != nil {
		c.Parameters = FormalParameterListProcessor(ctx.FormalParameters().FormalParameterList())
	}

	if ctx.QualifiedNameList() != nil {
		c.Throws = ctx.QualifiedNameList().GetText()
	}

	return c
}

func (c *Constructor) String() string {
	body := ""
	if c.Body != nil {
		body = c.Body.String()
	}

	throws := ""
	if c.Throws != "" {
		throws = " /* TODO throws " + c.Throws + " */"
	}

	return fmt.Sprintf("func New%s(%s) %s *%s{\n%s\n}\n\n", c.Name, ArgumentListToString(c.Parameters), throws, c.Name, body)
}
