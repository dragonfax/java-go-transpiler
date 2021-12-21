package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/trans/node"
)

/* scope is an entity used for variable resolving.
 * a method or constructo
 */

type HasScope interface {
	SetScope(node.Node)
}

type BaseHasScope struct {
	Scope node.Node // method or constructor for now.
}

func NewBaseHasScope() *BaseHasScope {
	return &BaseHasScope{}
}

func (s *BaseHasScope) SetScope(scope node.Node) {
	fmt.Println("setting scope")
	s.Scope = scope
}
