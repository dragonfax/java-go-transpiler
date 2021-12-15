package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/ast/exp"
	"github.com/dragonfax/java_converter/tool"
)

type File struct {
	Filename    string
	PackageName string
	Imports     []string
	Classes     []*Class
}

func NewFile() *File {
	f := &File{}
	f.Imports = make([]string, 0)
	return f
}

type Class struct {
	Name       string
	BaseClass  string
	Interfaces []exp.TypeNode
	Members    []Member
	Fields     []*Field
}

func NewClass() *Class {
	c := &Class{}
	c.Members = make([]Member, 0)
	c.Interfaces = make([]exp.TypeNode, 0)
	c.Fields = make([]*Field, 0)
	return c
}

type Member interface {
	String() string
}

type BaseMember struct {
	Name     string
	Modifier string
}

var _ Member = &Constructor{}

type Constructor struct {
	BaseMember
	Body exp.ExpressionNode
}

func NewConstructor() *Constructor {
	return &Constructor{}
}

func (c *Constructor) String() string {
	if c == nil {
		panic("nil constructor")
	}
	if tool.IsNilInterface(c.Body) {
		panic("nil constructor body")
	}

	return fmt.Sprintf("func New%s() *%s{\n%s\n}\n\n", c.Name, c.Name, c.Body)
}

type Method struct {
	BaseMember

	Body       exp.ExpressionNode
	Arguments  []exp.ExpressionNode
	ReturnType string
	Class      string
}

func NewMethod(modifier string, name string, class string, arguments []exp.ExpressionNode, returnType string, body exp.ExpressionNode) *Method {
	if class == "" {
		panic("no class")
	}
	return &Method{
		BaseMember: BaseMember{Modifier: modifier, Name: name},
		Class:      class,
		Arguments:  arguments,
		ReturnType: returnType,
		Body:       body,
	}
}

func (m *Method) String() string {
	if m == nil {
		panic("nil method")
	}

	body := ""
	if m.Body != nil {
		body = m.Body.String()
	}

	arguments := ""
	if len(m.Arguments) > 0 {
		arguments = exp.ArgumentListToString(m.Arguments)
	}

	return fmt.Sprintf("func (this *%s) %s(%s) %s{\n%s\n}\n\n", m.Class, m.Name, arguments, m.ReturnType, body)
}
