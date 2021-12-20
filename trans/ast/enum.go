package ast

import (
	"fmt"
	"text/template"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/ast/exp"
)

type Enum struct {
	Name       string
	Interfaces []exp.TypeNode
	Constants  []*EnumConstant
}

type EnumConstant struct {
	Name string
}

var enumTemplate = `
type {{.Name}} int
{{ $name := .Name }}

const (
{{range $index, $element := .Constants}}
	{{ $element.Name }} {{ $name }} = iota{{end}}

)
`

var enumTemplateCompiled = template.Must(template.New("enum").Parse(enumTemplate))

func (e *Enum) String() string {
	return exp.ExecuteTemplateToString(enumTemplateCompiled, e)
}

func NewEnum(ctx *parser.EnumDeclarationContext) *Enum {
	this := &Enum{
		Name:       ctx.IDENTIFIER().GetText(),
		Interfaces: make([]exp.TypeNode, 0),
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

	if ctx.EnumBodyDeclarations() != nil {
		fmt.Println("WARNING: enum has body")
	}

	return this
}
