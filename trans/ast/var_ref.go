package ast

import "github.com/dragonfax/java_converter/trans/node"

/* when a variable is used.
 * For now could be a field or a local variable
 */
type VarRef struct {
	*node.Base
	*BaseMemberScope

	VariableName string
	This         bool
	Super        bool

	VariableDecl *LocalVarDecl
}

func NewVarRef(name string) *VarRef {
	return &VarRef{Base: node.New(), BaseMemberScope: NewMemberScope(), VariableName: name}
}

func (vr *VarRef) String() string {
	return vr.VariableName
}

func (vr *VarRef) Children() []node.Node {
	return nil
}
