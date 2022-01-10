package ast

import (
	"fmt"

	"github.com/dragonfax/java-go-transpiler/tool"
	"github.com/dragonfax/java-go-transpiler/trans/node"
)

/* Field is a field declaaration,
 * Which is a type and a variable declaration.
 *
 * A variable declaration being a var name and
 * optionally an expression to set its value.
 *
 * Also fields can have various modifiers
 *
 * Appears only outside of a method. Inside a method, such a thing is a LocalVariableDeclaration
 */
type Field struct {
	*BaseExpression
	*BaseClassScope

	Name       string
	Expression Expression // what the value is set to on construction.
	TypePath   *TypePath

	Public    bool
	Transient bool
	Static    bool
}

func NewField(typ *TypePath, name string, expression Expression) *Field {
	return &Field{
		BaseExpression: NewExpression(),
		BaseClassScope: NewClassScope(),
		TypePath:       typ,
		Name:           name,
		Expression:     expression,
	}
}

func (f *Field) GetType() *Class {
	if f.TypePath == nil {
		fmt.Println("warning: no typepath in field (use of var?)")
		return nil
	}
	f.Type = f.TypePath.GetType()
	return f.Type
}

func (f *Field) Children() []node.Node {
	if f.Expression != nil {
		return []node.Node{f.TypePath, f.Expression}
	}
	return []node.Node{f.TypePath}
}

func (f *Field) Declaration() string {
	return fmt.Sprintf("%s %s", f.Name, f.TypePath)
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
