package node

type Node interface {
	String() string
	Children() []Node
	SetParent(Node)
}

type Base struct {
	Parent Node
}

func New() *Base {
	return &Base{}
}

func (bn *Base) SetParent(p Node) {
	bn.Parent = p
}
