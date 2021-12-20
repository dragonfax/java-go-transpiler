package ast

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/dragonfax/java_converter/trans/ast/exp"
)

type Class struct {
	Name       string
	BaseClass  string
	Interfaces []exp.TypeNode
	Members    []Member
	Fields     []*Field
	Package    string
}

func (c *Class) OutputFilename() string {
	if c.Package == "" {
		panic("no package for class")
	}
	return fmt.Sprintf("%s/%s/%s.go", strings.ReplaceAll(c.Package, ".", "/"), c.Name, c.Name)
}

var classTemplate = `
{{ $className := .Name }}
{{range .Interfaces }}var _ {{ . }} = &{{ $className}}{}
{{end}}
type {{ .Name }} struct {
    {{if .BaseClass}}*{{ .BaseClass }}{{end}}

    {{range .Fields}}{{ .Declaration }}
	{{end}}
}

func New{{.Name}}() *{{.Name}}{
    this := &{{.Name}}{}

    {{range .Fields}} {{if .HasInitializer}}this.{{ .Initializer }}{{end}}
    {{end}}

    return this
}

{{range .Members}}{{ . }}
{{end}}
`

var classTpl = template.Must(template.New("name").Parse(classTemplate))

func (c *Class) String() string {
	for _, m := range c.Members {
		if m == c {
			panic("recursion")
		}
	}
	b := &strings.Builder{}
	err := classTpl.Execute(b, c)
	if err != nil {
		return err.Error()
	}
	return b.String()
}

func (c *Class) PackageBasename() string {
	last := strings.LastIndex(c.Package, ".")
	return c.Package[last+1 : len(c.Package)]
}

func (c *Class) AsFile() string {
	return fmt.Sprintf("package %s\n\n%s", c.PackageBasename(), c)
}

func NewClass() *Class {
	c := &Class{}
	c.Members = make([]Member, 0)
	c.Interfaces = make([]exp.TypeNode, 0)
	c.Fields = make([]*Field, 0)
	return c
}

type SubClassTODO struct {
	Name string
}

func NewSubClassTODO(name string) *SubClassTODO {
	return &SubClassTODO{Name: name}
}

func (sc *SubClassTODO) String() string {
	return fmt.Sprintf("\n// TODO elevate subclass %s (pre-translation)\n\n", sc.Name)
}
