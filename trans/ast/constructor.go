package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/ast/exp"
)

var _ Member = &Constructor{}

type Constructor struct {
	*BaseMember
	Body       exp.ExpressionNode
	Parameters []exp.ExpressionNode
	Throws     string

	Public bool
}

func (c *Constructor) SetPublic(public bool) {
	c.Public = public
}

func NewConstructor(ctx *parser.ConstructorDeclarationContext) *Constructor {

	c := &Constructor{BaseMember: NewBaseMember(ctx.IDENTIFIER().GetText())}

	if ctx.GetConstructorBody() != nil {
		c.Body = exp.NewBlockNode(ctx.GetConstructorBody())
	}

	if ctx.FormalParameters().FormalParameterList() != nil {
		c.Parameters = exp.FormalParameterListProcessor(ctx.FormalParameters().FormalParameterList())
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

	return fmt.Sprintf("func New%s(%s) %s *%s{\n%s\n}\n\n", c.Name, exp.ArgumentListToString(c.Parameters), throws, c.Name, body)
}
