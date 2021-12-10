package main

type File struct {
	QualifiedPackageName string
	Imports              []string
	Class                *Class
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

type Member struct {
	Name      string
	Static    bool
	Output    Type
	Arguments []*Argument
	Body      []*CodeLine
}

type Type struct {
}

type Argument struct {
}

type CodeLine struct {
}
