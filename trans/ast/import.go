package ast

import "fmt"

type Import struct {
	Name string
}

func NewImport(name string) *Import {
	return &Import{Name: name}
}

func (i *Import) String() string {
	return fmt.Sprintf("import \"%s\"", i.Name)
}
