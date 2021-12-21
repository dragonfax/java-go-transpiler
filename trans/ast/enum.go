package ast

import (
	"text/template"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/ast/exp"
)

/*
type Enum struct {
	Name        string
	Interfaces  []exp.TypeNode
	Constants   []*EnumConstant
	BodyWarning bool
}
*/

type EnumConstant struct {
	Name string
}

func (ec *EnumConstant) String() string {
	return ec.Name
}

var enumTemplate = `
{{if .BodyWarning}}// TODO this enum has a body, fix pre-translation{{end}}
type {{.Name}} int
{{ $name := .Name }}

const (
{{range $index, $element := .Constants}}
	{{ $element.Name }} {{ $name }} = iota{{end}}

)
`

var enumTemplateCompiled = template.Must(template.New("enum").Parse(enumTemplate))

func NewEnum(ctx *parser.EnumDeclarationContext, fields FieldList, members []Member) *Class {
	this := &Class{
		Name:       ctx.IDENTIFIER().GetText(),
		Interfaces: make([]exp.TypeNode, 0),
		Members:    members,
		Fields:     fields,
		Enum:       true,
	}

	if ctx.TypeList() != nil {
		for _, typeType := range ctx.TypeList().AllTypeType() {
			this.Interfaces = append(this.Interfaces, exp.NewTypeNode(typeType))
		}
	}

	return this
}
