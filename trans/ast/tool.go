package ast

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/dragonfax/java-go-transpiler/tool"
	"github.com/dragonfax/java-go-transpiler/trans/node"
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

func DebugPrint(n node.Node) {
	fmt.Println(BuildParentBreadcrumbs(n))
	fmt.Println(node.JSONMarshalNode(n))
	if cs, ok := n.(ClassScope); ok && cs.GetClassScope() != nil {
		fmt.Printf("from class %s\n", cs.GetClassScope().Name)
	}
	if ms, ok := n.(MethodScope); ok && ms.GetMethodScope() != nil {
		methodScope := ms.GetMethodScope()
		fmt.Printf("from class %s and method %s\n", methodScope.GetClassScope().Name, methodScope.Name)
	}
}

func BuildParentBreadcrumbs(n node.Node) string {
	s := n.String()
	parent := n.GetParent()
	if parent != nil {
		s = BuildParentBreadcrumbs(parent) + "::" + s
	}
	return s
}
