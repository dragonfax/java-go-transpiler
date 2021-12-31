package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

/* Not a method call, but a reference to a method that can be passed around and called like a function later.
 * Thus no arguments to speak of there.
 * It IS an expression, though its type is a Function<A,R> with Arguments and Return values (both type parameters ignored in this system)
 */
type MethodRef struct {
	*BaseExpression

	Instance   node.Node
	MethodName string
}

func (mr *MethodRef) GetType() *Class {
	// TODO We drop the type parmaeter here, but we may want it in the future, as we may have to type dereferences of this.
	return RuntimePackage.GetClass("Function")
}

func (mr *MethodRef) Children() []node.Node {
	return []node.Node{mr.Instance}
}

func NewMethodRef(expression *parser.ExpressionContext) *MethodRef {
	ctx := expression

	methodName := ""
	if ctx.IDENTIFIER() != nil {
		methodName = ctx.IDENTIFIER().GetText()
	} else if ctx.NEW() != nil {
		methodName = "new"
	}

	if methodName == "" {
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

	return &MethodRef{
		BaseExpression: NewExpression(),
		MethodName:     methodName,
		Instance:       instance,
	}
}

func (mf *MethodRef) String() string {
	return fmt.Sprintf("%s.%s", mf.Instance, mf.MethodName)
}
