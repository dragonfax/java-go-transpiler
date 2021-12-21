package ast

import (
	"text/template"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/ast/exp"
	"github.com/dragonfax/java_converter/trans/node"
)

type EnumConstant struct {
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

func NewEnum(ctx *parser.EnumDeclarationContext, fields FieldList, members []node.Node) *Class {
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

	if ctx.EnumConstants() != nil {
		for _, constant := range ctx.EnumConstants().AllEnumConstant() {
			if constant.ClassBody() != nil {
				panic("enum constant has its own class body")
			}

			if constant.Arguments() != nil {
				panic("enum constant has arguments")
			}

			this.Constants = append(this.Constants, &EnumConstant{
				Name: constant.IDENTIFIER().GetText(),
			})
		}
	}

	return this
}
