package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

type Lambda struct {
	*node.Base
	Arguments []*LocalVarDecl
	Body      node.Node
}

func (ln *Lambda) Children() []node.Node {
	return node.AppendNodeLists(ln.Arguments, ln.Body)
}

func NewLambda(lambda *parser.LambdaExpressionContext) *Lambda {
	lambdaCtx := lambda

	bodyCtx := lambdaCtx.LambdaBody()
	var body node.Node
	if bodyCtx.Expression() != nil {
		body = ExpressionProcessor(bodyCtx.Expression())
	} else if bodyCtx.Block() != nil {
		body = NewBlock(bodyCtx.Block())
	} else {
		panic("no body for lambda")
	}

	if body == nil {
		panic("no body for lambda")
	}

	arguments := make([]*LocalVarDecl, 0)
	parametersCtx := lambdaCtx.LambdaParameters()
	if len(parametersCtx.AllIDENTIFIER()) > 0 {
		// java lambda can have just parameter names, without types.
		for _, id := range parametersCtx.AllIDENTIFIER() {
			arguments = append(arguments, NewLocalVarDecl(nil, id.GetText(), nil))
		}
	} else {
		// must have formal parameters list
		arguments = FormalParameterListProcessor(parametersCtx.FormalParameterList())
	}

	return &Lambda{Base: node.New(), Arguments: arguments, Body: body}
}

func (ln *Lambda) String() string {
	arguments := ""
	if ln.Arguments != nil {
		arguments = ArgumentListToString(node.ListOfNodesToNodeList(ln.Arguments))
	}
	return fmt.Sprintf("func (%s) {%s}", arguments, ln.Body)
}
