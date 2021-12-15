package exp

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/tool"
)

type BlockNode struct {
	Body []ExpressionNode
}

func (bn *BlockNode) String() string {
	if bn == nil {
		panic("nil block")
	}
	return fmt.Sprintf("{\n%s}\n", expressionListToString(bn.Body))
}

func NewBlockNode(block parser.IBlockContext) *BlockNode {
	if tool.IsNilInterface(block) {
		return nil
	}
	ctx := block.(*parser.BlockContext)

	l := make([]ExpressionNode, 0)

	for _, blockStatement := range ctx.AllBlockStatement() {
		blockStatementCtx := blockStatement.(*parser.BlockStatementContext)

		if blockStatementCtx.LocalVariableDeclaration() != nil {
			localVarCtx := blockStatementCtx.LocalVariableDeclaration().(*parser.LocalVariableDeclarationContext)
			l = append(l, NewVariableDeclNodeList(localVarCtx)...)
		} else if blockStatementCtx.Statement() != nil {
			statementCtx := blockStatementCtx.Statement().(*parser.StatementContext)
			l = append(l, StatementProcessor(statementCtx))
		} else if blockStatementCtx.LocalTypeDeclaration() != nil {
			panic("didn't anticipate this")
		}
	}

	return &BlockNode{Body: l}
}
