package exp

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/node"
)

type BlockNode struct {
	Body []node.Node
}

func (bn *BlockNode) Children() []node.Node {
	if bn == nil {
		panic("block node was nil")
	}
	return node.ListOfNodesToNodeList(bn.Body)
}

func (bn *BlockNode) String() string {
	if bn == nil {
		panic("nil block")
	}
	return fmt.Sprintf("{\n%s}\n", expressionListToString(bn.Body))
}

func NewBlockNode(block *parser.BlockContext) *BlockNode {
	if tool.IsNilInterface(block) {
		return nil
	}
	ctx := block

	l := make([]node.Node, 0)

	for _, blockStatement := range ctx.AllBlockStatement() {
		s := BlockStatementProcessor(blockStatement)
		l = append(l, s...)
	}

	return &BlockNode{Body: l}
}

func BlockStatementProcessor(ctx *parser.BlockStatementContext) []node.Node {
	if ctx.LocalVariableDeclaration() != nil {
		localVarCtx := ctx.LocalVariableDeclaration()
		nodes := NewVariableDeclNodeList(localVarCtx)
		for _, n := range nodes {
			if n == nil {
				panic("nil in expression list")
			}
		}
		return nodes
	} else if ctx.Statement() != nil {
		statementCtx := ctx.Statement()
		stmt := StatementProcessor(statementCtx)
		if stmt == nil {
			tool.PanicDebug("adding nil to expression list: ", ctx)
		}
		return []node.Node{stmt}
	} else if ctx.LocalTypeDeclaration() != nil {
		panic("didn't anticipate this")
	}

	panic("unknown block statement type")
}
