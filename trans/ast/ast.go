package ast

import (
	"fmt"
	"strings"
)

type Node interface {
	String() string
}

type File struct {
	Filename    string
	PackageName string
	Imports     []*Import
	Class       *Class
}

func NewFile() *File {
	f := &File{}
	f.Imports = make([]*Import, 0)
	return f
}

func (f *File) String() string {
	return fmt.Sprintf(`
package main // %s;

%s

%s
`, f.PackageName, strings.Join(NodeListToStringList(f.Imports), "\n"), f.Class)
}

type Member interface {
	String() string
}

type BaseMember struct {
	Name     string
	Modifier string
}

func (bm *BaseMember) SetModifier(modifier string) {
	bm.Modifier = modifier
}

type Import struct {
	Name string
}

func NewImport(name string) *Import {
	return &Import{Name: name}
}

func (i *Import) String() string {
	return fmt.Sprintf("import \"%s\"", i.Name)
}
