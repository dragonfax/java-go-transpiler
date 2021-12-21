package ast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/node"
)

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

func (vn *VariableNode) Children() []node.Node {
	return nil
}

func (vn *VariableNode) String() string {
	return vn.Name
}

type IfNode struct {
	Condition node.Node
	Body      node.Node
	Else      node.Node
}

func (in *IfNode) Children() []node.Node {
	list := []node.Node{in.Body}
	if in.Condition != nil {
		list = append(list, in.Condition)
	}
	if in.Else != nil {
		list = append(list, in.Else)
	}
	return list
}

func NewIfNode(condition, body, els node.Node) *IfNode {
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
	Expression node.Node
}

func (rn *ReturnNode) Children() []node.Node {
	if rn.Expression != nil {
		return []node.Node{rn.Expression}
	}
	return nil
}

func NewReturnNode(exp node.Node) *ReturnNode {
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
	Expression string
}

func (tn *ThrowNode) Children() []node.Node {
	return nil
}

func NewThrowNode(exp string) *ThrowNode {
	return &ThrowNode{Expression: exp}
}

func (tn *ThrowNode) String() string {
	return fmt.Sprintf("panic(%s) // TODO\n", strconv.Quote(tn.Expression))
}

type BreakNode struct {
	Label string
}

func (b *BreakNode) Children() []node.Node {
	return nil
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

func (c *ContinueNode) Children() []node.Node {
	return nil
}

func NewContinueNode(label string) *ContinueNode {
	return &ContinueNode{Label: label}
}

func (cn *ContinueNode) String() string {
	return fmt.Sprintf("continue %s\n", cn.Label)
}

type LabelNode struct {
	Label      string
	Expression node.Node
}

func (l *LabelNode) Children() []node.Node {
	return []node.Node{l.Expression}
}

func NewLabelNode(label string, exp node.Node) *LabelNode {
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

type InstanceAttributeReference struct {
	Attribute         string
	InstanceReference node.Node
}

func (ia *InstanceAttributeReference) Children() []node.Node {
	return []node.Node{ia.InstanceReference}
}

func NewInstanceAttributeReference(attribute string, instanceExpression node.Node) *InstanceAttributeReference {
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
	Instance   node.Node
	MethodName string
	Arguments  []node.Node
}

func (mc *MethodCall) Children() []node.Node {
	list := []node.Node{mc.Instance}
	list = append(list, mc.Arguments...)
	return list
}

func NewMethodCall(instance node.Node, methodCall *parser.MethodCallContext) *MethodCall {
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

	this := &MethodCall{Instance: instance, MethodName: methodName, Arguments: arguments}
	return this
}

func (mc *MethodCall) String() string {
	if mc.Instance == nil {
		return fmt.Sprintf("%s(%s)", mc.MethodName, ArgumentListToString(mc.Arguments))
	}
	return fmt.Sprintf("%s.%s(%s)", mc.Instance, mc.MethodName, ArgumentListToString(mc.Arguments))
}

type IdentifierNode struct {
	Identifier string
}

func (i *IdentifierNode) Children() []node.Node {
	return nil
}

func NewIdentifierNode(id string) *IdentifierNode {
	return &IdentifierNode{Identifier: id}
}

func (in *IdentifierNode) String() string {
	return in.Identifier
}

type ConstructorCall struct {
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

type ClassReference struct {
	Class string
}

func (cr *ClassReference) Children() []node.Node {
	return nil
}

func NewClassReference(class string) *ClassReference {
	if class == "" {
		panic("no class name")
	}
	return &ClassReference{Class: class}
}

func (cr *ClassReference) String() string {
	return cr.Class + ".class"
}

type LambdaNode struct {
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

	return &LambdaNode{Arguments: arguments, Body: body}
}

func (ln *LambdaNode) String() string {
	arguments := ""
	if ln.Arguments != nil {
		arguments = ArgumentListToString(ln.Arguments)
	}
	return fmt.Sprintf("func (%s) {%s}", arguments, ln.Body)
}

type MethodReference struct {
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
		instance = NewTypeNode(ctx.TypeType(0))
	} else if ctx.ClassType() != nil {
		instance = NewIdentifierNode(ctx.ClassType().GetText())
	}

	if instance == nil {
		panic("no instance/expression for method reference")
	}

	return &MethodReference{Method: method, Instance: instance}
}

func (mf *MethodReference) String() string {
	return fmt.Sprintf("%s.%s", mf.Instance, mf.Method)
}
