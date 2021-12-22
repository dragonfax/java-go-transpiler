package ast

import (
	"fmt"
	"strconv"

	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/node"
)

type ReturnNode struct {
	Expression node.Node
}

func (rn *ReturnNode) Children() []node.Node {
	if rn.Expression != nil {
		return []node.Node{rn.Expression}
	}
	return nil
}

func NewReturnNode(exp node.Node) *ReturnNode {
	return &ReturnNode{Expression: exp}
}

func (rn *ReturnNode) String() string {
	exp := ""
	if !tool.IsNilInterface(rn.Expression) {
		exp = rn.Expression.String()
	}
	return fmt.Sprintf("return %s\n", exp)
}

type ThrowNode struct {
	Expression string
}

func (tn *ThrowNode) Children() []node.Node {
	return nil
}

func NewThrowNode(exp string) *ThrowNode {
	return &ThrowNode{Expression: exp}
}

func (tn *ThrowNode) String() string {
	return fmt.Sprintf("panic(%s) // TODO\n", strconv.Quote(tn.Expression))
}

type BreakNode struct {
	Label string
}

func (b *BreakNode) Children() []node.Node {
	return nil
}

func NewBreakNode(label string) *BreakNode {
	return &BreakNode{Label: label}
}

func (bn *BreakNode) String() string {
	return fmt.Sprintf("break %s\n", bn.Label)
}

type ContinueNode struct {
	Label string
}

func (c *ContinueNode) Children() []node.Node {
	return nil
}

func NewContinueNode(label string) *ContinueNode {
	return &ContinueNode{Label: label}
}

func (cn *ContinueNode) String() string {
	return fmt.Sprintf("continue %s\n", cn.Label)
}

type LabelNode struct {
	Label      string
	Expression node.Node
}

func (l *LabelNode) Children() []node.Node {
	return []node.Node{l.Expression}
}

func NewLabelNode(label string, exp node.Node) *LabelNode {
	if label == "" {
		panic("label missing")
	}
	if tool.IsNilInterface(exp) {
		panic("expression missing")
	}
	return &LabelNode{Label: label, Expression: exp}
}

func (ln *LabelNode) String() string {
	return fmt.Sprintf("%s: %s\n", ln.Label, ln.Expression)
}

type IdentifierNode struct {
	Identifier string
}

func (i *IdentifierNode) Children() []node.Node {
	return nil
}

func NewIdentifierNode(id string) *IdentifierNode {
	return &IdentifierNode{Identifier: id}
}

func (in *IdentifierNode) String() string {
	return in.Identifier
}
