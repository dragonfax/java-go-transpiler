package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

/* static and non-static initializers.
 * act as members of a class.
 */
type Initializer struct {
	*node.BaseNode

	Static bool
	Block  *BlockNode
}

func NewInitializerBlock(ctx *parser.ClassBodyDeclarationContext) *Initializer {
	this := &Initializer{
		BaseNode: node.NewNode(),
		Static:   ctx.Block() != nil,
		Block:    NewBlockNode(ctx.Block()),
	}
	return this
}

func (ib *Initializer) Children() []node.Node {
	return []node.Node{ib.Block}
}

func (i *Initializer) String() string {
	if i.Static {
		return fmt.Sprintf("func init() %s", i.Block)
	} else {
		return fmt.Sprintf("// TODO join initializer with constructor\nfunc New() %s", i.Block)
	}
}
