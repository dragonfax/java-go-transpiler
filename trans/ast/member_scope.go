package ast

/* scope is an entity used for variable resolving.
 * a method or constructor
 */

type MemberScope interface {
	SetMemberScope(*Member)
	GetMemberScope() *Member
}

type BaseMemberScope struct {
	MemberScope *Member
}

func NewMemberScope() *BaseMemberScope {
	return &BaseMemberScope{}
}

func (s *BaseMemberScope) SetMemberScope(scope *Member) {
	s.MemberScope = scope
}

func (s *BaseMemberScope) GetMemberScope() *Member {
	return s.MemberScope
}
