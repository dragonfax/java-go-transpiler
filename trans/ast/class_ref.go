package ast

import "github.com/dragonfax/java_converter/trans/node"

/* Class reference is when you say Array.class,
 * in order to use the class as an instance value itself
 * for reflection or other nonsense.
 */
type ClassRef struct {
	*node.Base

	ClassName string
}

func (cr *ClassRef) Children() []node.Node {
	return nil
}

func NewClassRef(className string) *ClassRef {
	if className == "" {
		panic("no class name")
	}
	return &ClassRef{Base: node.New(), ClassName: className}
}

func (cr *ClassRef) String() string {
	return cr.ClassName + ".class"
}

func (cr *ClassRef) Name() string {
	return cr.String()
}
