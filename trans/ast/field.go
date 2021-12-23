package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/node"
)

/* Field is a field declaaration,
 * Which is a type and a variable declaration.
 *
 * A variable declaration being a var name and
 * optionally an expression to set its value.
 *
 * Also fields can have various modifiers
 *
 * Appears only outside of a method. Inside, such a thing is a LocalVariableDeclaration
 */
type Field struct {
	*node.Base
	*BaseClassScope

	Name       string
	Expression node.Node // for now
	Type       *Type

	Public    bool
	Transient bool
	Static    bool
}

func NewField(typ *Type, name string, expression node.Node) *Field {
	return &Field{
		Base:           node.New(),
		BaseClassScope: NewClassScope(),
		Type:           typ,
		Name:           name,
		Expression:     expression,
	}
}

func (f *Field) Children() []node.Node {
	if f.Expression != nil {
		return []node.Node{f.Type, f.Expression}
	}
	return []node.Node{f.Type}
}

func (f *Field) Declaration() string {
	return fmt.Sprintf("%s %s", f.Name, f.Type)
}

func (f *Field) HasInitializer() bool {
	return !tool.IsNilInterface(f.Expression)
}

func (f *Field) String() string {
	return f.Initializer()
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
