package ast

import (
	"fmt"

	"github.com/dragonfax/java-go-transpiler/trans/node"
)

/* referencing a field of a class, either this class or another */
type FieldRef struct {
	*node.Base

	FieldName          string
	InstanceExpression node.Node

	Field *Field
}

func (fr *FieldRef) GetType() *Class {
	if fr.Field != nil {
		return fr.Field.GetType()
	}
	return nil
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
	if fr.InstanceExpression == nil {
		panic("shouldn't happen")
	}
	name := fr.FieldName
	if fr.Field == nil {
		name += " /* Unresolved */"
	}
	return fmt.Sprintf("%s.%s", fr.InstanceExpression.String(), name)
}
