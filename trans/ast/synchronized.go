package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

type SynchronizedBlock struct {
	*node.Base

	Condition node.Node
	Block     *Block
}

func (sb *SynchronizedBlock) Children() []node.Node {
	return []node.Node{sb.Condition, sb.Block}
}

func NewSynchronizedBlock(ctx *parser.StatementContext) *SynchronizedBlock {
	s := &SynchronizedBlock{
		Base:      node.New(),
		Condition: ExpressionProcessor(ctx.ParExpression().Expression()),
		Block:     NewBlock(ctx.Block()),
	}

	return s
}

func (sb *SynchronizedBlock) String() string {
	return fmt.Sprintf("// TODO synchronized(%s)\n%s", sb.Condition, sb.Block)
}
