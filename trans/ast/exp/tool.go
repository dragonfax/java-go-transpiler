package exp

import (
	"runtime/debug"
	"strings"

	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/node"
)

func expressionListToString(list []node.Node) string {
	if list == nil {
		panic("list expression list")
	}
	s := ""
	for _, node := range list {
		if tool.IsNilInterface(node) {
			s += "nil node in expression list\n" + string(debug.Stack())
		} else {
			s += node.String() + "\n"
		}
	}
	return s
}

func ArgumentListToString(list []node.Node) string {
	s := make([]string, 0)
	for _, node := range list {
		if tool.IsNilInterface(node) {
			s = append(s, "nil node in argument list\n"+string(debug.Stack()))
		} else {
			s = append(s, node.String())
		}
	}
	return strings.Join(s, ",")
}
