package ast

import "github.com/dragonfax/java_converter/trans/node"

/* Refers to something by name.
 * But until until we resolve what, it could be any of
 * - field on any class
 * - local variable in the current method
 * - parameter of the current method
 * - name of a class thats imported.
 *
 * could be by itself as an expression.
 * could be the far left, or middle component of a dereferenc chain (DOT operators)
 * is never on the far right of a DOT chain.
 */
type IdentRef struct {
	*node.Base
	*BaseMethodScope

	Name  string
	This  bool // a reference to the current class (never on the right of a DOT)
	Super bool // a reference to the super of the current class (never on the right of a dot)

	// depending on what it refers to.
	LocalVariableDecl *LocalVarDecl // local var or method parameter
	Class             *Class        // a class reference ("Float.MAX", not "Float.class" )
	Field             *Field        // a field on a class.
}

func NewIdentRef(name string) *IdentRef {
	return &IdentRef{Base: node.New(), BaseMethodScope: NewMethodScope(), Name: name}
}

func (vr *IdentRef) GetType() *Class {
	if vr.LocalVariableDecl != nil {
		return vr.LocalVariableDecl.GetType()
	}
	if vr.Class != nil {
		return RuntimePackage.GetClass("JavaClass")
	}
	if vr.Field != nil {
		return vr.Field.GetType()
	}
	return nil
}

func (ir *IdentRef) IsResolved() bool {
	return !(ir.LocalVariableDecl == nil && ir.Class == nil && ir.Field == nil)
}

func (ir *IdentRef) NodeName() string {
	return ir.String()
}

func (ir *IdentRef) String() string {
	if !ir.IsResolved() {
		return ir.Name + " /* unresolved */"
	}
	if ir.Class != nil {
		return ir.Class.Name
	}
	if ir.LocalVariableDecl != nil {
		return ir.LocalVariableDecl.Name
	}
	if ir.Field != nil {
		return ir.Field.Name
	}
	panic("no solution")
}

func (vr *IdentRef) Children() []node.Node {
	return nil
}
