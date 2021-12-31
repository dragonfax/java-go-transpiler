package ast

import (
	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

/* FieldList is a temporary node.
 * To hold a list of fields with the same type that were defined
 * in the same class method statement.
 *
 * As soon as th class gets the node, it breaks it into its individual fields,
 * and adds them all to itself.
 */

type FieldList struct {
	*node.Base

	Fields []*Field
}

func NewFieldList(ctx *parser.FieldDeclarationContext) *FieldList {
	methods := make([]*Field, 0)

	typ := NewTypeNodeFromContext(ctx.TypeType())

	for _, varDec := range ctx.VariableDeclarators().AllVariableDeclarator() {
		varDecCtx := varDec

		name := varDecCtx.VariableDeclaratorId().GetText()

		var init Expression
		if varDecCtx.VariableInitializer() != nil {
			initCtx := varDecCtx.VariableInitializer()
			if initCtx.Expression() != nil {
				init = ExpressionProcessor(initCtx.Expression())
			} else if initCtx.ArrayInitializer() != nil {
				init = NewArrayLiteral(initCtx.ArrayInitializer())
			}
		}

		methods = append(methods, NewField(typ, name, init))
	}

	return &FieldList{Base: node.New(), Fields: methods}
}

func (fl *FieldList) Children() []node.Node {
	panic("shouldn't reach this point.")
}

func (fl *FieldList) String() string {
	panic("shouldn't reach this point.")
}

func (m *FieldList) SetPublic(public bool) {
	for _, f := range m.Fields {
		f.SetPublic(public)
	}
}

func (m *FieldList) SetStatic(static bool) {
	for _, f := range m.Fields {
		f.SetStatic(static)
	}
}

func (m *FieldList) SetTransient(transient bool) {
	for _, f := range m.Fields {
		f.SetTransient(transient)
	}
}
