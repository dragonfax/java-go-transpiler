package ast

/* scope is an entity used for variable resolving.
 * a method or constructo
 */

type ClassScope interface {
	SetClassScope(*Class)
}

type BaseClassScope struct {
	ClassScope *Class // method or constructor, for now.
}

func NewClassScope() *BaseClassScope {
	return &BaseClassScope{}
}

func (s *BaseClassScope) SetClassScope(scope *Class) {
	s.ClassScope = scope
}
