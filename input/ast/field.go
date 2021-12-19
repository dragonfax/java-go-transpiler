package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/ast/exp"
	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/tool"
)

type Field struct {
	exp.VariableDeclNode
}

func NewFields(ctx *parser.FieldDeclarationContext) []*Field {
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
