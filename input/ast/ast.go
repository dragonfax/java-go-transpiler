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
	Class       *Class
}

func NewFile() *File {
	f := &File{}
	f.Imports = make([]string, 0)
	return f
}

type Class struct {
	Name       string
	BaseClass  string
	Interfaces []string
	Members    []Member
}

func NewClass() *Class {
	c := &Class{}
	c.Members = make([]Member, 0)
	c.Interfaces = make([]string, 0)
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
	Arguments  string
	ReturnType string
	Class      string
}

func NewMethod(modifier string, name string, class string, arguments, returnType string, body exp.ExpressionNode) *Method {
	if class == "" {
		panic("no class")
	}
	if returnType == "" {
		panic("no return type")
	}
	if arguments == "" {
		panic("no arguments")
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

	return fmt.Sprintf("func (this *%s) %s(%s) *%s{\n%s\n}\n\n", m.Class, m.Name, m.Arguments, m.ReturnType, m.Body)
}
