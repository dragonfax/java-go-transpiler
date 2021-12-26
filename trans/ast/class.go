package ast

import (
	"fmt"
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

	/* could be a member, but could also be an interface, a nested class, or a few other things */
	Members      []*Member
	OtherMembers []node.Node

	Fields       []*Field
	PackageScope *Package
	PackageName  string
	Interface    bool
	Enum         bool
	Constants    []*EnumConstant // for enums
	Generated    bool

	FieldsByName map[string]*Field
}

func (c *Class) Children() []node.Node {
	list := make([]node.Node, 0)

	list = append(list, node.ListOfNodesToNodeList(c.Imports)...)
	for _, n := range list {
		if n == nil {
			panic("nil in list")
		}
	}
	list = append(list, node.ListOfNodesToNodeList(c.Interfaces)...)
	for _, n := range list {
		if n == nil {
			panic("nil in list")
		}
	}
	list = append(list, node.ListOfNodesToNodeList(c.Members)...)
	for _, n := range list {
		if n == nil {
			panic("nil in list")
		}
	}
	list = append(list, node.ListOfNodesToNodeList(c.Fields)...)
	for _, n := range list {
		if n == nil {
			panic("nil in list")
		}
	}
	list = append(list, node.ListOfNodesToNodeList(c.Constants)...)
	for _, n := range list {
		if n == nil {
			panic("nil in list")
		}
	}

	return list
}

var classTemplate = `

{{if .Generated}}
// TODO This class was not included in the original source, but detected by code accessing it. It will have no implementation
{{end}}

{{range .Imports}}
{{.}}
{{end}}

{{ $className := .Name }}
{{range .Interfaces }}var _ {{ . }} = &{{ $className}}{}
{{end}}
type {{ .Name }} struct {
    {{if .BaseClassName -}}
		{{- if .BaseClass -}} 
			{{- .BaseClass.Reference -}}
		{{- else -}}
			*{{ .BaseClassName }} /* unresolved */
		{{- end -}}
	{{- end}}

    {{range .Fields}}{{ .Declaration }}
	{{end}}
}

{{range .OtherMembers}}
// TODO other member in class
{{printf "%T" . }}
{{end}}

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

// Code is referning to this class by name.
// include the basename of the package its in.
func (c *Class) Reference() string {
	return "*" + c.PackageScope.Basename() + "." + c.Name
}

func (c *Class) Filename() string {
	return c.PackageScope.Dir() + "/" + c.Name + ".go"
}

func (c *Class) AsFile() string {
	return fmt.Sprintf("%s\n\n%s", c.PackageScope, c)
}

func NewClass() *Class {
	c := &Class{
		Base:         node.New(),
		Members:      make([]*Member, 0),
		OtherMembers: make([]node.Node, 0),
		Interfaces:   make([]*Type, 0),
		Fields:       make([]*Field, 0),
		Imports:      make([]*Import, 0),
		Constants:    make([]*EnumConstant, 0),
		FieldsByName: make(map[string]*Field),
	}
	return c
}

func (c *Class) AddField(field *Field) {
	c.FieldsByName[field.Name] = field
}

func (thisClass *Class) ResolveClassName(className string) *Class {
	// resolve a classname to a class from the scope of this class.
	// class could come from the same package as this class, or an imported class, or an imported package

	for _, imp := range thisClass.Imports {
		if imp.ImportClass != nil && imp.ImportClass.Name == className {
			// the referenced class is from this import.
			return imp.ImportClass
		}
	}

	resolvedClass := thisClass.PackageScope.HasClass(className)
	if resolvedClass != nil {
		// the references class is from the same package
		return resolvedClass
	}

	// go through all start imported packages
	for _, imp := range thisClass.Imports {
		if imp.Star {
			resolvedClass := imp.ImportPackage.HasClass(className)
			if resolvedClass != nil {
				// the referenced class game from a star imported package
				return resolvedClass
			}
		}
	}

	return nil
}
