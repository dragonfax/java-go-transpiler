package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/node"
)

/* method call
 * more like a function call,
 * just the method name, and arguments.
 * Doesn't have the class its connected to until after resolution phase is complete
 */
type MethodCall struct {
	*node.Base

	MethodName string
	Arguments  []node.Node
}

func (mc *MethodCall) Children() []node.Node {
	return mc.Arguments
}

func NewMethodCall(methodCall *parser.MethodCallContext) *MethodCall {
	if tool.IsNilInterface(methodCall) {
		panic("no method call")
	}

	methodCallCtx := methodCall

	methodName := ""
	if methodCallCtx.SUPER() != nil {
		methodName = "super"
	} else if methodCallCtx.THIS() != nil {
		methodName = "this"
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

	this := &MethodCall{Base: node.New(), MethodName: methodName, Arguments: arguments}
	return this
}

func (mc *MethodCall) String() string {
	return fmt.Sprintf("%s(%s)", mc.MethodName, ArgumentListToString(mc.Arguments))
}