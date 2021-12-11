package main

type File struct {
	QualifiedPackageName string
	Imports              []string
	Class                *Class
	BaseClass            string
	Filename             string
}

func NewFile() *File {
	f := &File{}
	f.Imports = make([]string, 0)
	return f
}

type Class struct {
	Name    string
	Members []*Member
}

func NewClass() *Class {
	c := &Class{}
	c.Members = make([]*Member, 0)
	return c
}

type Member struct {
	Name      string
	Static    bool
	Output    Type
	Type      string
	Arguments []*Argument
	Body      []*CodeLine
}

func NewMember() *Member {
	m := &Member{}
	m.Body = make([]*CodeLine, 0)
	return m
}

type Type struct {
}

type Argument struct {
}

type CodeLine struct {
	Body string
}
