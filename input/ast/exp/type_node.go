package exp

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
)

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
	if len(tn.TypeArguments) > 0 {

		list := make([]string, 0)
		for _, s := range tn.TypeArguments {
			list = append(list, s.String())
		}

		return fmt.Sprintf("%s[%s]", tn.Class, strings.Join(list, ","))
	}
	return tn.Class
}

func NewTypeNode(typ parser.ITypeTypeContext) TypeNode {
	ctx := typ.(*parser.TypeTypeContext)
	if ctx.PrimitiveType() != nil {
		return TypeNode([]*TypeElementNode{
			&TypeElementNode{
				Class: ctx.PrimitiveType().GetText(),
			},
		})
	}

	classCtx := ctx.ClassOrInterfaceType().(*parser.ClassOrInterfaceTypeContext)

	typeElements := make(TypeNode, 0)
	for i, typID := range classCtx.AllIDENTIFIER() {
		class := typID.GetText()

		typeComp := classCtx.TypeArguments(i)
		if typeComp == nil {
			// no typ arguments for this element of the type
			typeElements = append(typeElements, &TypeElementNode{
				Class: class,
			})
			continue
		}
		typeCompCtx := typeComp.(*parser.TypeArgumentsContext)

		countTypArgs := len(typeCompCtx.AllTypeArgument())

		if countTypArgs > 1 {
			panic("too many sets of type arguments")
		}

		var thisTypeArguments []TypeNode
		if countTypArgs == 1 {
			thisTypeArguments = make([]TypeNode, 0)
			typArgsCtx := classCtx.TypeArguments(0).(*parser.TypeArgumentsContext)
			for _, typArg := range typArgsCtx.AllTypeArgument() {
				typArgCtx := typArg.(*parser.TypeArgumentContext)

				if typArgCtx.QUESTION() != nil {
					panic("unknown")
				}

				thisTypeArguments = append(thisTypeArguments, NewTypeNode(typArgCtx.TypeType()))
			}
		}
		node := &TypeElementNode{
			Class:         class,
			TypeArguments: thisTypeArguments,
		}

		typeElements = append(typeElements, node)
	}

	return typeElements
}
