package ast

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

func NewFor(statementCtx *parser.StatementContext) node.Node {

	if statementCtx.ForControl().EnhancedForControl() != nil {
		return NewEnhancedFor(statementCtx)
	}
	return NewClassicFor(statementCtx)
}

type EnhancedFor struct {
	*node.Base

	Variable *LocalVarDecl
	Instance node.Node
	Body     node.Node
}

func (ef *EnhancedFor) Children() []node.Node {
	return []node.Node{ef.Variable, ef.Instance, ef.Body}
}

func NewEnhancedFor(statementCtx *parser.StatementContext) *EnhancedFor {

	forControlCtx := statementCtx.ForControl()
	enhancedCtx := forControlCtx.EnhancedForControl()

	instance := ExpressionProcessor(enhancedCtx.Expression())
	variable := NewLocalVarDecl(NewTypeNodeFromContext(enhancedCtx.TypeType()), enhancedCtx.VariableDeclaratorId().GetText(), nil)

	body := StatementProcessor(statementCtx.Statement(0))
	return &EnhancedFor{
		Base: node.New(),

		Variable: variable,
		Instance: instance,
		Body:     body,
	}
}

func (ef *EnhancedFor) String() string {
	return fmt.Sprintf("for %s := range %s %s", ef.Variable.Name, ef.Instance, ef.Body)
}

type ClassicFor struct {
	*node.Base

	Condition     node.Node
	Init          []node.Node
	Increment     []node.Node
	Body          node.Node
	ConditionLast bool // Do...While
}

func (cf *ClassicFor) Children() []node.Node {
	list := []node.Node{cf.Body}
	list = append(list, cf.Init...)
	list = append(list, cf.Increment...)
	if cf.Condition != nil {
		list = append(list, cf.Condition)
	}
	return list
}

func (fn *ClassicFor) String() string {
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

func NewClassicFor(statementCtx *parser.StatementContext) *ClassicFor {
	init, condition, increment := classicForControlProcessor(statementCtx.ForControl())
	return &ClassicFor{
		Base: node.New(),

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
			init = node.ListOfNodesToNodeList(NewLocalVarDeclNodeList(declCtx))
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
