package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/trans/ast/exp"
	"github.com/dragonfax/java_converter/trans/node"
)

type Method struct {
	Name string

	TypeParameters exp.TypeParameterList // implies generic method
	Body           node.Node
	Arguments      []node.Node
	ReturnType     node.Node
	Class          *Class
	Throws         string

	Public       bool
	Abstract     bool
	Static       bool
	Synchronized bool
}

func (m *Method) Children() []node.Node {
	return node.AppendNodeLists([]node.Node{m.Body, m.ReturnType}, m.Arguments...)
}

func NewMethod(name string, arguments []node.Node, returnType node.Node, body node.Node) *Method {
	return &Method{
		Name:       name,
		Arguments:  arguments,
		ReturnType: returnType,
		Body:       body,
	}
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

	body := ""
	if m.Body != nil {
		body = m.Body.String()
	}

	arguments := ""
	if len(m.Arguments) > 0 {
		arguments = exp.ArgumentListToString(m.Arguments)
	}

	throws := ""
	if m.Throws != "" {
		throws = " /* TODO throws " + m.Throws + "*/"
	}

	return fmt.Sprintf("func (this *%s) %s(%s) %s%s{\n%s\n}\n\n", m.Class, m.Name, arguments, m.ReturnType, throws, body)
}
