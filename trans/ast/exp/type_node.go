package exp

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
)

var primitiveTranslation = map[string]string{
	"float":   "float64",
	"double":  "float64",
	"boolean": "bool",
	"long":    "int64",
}

type TypeNode []*TypeElementNode

func (tn TypeNode) String() string {
	list := make([]string, 0)
	for _, ten := range tn {
		list = append(list, ten.String())
	}
	return strings.Join(list, ".")
}

// used when defining the type of variable/return value/parameter/etc
type TypeElementNode struct {
	Class         string
	TypeArguments []TypeNode
}

func (tn *TypeElementNode) String() string {
	class := tn.Class
	if tClass, ok := primitiveTranslation[class]; ok {
		class = tClass
	}

	if len(tn.TypeArguments) > 0 {

		list := make([]string, 0)
		for _, s := range tn.TypeArguments {
			list = append(list, s.String())
		}

		return fmt.Sprintf("%s[%s]", class, strings.Join(list, ","))
	}

	return class
}

func NewTypeOrVoidNode(typ *parser.TypeTypeOrVoidContext) TypeNode {
	typCtx := typ
	if typCtx.VOID() != nil {
		return TypeNode([]*TypeElementNode{{Class: "void"}})
	} else {
		return NewTypeNode(typCtx.TypeType())
	}
}

func NewTypeNode(typ *parser.TypeTypeContext) TypeNode {
	if typ == nil {
		return nil
	}
	ctx := typ
	if ctx.PrimitiveType() != nil {
		// simple primitive type, easy to parse
		return TypeNode([]*TypeElementNode{
			{
				Class: ctx.PrimitiveType().GetText(),
			},
		})
	}

	classCtx := ctx.ClassOrInterfaceType()

	// multile components to one type. Say Car.WheelEnum
	typeNode := make(TypeNode, 0)
	for i, typID := range classCtx.AllIDENTIFIER() {
		class := typID.GetText()

		typeComp := classCtx.TypeArguments(i)
		if typeComp == nil {
			// no typ arguments for this element of the type
			typeNode = append(typeNode, &TypeElementNode{
				Class: class,
			})
			continue
		}

		typeCompCtx := typeComp

		// (typeArguments) multiple type args between <> seperated by commas
		thisTypeArguments := make([]TypeNode, 0)
		for _, typCompArg := range typeCompCtx.AllTypeArgument() {

			childTypeNode := NewTypeNode(typCompArg.TypeType())
			thisTypeArguments = append(thisTypeArguments, childTypeNode)
		}

		node := &TypeElementNode{
			Class:         class,
			TypeArguments: thisTypeArguments,
		}
		typeNode = append(typeNode, node)
	}

	return typeNode
}
