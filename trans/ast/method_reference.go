package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

type MethodReference struct {
	*node.Base

	Instance node.Node
	Method   string
}

func (mr *MethodReference) Children() []node.Node {
	return []node.Node{mr.Instance}
}

func NewMethodReference(expression *parser.ExpressionContext) node.Node {
	ctx := expression

	method := ""
	if ctx.IDENTIFIER() != nil {
		method = ctx.IDENTIFIER().GetText()
	} else if ctx.NEW() != nil {
		method = "new"
	}

	if method == "" {
		panic("no method name in method reference")
	}

	var instance node.Node
	if ctx.Expression(0) != nil {
		instance = ExpressionProcessor(ctx.Expression(0))
	} else if ctx.TypeType(0) != nil {
		instance = NewTypeNodeFromContext(ctx.TypeType(0))
	} else if ctx.ClassType() != nil {
		panic("class type used: ") //  + ctx.ClassType().GetText())
		// instance = NewIdentifier(ctx.ClassType().GetText())
	}

	if instance == nil {
		panic("no instance/expression for method reference")
	}

	return &MethodReference{Base: node.New(), Method: method, Instance: instance}
}

func (mf *MethodReference) String() string {
	return fmt.Sprintf("%s.%s", mf.Instance, mf.Method)
}
