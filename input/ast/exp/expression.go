package exp

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/tool"
)

type ExpressionNode interface {
	String() string
}

func expressionListToString(list []ExpressionNode) string {
	if list == nil {
		panic("list expression list")
	}
	s := ""
	for _, node := range list {
		if tool.IsNilInterface(node) {
			panic("nil node in expression list")
		}
		s += node.String() + "\n"
	}
	return s
}

type LiteralNode struct {
	Value string
}

func NewLiteralNode(value string) *LiteralNode {
	if value == "" {
		panic("no value")
	}
	return &LiteralNode{Value: value}
}

func (ln *LiteralNode) String() string {
	return ln.Value
}

type VariableNode struct {
	Name string
}

func NewVariableNode(name string) *VariableNode {
	if name == "" {
		panic("missing name")
	}
	return &VariableNode{
		Name: name,
	}
}

func (vn *VariableNode) String() string {
	return vn.Name
}

type IfNode struct {
	Condition ExpressionNode
	Body      ExpressionNode
	Else      ExpressionNode
}

func NewIfNode(condition, body, els ExpressionNode) *IfNode {
	if tool.IsNilInterface(body) {
		panic("missing body")
	}
	if tool.IsNilInterface(condition) {
		panic("missing condition")
	}
	return &IfNode{
		Condition: condition,
		Body:      body,
		Else:      els,
	}
}

func (in *IfNode) String() string {
	if tool.IsNilInterface(in.Else) {
		return fmt.Sprintf("if %s {\n%s}\n", in.Condition, in.Body)
	}
	return fmt.Sprintf("if %s {\n%s} else {\n%s}\n", in.Condition, in.Body, in.Else)
}

type ReturnNode struct {
	Expression ExpressionNode
}

func NewReturnNode(exp ExpressionNode) *ReturnNode {
	return &ReturnNode{Expression: exp}
}

func (rn *ReturnNode) String() string {
	exp := ""
	if !tool.IsNilInterface(rn.Expression) {
		exp = rn.Expression.String()
	}
	return fmt.Sprintf("return %s\n", exp)
}

type ThrowNode struct {
	Expression ExpressionNode
}

func NewThrowNode(exp ExpressionNode) *ThrowNode {
	if tool.IsNilInterface(exp) {
		panic("missing expression")
	}
	return &ThrowNode{Expression: exp}
}

func (tn *ThrowNode) String() string {
	return fmt.Sprintf("panic(%s)\n", tn.Expression.String())
}

type BreakNode struct {
	Label string
}

func NewBreakNode(label string) *BreakNode {
	return &BreakNode{Label: label}
}

func (bn *BreakNode) String() string {
	return fmt.Sprintf("break %s\n", bn.Label)
}

type ContinueNode struct {
	Label string
}

func NewContinueNode(label string) *ContinueNode {
	return &ContinueNode{Label: label}
}

func (cn *ContinueNode) String() string {
	return fmt.Sprintf("continue %s\n", cn.Label)
}

type LabelNode struct {
	Label      string
	Expression ExpressionNode
}

func NewLabelNode(label string, exp ExpressionNode) *LabelNode {
	if label == "" {
		panic("label missing")
	}
	if tool.IsNilInterface(exp) {
		panic("expression missing")
	}
	return &LabelNode{Label: label, Expression: exp}
}

func (ln *LabelNode) String() string {
	return fmt.Sprintf("%s: %s\n", ln.Label, ln.Expression)
}

type SelfReference struct {
}

func NewSelfReference() *SelfReference {
	return &SelfReference{}
}

func (sr *SelfReference) String() string {
	return "this"
}

type InstanceAttributeReference struct {
	Attribute         string
	InstanceReference ExpressionNode
}

func NewInstanceAttributeReference(attribute string, instanceExpression ExpressionNode) *InstanceAttributeReference {
	if attribute == "" {
		panic("no attribute")
	}
	if tool.IsNilInterface(instanceExpression) {
		panic("no instance")
	}
	this := &InstanceAttributeReference{Attribute: attribute, InstanceReference: instanceExpression}
	return this
}

func (ia *InstanceAttributeReference) String() string {
	return fmt.Sprintf("%s.%s", ia.InstanceReference, ia.Attribute)
}

type MethodCall struct {
	Instance   string
	MethodName string
	Arguments  []ExpressionNode
}

func NewMethodCall(instance string, methodCall parser.IMethodCallContext) *MethodCall {
	if instance == "" {
		panic("no instance name in call")
	}
	if tool.IsNilInterface(methodCall) {
		panic("no method call")
	}

	methodCallCtx := methodCall.(*parser.MethodCallContext)

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

	arguments := make([]ExpressionNode, 0)

	for _, expression := range methodCallCtx.ExpressionList().(*parser.ExpressionListContext).AllExpression() {
		arguments = append(arguments, expressionProcessor(expression))
	}

	this := &MethodCall{Instance: instance, MethodName: methodName, Arguments: arguments}
	return this
}

func (mc *MethodCall) String() string {
	return fmt.Sprintf("%s.%s(%s)", mc.Instance, mc.MethodName, expressionListToString(mc.Arguments))
}
