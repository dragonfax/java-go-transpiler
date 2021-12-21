package exp

import "github.com/dragonfax/java_converter/trans/node"

type EnumRef struct {
	Name string
}

func (er *EnumRef) Children() []node.Node {
	return nil
}

func NewEnumRef(name string) *EnumRef {
	return &EnumRef{Name: name}
}

func (er *EnumRef) String() string {
	return er.Name
}
