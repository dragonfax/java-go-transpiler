package ast

import "github.com/dragonfax/java_converter/trans/node"

type Variable struct {
	*node.Base

	Name string
}

func NewVariable(name string) *Variable {
	if name == "" {
		panic("missing name")
	}
	return &Variable{
		Base: node.New(),
		Name: name,
	}
}

func (vn *Variable) Children() []node.Node {
	return nil
}

func (vn *Variable) String() string {
	return vn.Name
}
