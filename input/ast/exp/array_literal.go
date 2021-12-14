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

func NewArrayLiteral(ctx *parser.ArrayInitializerContext) *ArrayLiteral {

	if len(ctx.AllVariableInitializer()) == 0 {
		return &ArrayLiteral{}
	}

	l := make([]ExpressionNode, 0)
	for _, varInit := range ctx.AllVariableInitializer() {
		varInitCtx := varInit.(*parser.VariableInitializerContext)

		exp := variableInitializerProcessor(varInitCtx)

		l = append(l, exp)
	}

	return &ArrayLiteral{
		Elements: l,
	}
}
