package ast

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/dragonfax/java_converter/trans/node"
)

type Class struct {
	*node.Base

	Name          string
	Imports       []*Import
	BaseClassName string
	BaseClass     *Class
	Interfaces    []*Type
	Members       []node.Node
	Fields        []*Field
	PackageScope  *Package
	PackageName   string
	Interface     bool
	Enum          bool
	Constants     []*EnumConstant // for enums
}

func (c *Class) Children() []node.Node {
	return node.AppendNodeLists(
		node.AppendNodeLists(
			node.AppendNodeLists(
				node.ListOfNodesToNodeList(c.Members),
				node.ListOfNodesToNodeList(c.Fields)...),
			node.ListOfNodesToNodeList(c.Constants)...),
		c.Imports...,
	)
}

func (c *Class) OutputFilename() string {
	if c.PackageName == "" {
		panic("no package for class")
	}
	return fmt.Sprintf("%s/%s/%s.go", strings.ReplaceAll(c.PackageName, ".", "/"), c.Name, c.Name)
}

var classTemplate = `
{{ $className := .Name }}
{{range .Interfaces }}var _ {{ . }} = &{{ $className}}{}
{{end}}
type {{ .Name }} struct {
    {{if .BaseClassName}}*{{ .BaseClassName }}{{end}}

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

var interfaceTemplate = `
type {{ .Name }} interface {
{{range .Members}}
	{{.Name}}({{.ArgumentsString}}) .ReturnType
{{end}}
}

`

var interfaceTemplateCompiled = template.Must(template.New("interface").Parse(interfaceTemplate))

func (c *Class) String() string {
	if c.Interface {
		return ExecuteTemplateToString(interfaceTemplateCompiled, c)
	}
	if c.Enum {
		return ExecuteTemplateToString(enumTemplateCompiled, c)
	}

	return ExecuteTemplateToString(classTpl, c)
}

func (c *Class) PackageBasename() string {
	last := strings.LastIndex(c.PackageName, ".")
	return c.PackageName[last+1 : len(c.PackageName)]
}

func (c *Class) AsFile() string {
	return fmt.Sprintf("package %s\n\n%s", c.PackageBasename(), c)
}

func NewClass() *Class {
	c := &Class{
		Base:       node.New(),
		Members:    make([]node.Node, 0),
		Interfaces: make([]*Type, 0),
		Fields:     make([]*Field, 0),
		Imports:    make([]*Import, 0),
		Constants:  make([]*EnumConstant, 0),
	}
	return c
}
