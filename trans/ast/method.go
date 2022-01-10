package ast

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/dragonfax/java-go-transpiler/input/parser"
	"github.com/dragonfax/java-go-transpiler/tool"
	"github.com/dragonfax/java-go-transpiler/trans/node"
)

/* Method, a method or a constructor.
 */
type Method struct {
	*node.Base
	*BaseClassScope

	Name           string
	TypeParameters *TypeParameterList // implies generic method, nullable
	Arguments      []*LocalVarDecl    // nullable
	ReturnType     *TypePath          // nullable
	Body           node.Node          // nullable
	Throws         string

	Public       bool
	Abstract     bool
	Static       bool
	Synchronized bool
	Constructor  bool
	Interface    bool
	Initializer  bool

	LocalVars map[string]*LocalVarDecl
}

/* Not an expression, but has a GetType anyways.
 * The expression version is MethodCall
 */
func (m *Method) GetType() *Class {
	return m.ReturnType.GetType()
}

func NewConstructor(ctx *parser.ConstructorDeclarationContext) *Method {

	c := &Method{
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

func NewMethod(name string, arguments []*LocalVarDecl, returnType *TypePath, body node.Node) *Method {
	return &Method{
		Base:           node.New(),
		BaseClassScope: NewClassScope(),

		Name:       name,
		Arguments:  arguments,
		ReturnType: returnType,
		Body:       body,
		LocalVars:  make(map[string]*LocalVarDecl),
	}
}

func (m *Method) Children() []node.Node {
	list := make([]node.Node, 0)

	list = append(list, node.ListOfNodesToNodeList(m.Arguments)...)
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

func (m *Method) SetPublic(public bool) {
	m.Public = public
}

func (m *Method) SetAbstract(abstract bool) {
	m.Abstract = abstract
}

func (m *Method) SetStatic(static bool) {
	m.Static = static
}

func (m *Method) SetSynchronized(sync bool) {
	m.Synchronized = sync
}

func (m *Method) String() string {
	if m == nil {
		panic("nil method")
	}
	if m.Constructor {
		return m.ConstructorString()
	}
	return m.MethodString()
}
func (m *Method) MethodString() string {
	if m == nil {
		panic("nil method")
	}

	body := ""
	if !tool.IsNilInterface(m.Body) {
		body = m.Body.String()
	}

	arguments := ""
	if len(m.Arguments) > 0 {
		arguments = ArgumentListToString(node.ListOfNodesToNodeList(m.Arguments))
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

func (im *Method) ArgumentsString() string {
	return ArgumentListToString(node.ListOfNodesToNodeList(im.Arguments))
}

func ArgumentCount(method *Method) string {
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

func (c *Method) ConstructorString() string {
	writer := strings.Builder{}
	constructorTemplateCompiled.Execute(&writer, c)
	return writer.String()
}

func (m *Method) AddLocalVar(localVarDecl *LocalVarDecl) {
	if m == nil {
		panic("nil method")
	}
	m.LocalVars[localVarDecl.Name] = localVarDecl
}
