package ast

import (
	"strings"

	"github.com/dragonfax/java-go-transpiler/input/parser"
	"github.com/dragonfax/java-go-transpiler/trans/node"
)

type LiteralType string

// not part of primitive clases as these have to be available during parsing.
const (
	Integer LiteralType = "int"
	Short   LiteralType = "short"
	Long    LiteralType = "long"
	Float   LiteralType = "float"
	Double  LiteralType = "double"
	Char    LiteralType = "char"
	String  LiteralType = "string"
	Bool    LiteralType = "boolean"
	Null    LiteralType = "null"
)

type Literal struct {
	*BaseExpression

	LiteralType LiteralType
	Value       string
}

func (l *Literal) GetType() *Class {
	// will get its type from an early visitor pass
	return l.Type
}

func (ln *Literal) Children() []node.Node {
	return nil
}

func NewLiteral(typ LiteralType, value string) *Literal {
	return &Literal{
		BaseExpression: NewExpression(),
		LiteralType:    typ,
		Value:          value,
	}
}

func NewLiteralFromContext(literal *parser.LiteralContext) *Literal {
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
		BaseExpression: NewExpression(),
		LiteralType:    typ,
		Value:          literal.GetText(),
	}
}

func (ln *Literal) NodeName() string {
	return ln.String()
}

func (ln *Literal) String() string {
	if ln.LiteralType == Float {
		return strings.TrimSuffix(ln.Value, "f")
	}
	if ln.LiteralType == Null {
		return "nil"
	}
	return ln.Value
}
