package ast

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

type TryCatchNode struct {
	*node.BaseNode

	Body         node.Node
	Finally      node.Node
	CatchClauses []*CatchClause
}

func (tc *TryCatchNode) Children() []node.Node {
	list := []node.Node{tc.Body}
	if tc.Finally != nil {
		list = append(list, tc.Finally)
	}
	list = node.AppendNodeLists(list, tc.CatchClauses...)
	return list
}

func (tn *TryCatchNode) String() string {
	clauses := make([]string, 0)
	for _, clause := range tn.CatchClauses {
		s := fmt.Sprintf("catch (%s %s)%s", strings.Join(clause.CatchType, "."), clause.Variable, clause.Body)
		clauses = append(clauses, s)
	}

	finally := ""
	if tn.Finally != nil {
		finally = fmt.Sprintf("// finally\n%s\n", tn.Finally.String())
	}

	catches := ""

	if len(tn.CatchClauses) == 1 && len(tn.CatchClauses[0].CatchType) == 1 && tn.CatchClauses[0].CatchType[0] == "Exception" {
		// common case

		clause := tn.CatchClauses[0]

		catches = fmt.Sprintf(`
if err != nil {
	%s := err
	%s
}
`, clause.Variable, clause.Body)
	} else {
		for _, c := range tn.CatchClauses {
			catches += fmt.Sprintf("catch %s %s\n", c.CatchType, c.Body)
		}
		catches = "/* TODO\n" + catches + "\n*/"
	}

	return fmt.Sprintf("// try\n%s\n%s\n%s\n", tn.Body, catches, finally)
}

type CatchClause struct {
	*node.BaseNode

	Body      node.Node
	CatchType []string
	Variable  string
}

func (cc *CatchClause) String() string {
	return "catch"
}

func (cc *CatchClause) Children() []node.Node {
	return []node.Node{cc.Body}
}

func NewTryCatchNode(statement *parser.StatementContext) *TryCatchNode {

	ctx := statement

	if ctx.ResourceSpecification() != nil {
		panic("resource try/catch found.")
	}

	block := NewBlockNode(ctx.Block())

	var finallyBlock node.Node
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
			BaseNode: node.NewNode(),

			Variable:  variable,
			Body:      NewBlockNode(catchCtx.Block()),
			CatchType: catchType,
		})
	}

	return &TryCatchNode{
		BaseNode:     node.NewNode(),
		Body:         block,
		Finally:      finallyBlock,
		CatchClauses: clauses,
	}
}
