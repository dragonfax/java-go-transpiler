package ast

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java-go-transpiler/input/parser"
	"github.com/dragonfax/java-go-transpiler/trans/node"
)

type ArrayLiteral struct {
	*node.Base

	Elements []Expression
}

func (al *ArrayLiteral) GetType() *Class {
	if len(al.Elements) == 0 {
		return nil
	}
	return al.Elements[0].GetType()
}

func (al *ArrayLiteral) Children() []node.Node {

	return node.ListOfNodesToNodeList(al.Elements)
}

func (al *ArrayLiteral) String() string {
	l := make([]string, 0)
	for _, node := range al.Elements {
		l = append(l, node.String())
	}
	typ := al.GetType()
	var t = "/* Unresolved */"
	if typ == nil {
		t = typ.Name
	}
	return fmt.Sprintf("[]%s{%s}", t, strings.Join(l, ","))
}

func NewArrayLiteral(lit *parser.ArrayInitializerContext) *ArrayLiteral {
	ctx := lit

	if len(ctx.AllVariableInitializer()) == 0 {
		return &ArrayLiteral{Base: node.New()}
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
		Base:     node.New(),
		Elements: l,
	}
}
