package ast

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/node"
)

/* Member, a method or a constructor.
 */
type Member struct {
	*node.Base
	*BaseClassScope

	Name           string
	TypeParameters *TypeParameterList // implies generic method, nullable
	Arguments      []node.Node        // nullable
	ReturnType     node.Node          // nullable
	Body           node.Node          // nullable
	Throws         string

	Public       bool
	Abstract     bool
	Static       bool
	Synchronized bool
	Constructor  bool

	LocalVars map[string]*LocalVarDecl
}

func NewConstructor(ctx *parser.ConstructorDeclarationContext) *Member {

	c := &Member{
		Base:           node.New(),
		BaseClassScope: NewClassScope(),
		Name:           ctx.IDENTIFIER().GetText(),
		Constructor:    true,
		LocalVars:      make(map[string]*LocalVarDecl),
	}

	if ctx.GetConstructorBody() != nil {
		c.Body = NewBlock(ctx.GetConstructorBody())
	}

	if ctx.FormalParameters().FormalParameterList() != nil {
		c.Arguments = FormalParameterListProcessor(ctx.FormalParameters().FormalParameterList())
	}

	if ctx.QualifiedNameList() != nil {
		c.Throws = ctx.QualifiedNameList().GetText()
	}

	return c
}

func NewMethod(name string, arguments []node.Node, returnType node.Node, body node.Node) *Member {
	return &Member{
		Base:           node.New(),
		BaseClassScope: NewClassScope(),

		Name:       name,
		Arguments:  arguments,
		ReturnType: returnType,
		Body:       body,
		LocalVars:  make(map[string]*LocalVarDecl),
	}
}

func (m *Member) Children() []node.Node {
	list := make([]node.Node, 0)

	list = append(list, m.Arguments...)
	if m.TypeParameters != nil {
		list = append(list, m.TypeParameters)
	}
	if m.ReturnType != nil {
		list = append(list, m.ReturnType)
	}
	if m.Body != nil {
		list = append(list, m.Body)
	}

	return list
}

func (m *Member) SetPublic(public bool) {
	m.Public = public
}

func (m *Member) SetAbstract(abstract bool) {
	m.Abstract = abstract
}

func (m *Member) SetStatic(static bool) {
	m.Static = static
}

func (m *Member) SetSynchronized(sync bool) {
	m.Synchronized = sync
}

func (m *Member) String() string {
	if m == nil {
		panic("nil member")
	}
	if m.Constructor {
		return m.ConstructorString()
	}
	return m.MethodString()
}
func (m *Member) MethodString() string {
	if m == nil {
		panic("nil method")
	}

	body := ""
	if !tool.IsNilInterface(m.Body) {
		body = m.Body.String()
	}

	arguments := ""
	if len(m.Arguments) > 0 {
		arguments = ArgumentListToString(m.Arguments)
	}

	throws := ""
	if m.Throws != "" {
		throws = " /* TODO throws " + m.Throws + "*/"
	}

	returnType := ""
	if m.ReturnType != nil {
		returnType = m.ReturnType.String()
	}

	if m.ClassScope == nil {
		panic("no class scope for method")
	}

	return fmt.Sprintf("func (this *%s) %s(%s) %s%s{\n%s\n}\n\n", m.ClassScope.Name, m.Name, arguments, returnType, throws, body)
}

func ArgumentCount(method *Member) string {
	if len(method.Arguments) == 0 {
		return ""
	}
	return fmt.Sprintf("%d", len(method.Arguments))
}

var constructorTemplate = `
func New{{.Name}}{{ ArgumentCount . }}({{ Arguments .Arguments }})
	{{- if .Throws -}} 
		/* TODO throws {{ .Throws }} */
	{{- end -}}

	{{- with .ClassScope -}}
	*{{ .Name }} {

	this := &{{.Name}}{
		{{- if .BaseClass -}}
		{{- .BaseClass.Name}}: {{.BaseClass.PackageScope.Basename}}.New{{.BaseClass.Name}}(),
		{{- end -}}
	}

	{{range .Fields}} {{if .HasInitializer}}this.{{ .Initializer }}{{end}}
	{{end}}

	{{end}}

	{{if .Body}}
	{{- .Body -}}
	{{end}}

	return this
}
`
var constructorTemplateCompiled = template.Must(template.New("constructor").Funcs(map[string]interface{}{
	"Arguments":     ArgumentListToString,
	"ArgumentCount": ArgumentCount,
}).Parse(constructorTemplate))

func (c *Member) ConstructorString() string {
	writer := strings.Builder{}
	constructorTemplateCompiled.Execute(&writer, c)
	return writer.String()
}

func (m *Member) AddLocalVar(localVarDecl *LocalVarDecl) {
	if m == nil {
		panic("nil member")
	}
	m.LocalVars[localVarDecl.Name] = localVarDecl
}
