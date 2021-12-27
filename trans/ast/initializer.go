package ast

import (
	"github.com/dragonfax/java_converter/input/parser"
)

/* static and non-static initializers.
 * act as methods of a class.
 */
func NewInitializerBlock(ctx *parser.ClassBodyDeclarationContext) *Method {
	initializer := NewMethod("init", nil, nil, NewBlock(ctx.Block()))
	initializer.Initializer = true
	initializer.Static = ctx.STATIC() != nil
	return initializer
}
