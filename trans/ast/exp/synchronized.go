package exp

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

type SynchronizedBlock struct {
	Condition node.Node
	Block     *BlockNode
}

func (sb *SynchronizedBlock) Children() []node.Node {
	return []node.Node{sb.Condition, sb.Block}
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
