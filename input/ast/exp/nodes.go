package exp

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/tool"
)

type ExpressionNode interface {
	String() string
}

func argumentListToString(list []ExpressionNode) string {
	s := make([]string, 0)
	for _, node := range list {
		if tool.IsNilInterface(node) {
			panic("nil node in expression list")
		}
		s = append(s, node.String())
	}
	return strings.Join(s, ",")
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
	Instance   ExpressionNode
	MethodName string
	Arguments  []ExpressionNode
}

func NewMethodCall(instance ExpressionNode, methodCall parser.IMethodCallContext) *MethodCall {
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

	if methodCallCtx.ExpressionList() != nil {
		for _, expression := range methodCallCtx.ExpressionList().(*parser.ExpressionListContext).AllExpression() {
			arguments = append(arguments, ExpressionProcessor(expression))
		}
	}

	this := &MethodCall{Instance: instance, MethodName: methodName, Arguments: arguments}
	return this
}

func (mc *MethodCall) String() string {
	if mc.Instance == nil {
		return fmt.Sprintf("%s(%s)", mc.MethodName, argumentListToString(mc.Arguments))
	}
	return fmt.Sprintf("%s.%s(%s)", mc.Instance, mc.MethodName, argumentListToString(mc.Arguments))
}

type IdentifierNode struct {
	Identifier string
}

func NewIdentifierNode(id string) *IdentifierNode {
	return &IdentifierNode{Identifier: id}
}

func (in *IdentifierNode) String() string {
	return in.Identifier
}

type ConstructorCall struct {
	Class         string
	TypeArguments []string
	Arguments     []ExpressionNode
}

func NewConstructorCall(creator parser.ICreatorContext) *ConstructorCall {
	if creator == nil {
		panic("empty creator call")
	}
	creatorCtx := creator.(*parser.CreatorContext)

	creatorNameCtx := creatorCtx.CreatedName().(*parser.CreatedNameContext)
	class := ""
	if creatorNameCtx.IDENTIFIER(0) != nil {
		class = creatorNameCtx.IDENTIFIER(0).GetText()
	} else if creatorNameCtx.PrimitiveType() != nil {
		class = creatorNameCtx.PrimitiveType().GetText()
	} else {
		panic("constructor call with no class name")
	}

	typeArguments := make([]string, 0)
	if creatorNameCtx.TypeArgumentsOrDiamond(0) != nil {
		if creatorNameCtx.TypeArgumentsOrDiamond(0).(*parser.TypeArgumentsOrDiamondContext).TypeArguments() != nil {
			for _, typeArg := range creatorNameCtx.TypeArgumentsOrDiamond(0).(*parser.TypeArgumentsOrDiamondContext).TypeArguments().(*parser.TypeArgumentsContext).AllTypeArgument() {
				typeArgCtx := typeArg.(*parser.TypeArgumentContext)
				typeArguments = append(typeArguments, typeArgCtx.GetText())
			}
		}
	}

	arguments := make([]ExpressionNode, 0)
	if creatorCtx.ClassCreatorRest() != nil {
		if creatorCtx.ClassCreatorRest().(*parser.ClassCreatorRestContext).Arguments().(*parser.ArgumentsContext).ExpressionList() != nil {
			for _, expression := range creatorCtx.ClassCreatorRest().(*parser.ClassCreatorRestContext).Arguments().(*parser.ArgumentsContext).ExpressionList().(*parser.ExpressionListContext).AllExpression() {
				arguments = append(arguments, ExpressionProcessor(expression))
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
		return fmt.Sprintf("New%s(%s)", cc.Class, expressionListToString(cc.Arguments))
	}
	return fmt.Sprintf("New%s[%s](%s)", cc.Class, strings.Join(cc.TypeArguments, ",'"), expressionListToString(cc.Arguments))
}

type ClassReference struct {
	Class string
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
	Arguments []ExpressionNode
	Body      ExpressionNode
}

func NewLambdaNode(lambda parser.ILambdaExpressionContext) *LambdaNode {
	lambdaCtx := lambda.(*parser.LambdaExpressionContext)

	bodyCtx := lambdaCtx.LambdaBody().(*parser.LambdaBodyContext)
	var body ExpressionNode
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

	arguments := make([]ExpressionNode, 0)
	parametersCtx := lambdaCtx.LambdaParameters().(*parser.LambdaParametersContext)
	if len(parametersCtx.AllIDENTIFIER()) > 0 {
		// java lambda can have just parameter names, without types. thats valid
		for _, id := range parametersCtx.AllIDENTIFIER() {
			arguments = append(arguments, NewIdentifierNode(id.GetText()))
		}
	} else {
		// must have formal parameters list
		arguments = formalParameterListProcessor(parametersCtx.FormalParameterList())
	}

	return &LambdaNode{Arguments: arguments, Body: body}
}

func (ln *LambdaNode) String() string {
	arguments := ""
	if ln.Arguments != nil {
		arguments = expressionListToString(ln.Arguments)
	}
	return fmt.Sprintf("func (%s) {%s}", arguments, ln.Body)
}

type MethodReference struct {
	Instance ExpressionNode
	Method   string
}

func NewMethodReference(expression parser.IExpressionContext) ExpressionNode {
	ctx := expression.(*parser.ExpressionContext)

	method := ""
	if ctx.IDENTIFIER() != nil {
		method = ctx.IDENTIFIER().GetText()
	} else if ctx.NEW() != nil {
		method = "new"
	}

	if method == "" {
		panic("no method name in method reference")
	}

	var instance ExpressionNode
	if ctx.Expression(0) != nil {
		instance = ExpressionProcessor(ctx.Expression(0))
	} else if ctx.TypeType(0) != nil {
		instance = NewIdentifierNode(ctx.TypeType(0).GetText())
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
