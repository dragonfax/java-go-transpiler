package ast

import "github.com/dragonfax/java_converter/trans/node"

type Expression interface {
	node.Node

	GetType() *Class
}

type BaseExpression struct {
	*node.Base

	Type *Class
}

func NewExpression() *BaseExpression {
	return &BaseExpression{Base: node.New()}
}
