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

type TypeNode struct {
	*node.BaseNode

	Elements []*TypeElementNode
}

func NewTypeNode(elements []*TypeElementNode) *TypeNode {
	return &TypeNode{BaseNode: node.NewNode(), Elements: elements}
}

func NewTypeOrVoidNode(typ *parser.TypeTypeOrVoidContext) *TypeNode {
	typCtx := typ
	if typCtx.VOID() != nil {
		return NewTypeNode([]*TypeElementNode{{BaseNode: node.NewNode(), Class: "void"}})
	} else {
		return NewTypeNodeFromContext(typCtx.TypeType())
	}
}

func NewTypeNodeFromContext(typ *parser.TypeTypeContext) *TypeNode {
	if typ == nil {
		return nil
	}
	ctx := typ
	if ctx.PrimitiveType() != nil {
		// simple primitive type, easy to parse
		return NewTypeNode([]*TypeElementNode{
			{
				Class: ctx.PrimitiveType().GetText(),
			},
		})
	}

	classCtx := ctx.ClassOrInterfaceType()

	// multile components to one type. Say Car.WheelEnum
	elements := make([]*TypeElementNode, 0)
	for i, typID := range classCtx.AllIDENTIFIER() {
		class := typID.GetText()

		typeComp := classCtx.TypeArguments(i)
		if typeComp == nil {
			// no typ arguments for this element of the type
			elements = append(elements, &TypeElementNode{
				Class: class,
			})
			continue
		}

		typeCompCtx := typeComp

		// (typeArguments) multiple type args between <> seperated by commas
		thisTypeArguments := make([]*TypeNode, 0)
		for _, typCompArg := range typeCompCtx.AllTypeArgument() {

			childTypeNode := NewTypeNodeFromContext(typCompArg.TypeType())
			thisTypeArguments = append(thisTypeArguments, childTypeNode)
		}

		node := &TypeElementNode{
			BaseNode:      node.NewNode(),
			Class:         class,
			TypeArguments: thisTypeArguments,
		}
		elements = append(elements, node)
	}

	return NewTypeNode(elements)
}

func (tn TypeNode) String() string {
	list := make([]string, 0)
	for _, ten := range tn.Elements {
		list = append(list, ten.String())
	}
	return strings.Join(list, ".")
}

func (tn TypeNode) Children() []node.Node {
	return node.ListOfNodesToNodeList(tn.Elements)
}

// used when defining the type of variable/return value/parameter/etc
type TypeElementNode struct {
	*node.BaseNode
	Class         string
	TypeArguments []*TypeNode
}

func (te *TypeElementNode) Children() []node.Node {
	return node.ListOfNodesToNodeList(te.TypeArguments)
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
