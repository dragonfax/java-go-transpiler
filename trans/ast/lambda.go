package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

type LambdaNode struct {
	*node.BaseNode
	Arguments []node.Node
	Body      node.Node
}

func (ln *LambdaNode) Children() []node.Node {
	return node.AppendNodeLists(ln.Arguments, ln.Body)
}

func NewLambdaNode(lambda *parser.LambdaExpressionContext) *LambdaNode {
	lambdaCtx := lambda

	bodyCtx := lambdaCtx.LambdaBody()
	var body node.Node
	if bodyCtx.Expression() != nil {
		body = ExpressionProcessor(bodyCtx.Expression())
	} else if bodyCtx.Block() != nil {
		body = NewBlockNode(bodyCtx.Block())
	} else {
		panic("no body for lambda")
	}

	if body == nil {
		panic("no body for lambda")
	}

	arguments := make([]node.Node, 0)
	parametersCtx := lambdaCtx.LambdaParameters()
	if len(parametersCtx.AllIDENTIFIER()) > 0 {
		// java lambda can have just parameter names, without types. thats valid
		for _, id := range parametersCtx.AllIDENTIFIER() {
			arguments = append(arguments, NewIdentifierNode(id.GetText()))
		}
	} else {
		// must have formal parameters list
		arguments = FormalParameterListProcessor(parametersCtx.FormalParameterList())
	}

	return &LambdaNode{BaseNode: node.NewNode(), Arguments: arguments, Body: body}
}

func (ln *LambdaNode) String() string {
	arguments := ""
	if ln.Arguments != nil {
		arguments = ArgumentListToString(ln.Arguments)
	}
	return fmt.Sprintf("func (%s) {%s}", arguments, ln.Body)
}
