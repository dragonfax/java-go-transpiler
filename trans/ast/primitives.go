package ast

var primitiveClasses = []*PrimitiveClass{
	NewPrimitiveClass("boolean", "Boolean", "bool"),
	NewPrimitiveClass("void", "Void", ""),
	NewPrimitiveClass("char", "Character", "rune"),
	NewPrimitiveClass("String", "String", "string"),
	NewPrimitiveClass("byte", "Byte", "byte"),
	NewPrimitiveClass("int", "Integer", "int"),
	NewPrimitiveClass("short", "Short", "int"),
	NewPrimitiveClass("long", "Long", "int64"),
	NewPrimitiveClass("float", "Float", "float64"),
	NewPrimitiveClass("double", "Double", "float64"),
}

type PrimitiveClass struct {
	JavaPrimitive string
	GoPrimitive   string
	*Class
}

func NewPrimitiveClass(javaPrimitive, javaClassName, goPrimitive string) *PrimitiveClass {
	this := &PrimitiveClass{
		Class:         NewClass(),
		JavaPrimitive: javaPrimitive,
		GoPrimitive:   goPrimitive,
	}
	this.Name = javaClassName
	return this
}
