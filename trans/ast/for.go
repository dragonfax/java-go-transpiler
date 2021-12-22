package ast

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

func NewForNode(statementCtx *parser.StatementContext) node.Node {

	if statementCtx.ForControl().EnhancedForControl() != nil {
		return NewEnhancedForNode(statementCtx)
	}
	return NewClassicForNode(statementCtx)
}

type EnhancedForNode struct {
	*node.BaseNode

	Variable *VariableDeclNode
	Instance node.Node
	Body     node.Node
}

func (ef *EnhancedForNode) Children() []node.Node {
	return []node.Node{ef.Variable, ef.Instance, ef.Body}
}

func NewEnhancedForNode(statementCtx *parser.StatementContext) *EnhancedForNode {

	forControlCtx := statementCtx.ForControl()
	enhancedCtx := forControlCtx.EnhancedForControl()

	instance := ExpressionProcessor(enhancedCtx.Expression())
	variable := NewVariableDecl(NewTypeNode(enhancedCtx.TypeType()), enhancedCtx.VariableDeclaratorId().GetText(), nil)

	body := StatementProcessor(statementCtx.Statement(0))
	return &EnhancedForNode{
		BaseNode: node.NewNode(),

		Variable: variable,
		Instance: instance,
		Body:     body,
	}
}

func (ef *EnhancedForNode) String() string {
	return fmt.Sprintf("for %s := range %s %s", ef.Variable.Name, ef.Instance, ef.Body)
}

type ClassicForNode struct {
	*node.BaseNode

	Condition     node.Node
	Init          []node.Node
	Increment     []node.Node
	Body          node.Node
	ConditionLast bool // Do...While
}

func (cf *ClassicForNode) Children() []node.Node {
	list := []node.Node{cf.Body}
	list = append(list, cf.Init...)
	list = append(list, cf.Increment...)
	if cf.Condition != nil {
		list = append(list, cf.Condition)
	}
	return list
}

func (fn *ClassicForNode) String() string {
	// TODO ConditionLast
	// TODO remove unnecessary semicolons

	init := make([]string, 0)
	for _, i := range fn.Init {
		init = append(init, i.String())
	}

	incr := make([]string, 0)
	for _, i := range fn.Increment {
		incr = append(incr, i.String())
	}

	return fmt.Sprintf("for %s;%s;%s {\n%s}\n", strings.Join(init, ",'"), fn.Condition, strings.Join(incr, ","), fn.Body)
}

func NewClassicForNode(statementCtx *parser.StatementContext) *ClassicForNode {
	init, condition, increment := classicForControlProcessor(statementCtx.ForControl())
	return &ClassicForNode{
		BaseNode: node.NewNode(),

		Condition: condition,
		Init:      init,
		Increment: increment,
		Body:      StatementProcessor(statementCtx.Statement(0)),
	}
}

func classicForControlProcessor(forControlCtx *parser.ForControlContext) (init []node.Node, condition node.Node, increment []node.Node) {
	if forControlCtx.GetForUpdate() != nil {
		increment = make([]node.Node, 0)
		for _, exp := range forControlCtx.GetForUpdate().AllExpression() {
			node := ExpressionProcessor(exp)
			if node == nil {
				panic("nil into node list")
			}
			increment = append(increment, node)
		}
	}

	if forControlCtx.Expression() != nil {
		condition = ExpressionProcessor(forControlCtx.Expression())
	}

	if forControlCtx.ForInit() != nil {
		initCtx := forControlCtx.ForInit()
		init = make([]node.Node, 0)
		if initCtx.LocalVariableDeclaration() != nil {
			// variable declaractions
			declCtx := initCtx.LocalVariableDeclaration()
			init = NewVariableDeclNodeList(declCtx)
			for _, n := range init {
				if n == nil {
					panic("nil in node list")
				}
			}
		} else {
			// expression list
			for _, exp := range initCtx.ExpressionList().AllExpression() {
				node := ExpressionProcessor(exp)
				if node == nil {
					panic("nil into node list")
				}
				init = append(init, node)
			}
		}
	}

	return
}
