package ast

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/ast/exp"
	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/tool"
)

func NodeListToStringLisT[T Node](list []T) []string {
	s := make([]string, 0, len(list))
	for _, n := range list {
		s = append(s, n.String())
	}
	return s
}

type FieldList []*Field

func (fl FieldList) String() string {
	return strings.Join(NodeListToStringLisT(fl), ",")
}

type Field struct {
	exp.VariableDeclNode
}

func NewFields(ctx *parser.FieldDeclarationContext) FieldList {
	members := make([]*Field, 0)

	typ := exp.NewTypeNode(ctx.TypeType())

	for _, varDec := range ctx.VariableDeclarators().AllVariableDeclarator() {
		varDecCtx := varDec

		name := varDecCtx.VariableDeclaratorId().GetText()

		var init exp.ExpressionNode
		if varDecCtx.VariableInitializer() != nil {
			initCtx := varDecCtx.VariableInitializer()
			if initCtx.Expression() != nil {
				init = exp.ExpressionProcessor(initCtx.Expression())
			} else if initCtx.ArrayInitializer() != nil {
				init = exp.NewArrayLiteral(initCtx.ArrayInitializer())
			}
		}

		node := exp.NewVariableDecl(typ, name, init)
		members = append(members, &Field{VariableDeclNode: *node})
	}

	return members
}

func (f *Field) Declaration() string {
	return fmt.Sprintf("%s %s", f.Name, f.Type)
}

func (f *Field) HasInitializer() bool {
	return !tool.IsNilInterface(f.Expression)
}

func (f *Field) Initializer() string {
	return fmt.Sprintf("%s = %s", f.Name, f.Expression)
}
