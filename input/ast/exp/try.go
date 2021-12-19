package exp

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
)

type TryCatchNode struct {
	Body         ExpressionNode
	Finally      ExpressionNode
	CatchClauses []*CatchClause
}

func (tn *TryCatchNode) String() string {
	clauses := make([]string, 0)
	for _, clause := range tn.CatchClauses {
		s := fmt.Sprintf("catch (%s %s)%s", strings.Join(clause.CatchType, "."), clause.Variable, clause.Body)
		clauses = append(clauses, s)
	}

	if tn.Finally != nil {
		return fmt.Sprintf("try %s\n%sfinally %s", tn.Body, strings.Join(clauses, "\n"), tn.Finally)
	}
	return fmt.Sprintf("try %s %s", tn.Body, strings.Join(clauses, "\n"))
}

type CatchClause struct {
	Body      ExpressionNode
	CatchType []string
	Variable  string
}

func NewTryCatchNode(statement *parser.StatementContext) *TryCatchNode {

	ctx := statement

	if ctx.ResourceSpecification() != nil {
		panic("resource try/catch found.")
	}

	block := NewBlockNode(ctx.Block())

	var finallyBlock ExpressionNode
	finally := ctx.FinallyBlock()
	if finally != nil {
		finallyBlock = NewBlockNode(finally.Block())
	}

	clauses := make([]*CatchClause, 0)
	for _, catch := range ctx.AllCatchClause() {
		catchCtx := catch
		variable := catchCtx.IDENTIFIER().GetText()

		typeCtx := catchCtx.CatchType()
		if len(typeCtx.AllQualifiedName()) > 1 {
			panic("too many catch types.")
		}

		catchType := make([]string, 0)
		for _, ct := range typeCtx.QualifiedName(0).AllIDENTIFIER() {
			catchType = append(catchType, ct.GetText())
		}

		clauses = append(clauses, &CatchClause{
			Variable:  variable,
			Body:      NewBlockNode(catchCtx.Block()),
			CatchType: catchType,
		})
	}

	return &TryCatchNode{
		Body:         block,
		Finally:      finallyBlock,
		CatchClauses: clauses,
	}
}
