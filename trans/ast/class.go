package ast

import (
	"fmt"
	"text/template"

	"github.com/dragonfax/java_converter/trans/node"
)

var PrimitiveClasses = map[string]*Class{
	"float":   NewPrimitiveClass("float"),
	"double":  NewPrimitiveClass("double"),
	"boolean": NewPrimitiveClass("boolean"),
	"long":    NewPrimitiveClass("long"),
	"String":  NewPrimitiveClass("String"),
	"void":    NewPrimitiveClass("void"),
	"int":     NewPrimitiveClass("int"),
}

type Class struct {
	*node.Base

	Name          string
	Imports       []*Import
	BaseClassName string
	BaseClass     *Class
	Interfaces    []*TypePath

	/* could be a method, but could also be an interface, a nested class, or a few other things */
	Methods       []*Method
	NestedClasses []*Class

	Fields       []*Field
	PackageScope *Package
	PackageName  string
	Interface    bool
	Enum         bool
	Constants    []*EnumConstant // for enums
	Generated    bool

	FieldsByName map[string]*Field
}

func (c *Class) NodeName() string {
	if c.BaseClassName != "" {
		if c.BaseClass == nil {
			return fmt.Sprintf("%s (%s /* unresolved */)", c.Name, c.BaseClassName)
		}
		return fmt.Sprintf("%s (%s)", c.Name, c.BaseClass.Name)
	}
	return c.Name
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
	list = append(list, node.ListOfNodesToNodeList(c.Methods)...)
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

{{range .NestedClasses}}
// TODO other method in class
{{printf "%T" . }}
{{end}}

{{range .Methods}}{{ . }}
{{end}}
`

var classTpl = template.Must(template.New("name").Parse(classTemplate))

var interfaceTemplate = `
type {{ .Name }} interface {
{{range .Methods}}
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
		Base:          node.New(),
		Methods:       make([]*Method, 0),
		NestedClasses: make([]*Class, 0),
		Interfaces:    make([]*TypePath, 0),
		Fields:        make([]*Field, 0),
		Imports:       make([]*Import, 0),
		Constants:     make([]*EnumConstant, 0),
		FieldsByName:  make(map[string]*Field),
	}
	return c
}

func NewPrimitiveClass(name string) *Class {
	c := NewClass()
	c.Name = name
	return c
}

func (c *Class) AddField(field *Field) {
	c.FieldsByName[field.Name] = field
}

func (c *Class) NestedClassByName(className string) *Class {
	for _, nc := range c.NestedClasses {
		if nc.Name == className {
			return nc
		}
	}
	return nil
}

func (thisClass *Class) ResolveClassName(className string) *Class {
	// resolve a classname to a another class from the scope of this class.
	// class could come from the same package as this class, or an imported class, or an imported package
	// TODO or a primitive boxing class

	// primitives
	if pClass, ok := PrimitiveClasses[className]; ok {
		return pClass
	}

	// primitive boxes (via runtime package)
	if _, ok := boxingClassesSet[className]; ok {
		return thisClass.PackageScope.GetParent().(*Hierarchy).GetPackage("runtime").GetClass(className)
	}

	// nested classes
	nc := thisClass.NestedClassByName(className)
	if nc != nil {
		return nc
	}

	// classes imported directly
	for _, imp := range thisClass.Imports {
		if imp.ImportClass != nil && imp.ImportClass.Name == className {
			// the referenced class is from this import.
			return imp.ImportClass
		}
	}

	// classes in the same package as this one
	resolvedClass := thisClass.PackageScope.HasClass(className)
	if resolvedClass != nil {
		// the references class is from the same package
		return resolvedClass
	}

	// classes imported indirectly via '*' package imports
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
