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

	FieldName string
}

func (fr *FieldRef) Name() string {
	return fmt.Sprintf("FieldRef = %s", fr.FieldName)
}

func (ia *FieldRef) Children() []node.Node {
	return nil
}

func NewFieldRef(fieldName string) *FieldRef {
	if fieldName == "" {
		panic("no fielde")
	}
	this := &FieldRef{Base: node.New(), FieldName: fieldName}
	return this
}

func (ia *FieldRef) String() string {
	return fmt.Sprintf("%s", ia.FieldName)
}
