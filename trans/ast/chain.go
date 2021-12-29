package ast

import (
	"strings"

	"github.com/dragonfax/java_converter/trans/node"
)

/* A chain of DOT binary operators. a series of field/method references together.
 * In the initial build of the AST, these turn into a tree of DOT BinaryOperator nodes.
 * But that is hard to reason about, so we use a visitor to reduce those trees
 * into individual Chain nodes.
 */
type Chain struct {
	*node.Base

	Elements []Expression
}

/* given the root of a tree of DOT binary operators, create a single chain out of them and replace it in their parent. */
func NewChain(left, right Expression) *Chain {

	// right is never another BOP or Chain, right is always a terminal of some form.

	var elements []Expression
	if subChain, ok := left.(*Chain); ok {
		// merge this chain with the sub chain
		elements = append(subChain.Elements, right)
	} else {
		elements = []Expression{left, right}
	}

	this := &Chain{
		Base:     node.New(),
		Elements: elements,
	}

	return this
}

func (c *Chain) String() string {
	return strings.Join(node.NodeListToStringList(c.Elements), ".")
}

func (c *Chain) Children() []node.Node {
	return c.Elements

}
