package ast

import "fmt"

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
	return c
}

type Member interface {
	String() string
}

type BaseMember struct {
	Name string
}

var _ Member = &Constructor{}

type Constructor struct {
	BaseMember
	Expressions []OperatorNode
}

func (c *Constructor) String() string {
	if c == nil {
		return ""
	}

	prefix := "     "
	body := ""
	for _, node := range c.Expressions {
		if node != nil {
			body += prefix + node.String() + "\n"
		}
	}
	return fmt.Sprintf("func New%s() *%s{\n%s\n}\n\n", c.Name, c.Name, body)
}
