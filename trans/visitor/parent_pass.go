package visitor

import (
	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/node"
)

/* Set the parent of each node */
func ParentPass(n node.Node) {
	children := n.Children()
	for _, child := range children {
		if tool.IsNilInterface(child) {
			continue
		}

		child.SetParent(n)
		ParentPass(child)
	}
}
