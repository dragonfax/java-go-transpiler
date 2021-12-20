package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/trans/ast/exp"
)

type Method struct {
	*BaseMember

	Body       exp.ExpressionNode
	Arguments  []exp.ExpressionNode
	ReturnType exp.ExpressionNode
	Class      string
	Throws     string

	Public       bool
	Abstract     bool
	Static       bool
	Synchronized bool
}

func NewMethod(name string, class string, arguments []exp.ExpressionNode, returnType exp.ExpressionNode, body exp.ExpressionNode) *Method {
	return &Method{
		BaseMember: NewBaseMember(name),
		Class:      class,
		Arguments:  arguments,
		ReturnType: returnType,
		Body:       body,
	}
}

func (m *Method) SetClass(class string) {
	m.Class = class
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
