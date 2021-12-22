package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/trans/node"
)

/* starts off as just the field name.
 * we won't have the instance expression until after resolving
 */
type FieldReference struct {
	*node.Base

	FieldName          string
	InstanceExpression node.Node
}

func (ia *FieldReference) Children() []node.Node {
	return []node.Node{ia.InstanceExpression}
}

func NewFieldReference(fieldName string) *FieldReference {
	if fieldName == "" {
		panic("no fielde")
	}
	this := &FieldReference{Base: node.New(), FieldName: fieldName}
	return this
}

func (ia *FieldReference) String() string {
	if ia.InstanceExpression != nil {
		return fmt.Sprintf("%s.%s", ia.InstanceExpression, ia.FieldName)
	}
	return fmt.Sprintf("%s", ia.FieldName)
}
