package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/node"
)

type Field struct {
	*node.BaseNode
	*BaseClassScope
	*VariableDeclNode

	Public    bool
	Transient bool
	Static    bool
}

func NewField(vardecl *VariableDeclNode) *Field {
	return &Field{BaseNode: node.NewNode(), BaseClassScope: NewClassScope(), VariableDeclNode: vardecl}
}

func (f *Field) Children() []node.Node {
	return nil
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

func (f *Field) SetPublic(public bool) {
	f.Public = public
}

func (f *Field) SetStatic(static bool) {
	f.Static = static
}

func (f *Field) SetTransient(transient bool) {
	f.Transient = transient
}
