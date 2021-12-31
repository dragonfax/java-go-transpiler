package ast

import "github.com/dragonfax/java_converter/trans/node"

/* Class reference is when you say Array.class,
 * in order to use the class as an instance value itself
 * for reflection or other nonsense.
 */
type ClassRef struct {
	*node.Base
	*BaseClassScope

	ClassName string

	Class *Class
}

func (cr *ClassRef) Children() []node.Node {
	return nil
}

func NewClassRef(className string) *ClassRef {
	if className == "" {
		panic("no class name")
	}
	return &ClassRef{
		Base:           node.New(),
		BaseClassScope: NewClassScope(),
		ClassName:      className,
	}
}

func (cr *ClassRef) GetType() *Class {
	return RuntimePackage.GetClass("JavaClass")
}

func (cr *ClassRef) String() string {
	if cr.Class == nil {
		return cr.ClassName + ".class /* unresolved */"

	}
	return cr.Class.Name + ".class"
}

func (cr *ClassRef) NodeName() string {
	return cr.String()
}
