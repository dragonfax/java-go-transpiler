package ast

import "github.com/dragonfax/java_converter/trans/node"

/* Class reference is when you say Array.class,
 * in order to use the class as an instance value itself
 * for reflection or other nonsense.
 */
type ClassReference struct {
	*node.Base

	ClassName string
}

func (cr *ClassReference) Children() []node.Node {
	return nil
}

func NewClassReference(className string) *ClassReference {
	if className == "" {
		panic("no class name")
	}
	return &ClassReference{Base: node.New(), ClassName: className}
}

func (cr *ClassReference) String() string {
	return cr.ClassName + ".class"
}
