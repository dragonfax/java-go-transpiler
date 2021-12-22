package ast

import "github.com/dragonfax/java_converter/trans/node"

type ClassReference struct {
	*node.Base

	Class string
}

func (cr *ClassReference) Children() []node.Node {
	return nil
}

func NewClassReference(class string) *ClassReference {
	if class == "" {
		panic("no class name")
	}
	return &ClassReference{Base: node.New(), Class: class}
}

func (cr *ClassReference) String() string {
	return cr.Class + ".class"
}