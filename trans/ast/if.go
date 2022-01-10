package ast

import (
	"fmt"

	"github.com/dragonfax/java-go-transpiler/tool"
	"github.com/dragonfax/java-go-transpiler/trans/node"
)

type If struct {
	*node.Base

	Condition node.Node
	Body      node.Node
	Else      node.Node
}

func (in *If) Children() []node.Node {
	list := []node.Node{in.Body}
	if in.Condition != nil {
		list = append(list, in.Condition)
	}
	if in.Else != nil {
		list = append(list, in.Else)
	}
	return list
}

func NewIf(condition, body, els node.Node) *If {
	if tool.IsNilInterface(body) {
		panic("missing body")
	}
	if tool.IsNilInterface(condition) {
		panic("missing condition")
	}
	return &If{
		Base:      node.New(),
		Condition: condition,
		Body:      body,
		Else:      els,
	}
}

func (in *If) String() string {
	if tool.IsNilInterface(in.Else) {
		return fmt.Sprintf("if %s {\n%s}\n", in.Condition, in.Body)
	}
	return fmt.Sprintf("if %s {\n%s} else {\n%s}\n", in.Condition, in.Body, in.Else)
}
