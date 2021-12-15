package exp

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
)

// used when defining the type of variable/return value/parameter/etc
type TypeNode struct {
	Class         string
	TypeArguments []*TypeNode
}

func (tn *TypeNode) String() string {
	if len(tn.TypeArguments) > 0 {
		typArgs := make([]string, 0)
		for _, ta := range tn.TypeArguments {
			typArgs = append(typArgs, ta.String())
		}
		return fmt.Sprintf("%s[%s]", tn.Class, strings.Join(typArgs, ","))
	}
	return tn.Class
}

func NewTypeNode(typ parser.ITypeTypeContext) *TypeNode {
	ctx := typ.(*parser.TypeTypeContext)
	if ctx.PrimitiveType() != nil {
		return &TypeNode{
			Class: ctx.PrimitiveType().GetText(),
		}
	}

	ctx2 := ctx.ClassOrInterfaceType().(*parser.ClassOrInterfaceTypeContext)
	if len(ctx2.AllIDENTIFIER()) != 1 {
		panic("wrong number of identifers in type type: " + typ.GetText() + "\n\n" + typ.ToStringTree(parser.RuleNames, nil))
	}

	class := ctx2.IDENTIFIER(0).GetText()

	countTypArgs := len(ctx2.AllTypeArguments())

	if countTypArgs > 1 {
		panic("too many sets of type arguments")
	}

	var typeArguments []*TypeNode
	if countTypArgs == 1 {
		typeArguments = make([]*TypeNode, 0)
		typArgsCtx := ctx2.TypeArguments(0).(*parser.TypeArgumentsContext)
		for _, typArg := range typArgsCtx.AllTypeArgument() {
			typArgCtx := typArg.(*parser.TypeArgumentContext)

			if typArgCtx.QUESTION() != nil {
				panic("unknown")
			}

			typeArguments = append(typeArguments, NewTypeNode(typArgCtx.TypeType()))
		}
	}

	return &TypeNode{
		Class:         class,
		TypeArguments: typeArguments,
	}
}
