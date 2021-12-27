package ast

import "github.com/dragonfax/java_converter/trans/node"

/* when a variable is used.
 * For now could be a field or a local variable
 */
type VarRef struct {
	*node.Base
	*BaseMethodScope

	VariableName string
	This         bool
	Super        bool

	// Field or LocalVarDecl (variable or method/constructor parameter)
	VariableDecl node.Node
}

func NewVarRef(name string) *VarRef {
	return &VarRef{Base: node.New(), BaseMethodScope: NewMethodScope(), VariableName: name}
}

func (vr *VarRef) NodeName() string {
	return vr.String()
}

func (vr *VarRef) String() string {
	if !vr.This && !vr.Super && vr.VariableDecl == nil {
		return vr.VariableName + " /* unresolved */"
	}
	return vr.VariableName
}

func (vr *VarRef) Children() []node.Node {
	return nil
}
