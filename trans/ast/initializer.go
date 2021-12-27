package ast

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

/* static and non-static initializers.
 * act as methods of a class.
 */
type Initializer struct {
	*node.Base

	Static bool
	Block  *Block
}

func NewInitializerBlock(ctx *parser.ClassBodyDeclarationContext) *Initializer {
	this := &Initializer{
		Base:   node.New(),
		Static: ctx.Block() != nil,
		Block:  NewBlock(ctx.Block()),
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
