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
	Expressions []exp.ExpressionNode
	Arguments   string
	ReturnType  string
	Class       string
}

func NewMethod(class string) *Method {
	return &Method{
		Class:       class,
		Expressions: make([]exp.ExpressionNode, 0),
	}
}

func (m *Method) String() string {
	if m == nil {
		panic("nil method")
	}
	prefix := "     "
	body := ""
	if m.Expressions != nil {
		for _, node := range m.Expressions {
			if node != nil {
				body += prefix + node.String() + "\n"
			}
		}
	}

	return fmt.Sprintf("func (this *%s) %s(%s) *%s{\n%s\n}\n\n", m.Class, m.Name, m.Arguments, m.ReturnType, body)
}
