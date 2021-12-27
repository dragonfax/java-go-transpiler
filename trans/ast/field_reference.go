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

	FieldName string
}

func (fr *FieldReference) Name() string {
	return fmt.Sprintf("FieldRef = %s", fr.FieldName)
}

func (ia *FieldReference) Children() []node.Node {
	return nil
}

func NewFieldReference(fieldName string) *FieldReference {
	if fieldName == "" {
		panic("no fielde")
	}
	this := &FieldReference{Base: node.New(), FieldName: fieldName}
	return this
}

func (ia *FieldReference) String() string {
	return fmt.Sprintf("%s", ia.FieldName)
}
