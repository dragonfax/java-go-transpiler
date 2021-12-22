package ast

import (
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/node"
)

type LiteralType int

const (
	Integer LiteralType = iota
	Float
	Char
	String
	Bool
	Null
)

type Literal struct {
	*node.Base

	Type  LiteralType
	Value string
}

func (ln *Literal) Children() []node.Node {
	return nil
}

func NewLiteral(literal *parser.LiteralContext) *Literal {
	if literal == nil {
		panic("no literal value")
	}
	literalCtx := literal

	var typ LiteralType
	if literalCtx.CHAR_LITERAL() != nil {
		typ = Char
	}
	if literalCtx.STRING_LITERAL() != nil {
		typ = String
	}
	if literalCtx.BOOL_LITERAL() != nil {
		typ = Bool
	}
	if literalCtx.NULL_LITERAL() != nil {
		typ = Null
	}
	if literalCtx.FloatLiteral() != nil {
		typ = Float
	}
	if literalCtx.IntegerLiteral() != nil {
		typ = Integer
	}

	return &Literal{
		Base:  node.New(),
		Type:  typ,
		Value: literal.GetText(),
	}
}

func (ln *Literal) String() string {
	if ln.Type == Float {
		return strings.TrimSuffix(ln.Value, "f")
	}
	if ln.Type == Null {
		return "nil"
	}
	return ln.Value
}