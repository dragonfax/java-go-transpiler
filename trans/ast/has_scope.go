package ast

import (
	"github.com/dragonfax/java_converter/trans/node"
)

/* scope is an entity used for variable resolving.
 * a method or constructo
 */

type Scope interface {
	SetScope(node.Node)
}

type BaseScope struct {
	Scope node.Node // method or constructor for now.
}

func NewScope() *BaseScope {
	return &BaseScope{}
}

func (s *BaseScope) SetScope(scope node.Node) {
	s.Scope = scope
}
