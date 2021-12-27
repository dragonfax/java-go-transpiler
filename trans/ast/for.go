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
	Body     node.Node
}

func (ef *EnhancedFor) Children() []node.Node {
	return []node.Node{ef.Variable, ef.Body}
}

func NewEnhancedFor(statementCtx *parser.StatementContext) *EnhancedFor {

	forControlCtx := statementCtx.ForControl()
	enhancedCtx := forControlCtx.EnhancedForControl()

	instance := ExpressionProcessor(enhancedCtx.Expression())
	variable := NewLocalVarDecl(NewTypeNodeFromContext(enhancedCtx.TypeType()), enhancedCtx.VariableDeclaratorId().GetText(), instance)

	body := StatementProcessor(statementCtx.Statement(0))
	return &EnhancedFor{
		Base: node.New(),

		Variable: variable,
		Body:     body,
	}
}

func (ef *EnhancedFor) String() string {
	return fmt.Sprintf("for %s := range %s %s", ef.Variable.Name, ef.Variable.Expression, ef.Body)
}

type ClassicFor struct {
	*node.Base

	Condition     node.Node       // should only be an expression with a bool type
	Init          []*LocalVarDecl // should only be localvardecl
	Increment     []node.Node     // should only be expression, no localvar dec
	Body          node.Node
	ConditionLast bool // Do...While
}

func (cf *ClassicFor) Children() []node.Node {
	list := []node.Node{cf.Body}
	list = append(list, node.ListOfNodesToNodeList(cf.Init)...)
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

func classicForControlProcessor(forControlCtx *parser.ForControlContext) (init []*LocalVarDecl, condition node.Node, increment []node.Node) {
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
		init = make([]*LocalVarDecl, 0)
		if initCtx.LocalVariableDeclaration() != nil {
			// variable declaractions
			declCtx := initCtx.LocalVariableDeclaration()
			init = NewLocalVarDeclNodeList(declCtx)
			for _, n := range init {
				if n == nil {
					panic("nil in node list")
				}
			}
		} else {
			// expression list
			// These should be replace-able in the original source before translation. to make them a simple local var
			fmt.Println("warning: for loops without localvardecl, just expression")
			for _, exp := range initCtx.ExpressionList().AllExpression() {
				node := ExpressionProcessor(exp)
				if node == nil {
					panic("nil into node list")
				}
				init = append(init, NewLocalVarDecl(nil, "TODO", node))
			}
		}
	}

	return
}
