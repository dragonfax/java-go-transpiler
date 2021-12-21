package exp

import "github.com/dragonfax/java_converter/input/parser"

type TypeParameterList []*TypeParameter

func NewTypeParameterList(ctx *parser.TypeParametersContext) TypeParameterList {
	list := make(TypeParameterList, 0)
	for _, tpCtx := range ctx.AllTypeParameter() {
		list = append(list, NewTypeParameter(tpCtx))
	}

	return list
}

type TypeParameter struct {
	Name string
}

func NewTypeParameter(ctx *parser.TypeParameterContext) *TypeParameter {
	if ctx.EXTENDS() != nil {
		panic("type parameter with bounds")
	}
	return &TypeParameter{Name: ctx.IDENTIFIER().GetText()}
}
