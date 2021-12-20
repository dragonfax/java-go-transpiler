package exp

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
)

type SynchronizedBlock struct {
	Condition ExpressionNode
	Block     *BlockNode
}

func NewSynchronizedBlock(ctx *parser.StatementContext) *SynchronizedBlock {
	s := &SynchronizedBlock{
		Condition: ExpressionProcessor(ctx.ParExpression().Expression()),
		Block:     NewBlockNode(ctx.Block()),
	}

	return s
}

func (sb *SynchronizedBlock) String() string {
	return fmt.Sprintf("// TODO synchronized(%s)\n%s", sb.Condition, sb.Block)
}
