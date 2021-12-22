package ast

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

var primitiveTranslation = map[string]string{
	"float":   "float64",
	"double":  "float64",
	"boolean": "bool",
	"long":    "int64",
}

type Type struct {
	*node.Base

	Elements []*TypeElement
}

func NewType(elements []*TypeElement) *Type {
	return &Type{Base: node.New(), Elements: elements}
}

func NewTypeOrVoid(typ *parser.TypeTypeOrVoidContext) *Type {
	typCtx := typ
	if typCtx.VOID() != nil {
		return NewType([]*TypeElement{{Base: node.New(), Class: "void"}})
	} else {
		return NewTypeNodeFromContext(typCtx.TypeType())
	}
}

func NewTypeNodeFromContext(typ *parser.TypeTypeContext) *Type {
	if typ == nil {
		return nil
	}
	ctx := typ
	if ctx.PrimitiveType() != nil {
		// simple primitive type, easy to parse
		return NewType([]*TypeElement{
			{
				Class: ctx.PrimitiveType().GetText(),
			},
		})
	}

	classCtx := ctx.ClassOrInterfaceType()

	// multile components to one type. Say Car.WheelEnum
	elements := make([]*TypeElement, 0)
	for i, typID := range classCtx.AllIDENTIFIER() {
		class := typID.GetText()

		typeComp := classCtx.TypeArguments(i)
		if typeComp == nil {
			// no typ arguments for this element of the type
			elements = append(elements, &TypeElement{
				Class: class,
			})
			continue
		}

		typeCompCtx := typeComp

		// (typeArguments) multiple type args between <> seperated by commas
		thisTypeArguments := make([]*Type, 0)
		for _, typCompArg := range typeCompCtx.AllTypeArgument() {

			childType := NewTypeNodeFromContext(typCompArg.TypeType())
			thisTypeArguments = append(thisTypeArguments, childType)
		}

		node := &TypeElement{
			Base:          node.New(),
			Class:         class,
			TypeArguments: thisTypeArguments,
		}
		elements = append(elements, node)
	}

	return NewType(elements)
}

func (tn Type) String() string {
	list := make([]string, 0)
	for _, ten := range tn.Elements {
		list = append(list, ten.String())
	}
	return strings.Join(list, ".")
}

func (tn *Type) Children() []node.Node {
	if tn == nil {
		return nil
	}
	return node.ListOfNodesToNodeList(tn.Elements)
}

// used when defining the type of variable/return value/parameter/etc
type TypeElement struct {
	*node.Base
	Class         string
	TypeArguments []*Type
}

func (te *TypeElement) Children() []node.Node {
	return node.ListOfNodesToNodeList(te.TypeArguments)
}

func (tn *TypeElement) String() string {
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
