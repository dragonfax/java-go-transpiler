package ast

import (
	"text/template"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

type EnumConstant struct {
	*node.Base

	Name string
}

func (ec *EnumConstant) Children() []node.Node {
	return nil
}

func (ec *EnumConstant) String() string {
	return ec.Name
}

var enumTemplate = `
type {{.Name}} int
{{ $name := .Name }}

const (
{{range $index, $element := .Constants}}
	{{ $element.Name }} {{ $name }} = iota{{end}}

)
{{if (or .Members .Fields)}}// TODO not including enum body
{{end}}
`

var enumTemplateCompiled = template.Must(template.New("enum").Parse(enumTemplate))

func NewEnum(ctx *parser.EnumDeclarationContext, fields []*Field, members []node.Node) *Class {
	this := NewClass()
	this.Name = ctx.IDENTIFIER().GetText()
	this.Members = members
	this.Fields = fields
	this.Enum = true

	if ctx.TypeList() != nil {
		for _, typeType := range ctx.TypeList().AllTypeType() {
			this.Interfaces = append(this.Interfaces, NewTypeNodeFromContext(typeType))
		}
	}

	if ctx.EnumConstants() != nil {
		for _, constant := range ctx.EnumConstants().AllEnumConstant() {
			if constant.ClassBody() != nil {
				panic("enum constant has its own class body")
			}

			if constant.Arguments() != nil {
				panic("enum constant has arguments")
			}

			this.Constants = append(this.Constants, &EnumConstant{
				Base: node.New(),
				Name: constant.IDENTIFIER().GetText(),
			})
		}
	}

	return this
}
