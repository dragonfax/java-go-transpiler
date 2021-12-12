package main

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

type Member interface{}

type BaseMember struct {
	Name string
}

var _ Member = &Constructor{}

type Constructor struct {
	Name        string
	Expressions []OperatorNode
}
