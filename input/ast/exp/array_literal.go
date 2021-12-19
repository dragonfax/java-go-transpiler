package exp

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
)

type ArrayLiteral struct {
	Type     string
	Elements []ExpressionNode
}

func (al *ArrayLiteral) String() string {
	l := make([]string, 0)
	for _, node := range al.Elements {
		l = append(l, node.String())
	}
	return fmt.Sprintf("[]%s{%s}", al.Type, strings.Join(l, ","))
}

func NewArrayLiteral(lit *parser.ArrayInitializerContext) *ArrayLiteral {
	ctx := lit

	if len(ctx.AllVariableInitializer()) == 0 {
		return &ArrayLiteral{}
	}

	l := make([]ExpressionNode, 0)
	for _, varInit := range ctx.AllVariableInitializer() {
		varInitCtx := varInit

		exp := variableInitializerProcessor(varInitCtx)

		if exp == nil {
			panic("almost added nil expression to expression list.")
		}

		l = append(l, exp)
	}

	return &ArrayLiteral{
		Elements: l,
	}
}
