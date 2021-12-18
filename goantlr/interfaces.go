package goantlr

import (
	"github.com/dragonfax/java_converter/ast"
	"github.com/dragonfax/java_converter/parser"
)

type HasAggregateResult interface {
	AggregateResult(aggregate ast.Node, nextResult ast.Node) ast.Node
}

type HasVisitClassDeclaration interface {
	VisitClassDeclaration(ctx *parser.ClassDeclarationContext) ast.Node
}

type HasVisitFieldDeclaration interface {
	VisitFieldDeclaration(ctx *parser.FieldDeclarationContext) ast.Node
}
