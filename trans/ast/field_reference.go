package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/node"
)

type InstanceAttributeReference struct {
	Attribute         string
	InstanceReference node.Node
}

func (ia *InstanceAttributeReference) Children() []node.Node {
	return []node.Node{ia.InstanceReference}
}

func NewInstanceAttributeReference(attribute string, instanceExpression node.Node) *InstanceAttributeReference {
	if attribute == "" {
		panic("no attribute")
	}
	if tool.IsNilInterface(instanceExpression) {
		panic("no instance")
	}
	this := &InstanceAttributeReference{Attribute: attribute, InstanceReference: instanceExpression}
	return this
}

func (ia *InstanceAttributeReference) String() string {
	return fmt.Sprintf("%s.%s", ia.InstanceReference, ia.Attribute)
}
