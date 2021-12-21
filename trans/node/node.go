package node

type Node interface {
	String() string
	Children() []Node
}
