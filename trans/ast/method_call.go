package ast

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/node"
)

/* method or constructor call
 * more like a function call,
 * just the method name, and arguments.
 * Doesn't have the class its connected to until after resolution phase is complete
 */
type MethodCall struct {
	*BaseExpression
	*BaseMethodScope

	MethodName         string // or class name for constructors
	TypeArguments      []*TypePath
	Arguments          []node.Node
	Super              bool
	This               bool
	Constructor        bool
	InstanceExpression node.Node

	Class  *Class // class with method, or to be constructed, after resolution
	Method *Method
}

func (mc *MethodCall) GetType() *Class {
	if mc.Constructor {
		mc.Type = mc.Class
	} else if mc.Method != nil {
		mc.Type = mc.Method.GetType()
	}
	return mc.Type
}

func (mc *MethodCall) NodeName() string {
	t := "MethodCall"
	if mc.Constructor {
		t = "ConstructorCall"
	}
	return t + " = " + mc.MethodName
}

func (mc *MethodCall) Children() []node.Node {
	return node.AppendNodeLists(mc.TypeArguments, mc.Arguments...)
}

func NewMethodCall(methodCall *parser.MethodCallContext, instance node.Node) *MethodCall {
	if tool.IsNilInterface(methodCall) {
		panic("no method call")
	}

	methodCallCtx := methodCall

	methodName := ""
	var super bool
	var isThis bool
	if methodCallCtx.SUPER() != nil {
		methodName = "super"
		super = true
	} else if methodCallCtx.THIS() != nil {
		methodName = "this"
		isThis = true
	} else if methodCallCtx.IDENTIFIER() != nil {
		methodName = methodCallCtx.IDENTIFIER().GetText()
	} else {
		panic("no method name in method call")
	}

	arguments := make([]node.Node, 0)

	if methodCallCtx.ExpressionList() != nil {
		for _, expression := range methodCallCtx.ExpressionList().AllExpression() {
			node := ExpressionProcessor(expression)
			if node == nil {
				panic("nil in node list")
			}
			arguments = append(arguments, node)
		}
	}

	this := &MethodCall{
		BaseExpression:     NewExpression(),
		BaseMethodScope:    NewMethodScope(),
		MethodName:         methodName,
		Arguments:          arguments,
		Super:              super,
		This:               isThis,
		InstanceExpression: instance,
	}

	return this
}

func NewConstructorCall(creator *parser.CreatorContext) *MethodCall {
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

	typeArguments := make([]*TypePath, 0)
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

	return &MethodCall{
		BaseExpression:  NewExpression(),
		BaseMethodScope: NewMethodScope(),
		MethodName:      className,
		TypeArguments:   typeArguments,
		Arguments:       arguments,
	}
}

func (mc *MethodCall) String() string {
	if mc.Constructor {
		return mc.ConstructorString()
	}

	var instance = ""
	if mc.InstanceExpression != nil {
		instance = mc.InstanceExpression.String()
	}
	return fmt.Sprintf("%s%s(%s)", instance, mc.MethodName, ArgumentListToString(mc.Arguments))
}

func (cc *MethodCall) ConstructorString() string {
	if len(cc.TypeArguments) == 0 {
		return fmt.Sprintf("New%s(%s)", cc.MethodName, ArgumentListToString(cc.Arguments))
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

	return fmt.Sprintf("New%s%s[%s](%s)", cc.MethodName, argumentCount, strings.Join(list, ","), ArgumentListToString(cc.Arguments))
}
