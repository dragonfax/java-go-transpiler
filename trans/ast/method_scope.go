package ast

/* scope is an entity used for variable resolving.
 * a method or constructor
 */

type MethodScope interface {
	SetMethodScope(*Method)
	GetMethodScope() *Method
}

type BaseMethodScope struct {
	MethodScope *Method
}

func NewMethodScope() *BaseMethodScope {
	return &BaseMethodScope{}
}

func (s *BaseMethodScope) SetMethodScope(scope *Method) {
	s.MethodScope = scope
}

func (s *BaseMethodScope) GetMethodScope() *Method {
	return s.MethodScope
}
