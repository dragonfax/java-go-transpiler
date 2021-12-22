package ast

import "github.com/dragonfax/java_converter/trans/node"

/* when a variable is used.
 * For now could be a field or a local variable
 */
type VarRef struct {
	*node.Base

	VariableName string
	This         bool
	Super        bool
}

func NewVarRef(name string) *VarRef {
	return &VarRef{Base: node.New(), VariableName: name}
}

func (vr *VarRef) String() string {
	return vr.VariableName
}

func (vr *VarRef) Children() []node.Node {
	return nil
}
