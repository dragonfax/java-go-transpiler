package node

type Node interface {
	String() string
	Children() []Node
	SetParent(Node)
}

type BaseNode struct {
	Parent Node
}

func NewNode() *BaseNode {
	return &BaseNode{}
}

func (bn *BaseNode) SetParent(p Node) {
	bn.Parent = p
}
