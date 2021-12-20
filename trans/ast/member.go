package ast

type Member interface {
	String() string
}

type BaseMember struct {
	Name string
}

func NewBaseMember(name string) *BaseMember {
	return &BaseMember{Name: name}
}
