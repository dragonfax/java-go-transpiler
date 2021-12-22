package ast

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

type ConstructorCall struct {
	*node.BaseNode

	Class         string
	TypeArguments []TypeNode
	Arguments     []node.Node
}

func (cc *ConstructorCall) Children() []node.Node {
	return node.AppendNodeLists(cc.TypeArguments, cc.Arguments...)
}

func NewConstructorCall(creator *parser.CreatorContext) *ConstructorCall {
	if creator == nil {
		panic("empty creator call")
	}
	creatorCtx := creator

	creatorNameCtx := creatorCtx.CreatedName()
	class := ""
	if creatorNameCtx.IDENTIFIER(0) != nil {
		class = creatorNameCtx.IDENTIFIER(0).GetText()
	} else if creatorNameCtx.PrimitiveType() != nil {
		class = creatorNameCtx.PrimitiveType().GetText()
	} else {
		panic("constructor call with no class name")
	}

	typeArguments := make([]TypeNode, 0)
	if creatorNameCtx.TypeArgumentsOrDiamond(0) != nil {
		if creatorNameCtx.TypeArgumentsOrDiamond(0).TypeArguments() != nil {
			for _, typeArg := range creatorNameCtx.TypeArgumentsOrDiamond(0).TypeArguments().AllTypeArgument() {
				typeArgCtx := typeArg
				node := NewTypeNode(typeArgCtx.TypeType())
				typeArguments = append(typeArguments, node)
			}
		}
	}

	arguments := make([]node.Node, 0)
	if creatorCtx.ClassCreatorRest() != nil {
		if creatorCtx.ClassCreatorRest().Arguments().ExpressionList() != nil {
			for _, expression := range creatorCtx.ClassCreatorRest().Arguments().ExpressionList().AllExpression() {
				node := ExpressionProcessor(expression)
				if node == nil {
					panic("nil in node list")
				}
				arguments = append(arguments, node)
			}
		}
	}

	return &ConstructorCall{
		BaseNode:      node.NewNode(),
		Class:         class,
		TypeArguments: typeArguments,
		Arguments:     arguments,
	}
}

func (cc *ConstructorCall) String() string {
	if len(cc.TypeArguments) == 0 {
		return fmt.Sprintf("New%s(%s)", cc.Class, ArgumentListToString(cc.Arguments))
	}

	list := make([]string, 0)
	for _, ta := range cc.TypeArguments {
		list = append(list, ta.String())
	}
	return fmt.Sprintf("New%s[%s](%s)", cc.Class, strings.Join(list, ","), ArgumentListToString(cc.Arguments))
}
