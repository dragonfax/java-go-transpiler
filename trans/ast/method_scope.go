package ast

import (
	"github.com/dragonfax/java_converter/trans/node"
)

/* scope is an entity used for variable resolving.
 * a method or constructo
 */

type MethodScope interface {
	SetMethodScope(node.Node)
}

type BaseMethodScope struct {
	MethodScope node.Node // method or constructor, for now.
}

func NewMethodScope() *BaseMethodScope {
	return &BaseMethodScope{}
}

func (s *BaseMethodScope) SetMethodScope(scope node.Node) {
	s.MethodScope = scope
}
