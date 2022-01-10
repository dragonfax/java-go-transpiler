package ast

import "github.com/dragonfax/java-go-transpiler/trans/node"

type Expression interface {
	node.Node

	GetType() *Class
}

type BaseExpression struct {
	*node.Base

	Type *Class // cached return type of the expression.
}

func NewExpression() *BaseExpression {
	return &BaseExpression{Base: node.New()}
}
