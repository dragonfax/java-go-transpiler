package ast

import (
	"strings"

	"github.com/dragonfax/java-go-transpiler/input/parser"
	"github.com/dragonfax/java-go-transpiler/trans/node"
)

type TypeParameterList struct {
	*node.Base

	TypeParameters []*TypeParameter
}

func NewTypeParameterList(ctx *parser.TypeParametersContext) *TypeParameterList {
	this := &TypeParameterList{
		Base:           node.New(),
		TypeParameters: make([]*TypeParameter, 0),
	}
	for _, tpCtx := range ctx.AllTypeParameter() {
		this.TypeParameters = append(this.TypeParameters, NewTypeParameter(tpCtx))
	}

	return this
}

func (tl *TypeParameterList) String() string {
	return strings.Join(node.NodeListToStringList(tl.TypeParameters), ",")
}

func (tl *TypeParameterList) Children() []node.Node {
	if tl == nil {
		panic("nil type parameter list")
	}
	return node.ListOfNodesToNodeList(tl.TypeParameters)
}

type TypeParameter struct {
	*node.Base

	Name string
}

func NewTypeParameter(ctx *parser.TypeParameterContext) *TypeParameter {
	if ctx.EXTENDS() != nil {
		panic("type parameter with bounds")
	}
	return &TypeParameter{Base: node.New(), Name: ctx.IDENTIFIER().GetText()}
}

func (tp *TypeParameter) String() string {
	return tp.Name
}

func (tp *TypeParameter) Children() []node.Node {
	return nil
}
