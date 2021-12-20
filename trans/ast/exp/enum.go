package exp

type EnumRef struct {
	Name string
}

func NewEnumRef(name string) *EnumRef {
	return &EnumRef{Name: name}
}

func (er *EnumRef) String() string {
	return er.Name
}
