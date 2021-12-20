package exp

type UnimplementedNode struct {
	Msg string
}

func (un *UnimplementedNode) String() string {
	return un.Msg
}

func NewUnimplementedNode(msg string) *UnimplementedNode {
	return &UnimplementedNode{Msg: msg}
}
