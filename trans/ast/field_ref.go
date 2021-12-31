package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/trans/node"
)

/* starts off as just the field name.
 * we won't have the instance expression until after resolving
 */
type FieldRef struct {
	*node.Base

	FieldName          string
	InstanceExpression node.Node
}

func (fr *FieldRef) NodeName() string {
	return fmt.Sprintf("FieldRef = %s", fr.FieldName)
}

func (fr *FieldRef) Children() []node.Node {
	if fr.InstanceExpression != nil {
		return []node.Node{fr.InstanceExpression}
	}
	return nil
}

func NewFieldRef(fieldName string, instance node.Node) *FieldRef {
	if fieldName == "" {
		panic("no fielde")
	}
	this := &FieldRef{
		Base:               node.New(),
		FieldName:          fieldName,
		InstanceExpression: instance,
	}
	return this
}

func (fr *FieldRef) String() string {
	if fr.InstanceExpression != nil {
		return fmt.Sprintf("%s.%s", fr.InstanceExpression.String(), fr.FieldName)

	}
	return fmt.Sprintf("%s", fr.FieldName)
}
