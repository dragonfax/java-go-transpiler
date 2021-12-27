package ast

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

type ConstructorCall struct {
	*node.Base

	ClassName     string
	TypeArguments []*Type
	Arguments     []node.Node

	Class *Class // for resolution
}

func (cc *ConstructorCall) Name() string {
	return "ConstructorCall = " + cc.ClassName
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
	className := ""
	if creatorNameCtx.IDENTIFIER(0) != nil {
		className = creatorNameCtx.IDENTIFIER(0).GetText()
	} else if creatorNameCtx.PrimitiveType() != nil {
		className = creatorNameCtx.PrimitiveType().GetText()
	} else {
		panic("constructor call with no class name")
	}

	typeArguments := make([]*Type, 0)
	if creatorNameCtx.TypeArgumentsOrDiamond(0) != nil {
		if creatorNameCtx.TypeArgumentsOrDiamond(0).TypeArguments() != nil {
			for _, typeArg := range creatorNameCtx.TypeArgumentsOrDiamond(0).TypeArguments().AllTypeArgument() {
				typeArgCtx := typeArg
				node := NewTypeNodeFromContext(typeArgCtx.TypeType())
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
		Base:          node.New(),
		ClassName:     className,
		TypeArguments: typeArguments,
		Arguments:     arguments,
	}
}

func (cc *ConstructorCall) String() string {
	if len(cc.TypeArguments) == 0 {
		return fmt.Sprintf("New%s(%s)", cc.ClassName, ArgumentListToString(cc.Arguments))
	}

	list := make([]string, 0)
	for _, ta := range cc.TypeArguments {
		list = append(list, ta.String())
	}

	// We include the argument count to methods and constructors (when its > 0), as golang doesn't support method overloading
	argumentCount := ""
	if len(cc.Arguments) > 0 {
		argumentCount = fmt.Sprintf("%d", len(cc.Arguments))
	}

	return fmt.Sprintf("New%s%s[%s](%s)", cc.ClassName, argumentCount, strings.Join(list, ","), ArgumentListToString(cc.Arguments))
}
