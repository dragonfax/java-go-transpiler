package ast

import "github.com/dragonfax/java_converter/trans/node"

type Unimplemented struct {
	*node.Base
	Msg string
}

func (un *Unimplemented) Children() []node.Node {
	return nil
}

func (un *Unimplemented) String() string {
	return un.Msg
}

func NewUnimplemented(msg string) *Unimplemented {
	return &Unimplemented{Base: node.New(), Msg: msg}
}
