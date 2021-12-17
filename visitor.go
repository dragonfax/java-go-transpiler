/* Instead of a baseclass for visitor.
 * This struct would manage the tree walking visit process.
 *
 * This file would be auto-generated by antlr.
 */
package main

import "github.com/antlr/antlr4/runtime/Go/antlr"

var _ antlr.ParseTreeVisitor = &JavaVisitor{}

type JavaVisitor struct {
	V interface{}
}

func NewJavaVisitor(v interface{}) *JavaVisitor {
	return &JavaVisitor{V: v}
}

func (jv *JavaVisitor) Visit(tree antlr.ParseTree) Node {
	return tree.Accept(this)
}
