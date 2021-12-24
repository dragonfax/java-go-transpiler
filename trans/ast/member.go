package ast

import (
	"fmt"

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

	return fmt.Sprintf("func (this *%s) %s(%s) %s%s{\n%s\n}\n\n", m.ClassScope.Name, m.Name, arguments, returnType, throws, body)
}

func (c *Member) ConstructorString() string {
	body := ""
	if c.Body != nil {
		body = c.Body.String()
	}

	throws := ""
	if c.Throws != "" {
		throws = " /* TODO throws " + c.Throws + " */"
	}

	return fmt.Sprintf("func New%s(%s) %s *%s{\n%s\n}\n\n", c.Name, ArgumentListToString(c.Arguments), throws, c.Name, body)
}

func (m *Member) AddLocalVar(localVarDecl *LocalVarDecl) {
	if m == nil {
		panic("nil member")
	}
	m.LocalVars[localVarDecl.Name] = localVarDecl
}
