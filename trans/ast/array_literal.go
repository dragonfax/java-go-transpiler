package ast

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

type ArrayLiteral struct {
	*node.Base

	TypePath string
	Elements []node.Node
}

func (al *ArrayLiteral) Children() []node.Node {
	return node.ListOfNodesToNodeList(al.Elements)
}

func (al *ArrayLiteral) String() string {
	l := make([]string, 0)
	for _, node := range al.Elements {
		l = append(l, node.String())
	}
	return fmt.Sprintf("[]%s{%s}", al.TypePath, strings.Join(l, ","))
}

func NewArrayLiteral(lit *parser.ArrayInitializerContext) *ArrayLiteral {
	ctx := lit

	if len(ctx.AllVariableInitializer()) == 0 {
		return &ArrayLiteral{Base: node.New()}
	}

	l := make([]node.Node, 0)
	for _, varInit := range ctx.AllVariableInitializer() {
		varInitCtx := varInit

		exp := variableInitializerProcessor(varInitCtx)

		if exp == nil {
			panic("almost added nil expression to expression list.")
		}

		l = append(l, exp)
	}

	return &ArrayLiteral{
		Base:     node.New(),
		Elements: l,
	}
}
