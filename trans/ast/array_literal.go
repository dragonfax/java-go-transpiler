package ast

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

type ArrayLiteral struct {
	*BaseExpression

	Elements []Expression
}

func (al *ArrayLiteral) GetType() *Class {
	if len(al.Elements) == 0 {
		return nil
	}
	al.Type = al.Elements[0].GetType()
	return al.Type
}

func (al *ArrayLiteral) Children() []node.Node {

	return node.ListOfNodesToNodeList(al.Elements)
}

func (al *ArrayLiteral) String() string {
	l := make([]string, 0)
	for _, node := range al.Elements {
		l = append(l, node.String())
	}
	return fmt.Sprintf("[]{%s}", strings.Join(l, ","))
}

func NewArrayLiteral(lit *parser.ArrayInitializerContext) *ArrayLiteral {
	ctx := lit

	if len(ctx.AllVariableInitializer()) == 0 {
		return &ArrayLiteral{BaseExpression: NewExpression()}
	}

	l := make([]Expression, 0)
	for _, varInit := range ctx.AllVariableInitializer() {
		varInitCtx := varInit

		exp := variableInitializerProcessor(varInitCtx)

		if exp == nil {
			panic("almost added nil expression to expression list.")
		}

		l = append(l, exp)
	}

	return &ArrayLiteral{
		BaseExpression: NewExpression(),
		Elements:       l,
	}
}
