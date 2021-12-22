package ast

import (
	"fmt"
	"strconv"

	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/node"
)

type Return struct {
	*node.Base

	Expression node.Node
}

func (rn *Return) Children() []node.Node {
	if rn.Expression != nil {
		return []node.Node{rn.Expression}
	}
	return nil
}

func NewReturn(exp node.Node) *Return {
	return &Return{Base: node.New(), Expression: exp}
}

func (rn *Return) String() string {
	exp := ""
	if !tool.IsNilInterface(rn.Expression) {
		exp = rn.Expression.String()
	}
	return fmt.Sprintf("return %s\n", exp)
}

type Throw struct {
	*node.Base
	Expression string
}

func (tn *Throw) Children() []node.Node {
	return nil
}

func NewThrow(exp string) *Throw {
	return &Throw{Base: node.New(), Expression: exp}
}

func (tn *Throw) String() string {
	return fmt.Sprintf("panic(%s) // TODO\n", strconv.Quote(tn.Expression))
}

type Break struct {
	*node.Base

	Label string
}

func (b *Break) Children() []node.Node {
	return nil
}

func NewBreak(label string) *Break {
	return &Break{Base: node.New(), Label: label}
}

func (bn *Break) String() string {
	return fmt.Sprintf("break %s\n", bn.Label)
}

type Continue struct {
	*node.Base

	Label string
}

func (c *Continue) Children() []node.Node {
	return nil
}

func NewContinue(label string) *Continue {
	return &Continue{Base: node.New(), Label: label}
}

func (cn *Continue) String() string {
	return fmt.Sprintf("continue %s\n", cn.Label)
}

type Label struct {
	*node.Base

	Label      string
	Expression node.Node
}

func (l *Label) Children() []node.Node {
	return []node.Node{l.Expression}
}

func NewLabel(label string, exp node.Node) *Label {
	if label == "" {
		panic("label missing")
	}
	if tool.IsNilInterface(exp) {
		panic("expression missing")
	}
	return &Label{Base: node.New(), Label: label, Expression: exp}
}

func (ln *Label) String() string {
	return fmt.Sprintf("%s: %s\n", ln.Label, ln.Expression)
}

type Identifier struct {
	*node.Base

	Identifier string
}

func (i *Identifier) Children() []node.Node {
	return nil
}

func NewIdentifier(id string) *Identifier {
	return &Identifier{Base: node.New(), Identifier: id}
}

func (in *Identifier) String() string {
	return in.Identifier
}
