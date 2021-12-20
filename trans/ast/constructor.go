package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/ast/exp"
)

var _ Member = &Constructor{}

type Constructor struct {
	*BaseMember
	Body exp.ExpressionNode

	Public bool
}

func (c *Constructor) SetPublic(public bool) {
	c.Public = public
}

func NewConstructor() *Constructor {
	return &Constructor{}
}

func (c *Constructor) String() string {
	if c == nil {
		panic("nil constructor")
	}
	if tool.IsNilInterface(c.Body) {
		panic("nil constructor body")
	}

	return fmt.Sprintf("func New%s() *%s{\n%s\n}\n\n", c.Name, c.Name, c.Body)
}
