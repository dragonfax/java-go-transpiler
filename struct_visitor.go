/* visit a parse tree and build a golang struct from the java fields within the class in the file */
package main

import (
	"github.com/dragonfax/java_converter/ast"
	"github.com/dragonfax/java_converter/parser"
)

type StructVisitor struct{}

func (sv *StructVisitor) AggregateResult(aggregate ast.Node, nextResult ast.Node) ast.Node {
	/* 1. drop nils
	 * 2. merge FieldLists and Fields
	 */

	if aggregate == nil {
		return nextResult
	}

	if nextResult == nil {
		return aggregate
	}

	// this should be replaced with visitMemberDeclaration(),
	// but I wanted an example aggregate function.
	aggFieldList, ok := aggregate.(ast.FieldListNode)
	nextFieldList, ok2 := nextResult.(ast.FieldListNode)
	if ok && ok2 {
		aggFieldList = append(aggFieldList, nextFieldList...)
		return aggFieldList
	}

	return super.aggregateResult(aggregate, nextResult)
}

func (sv *StructVisitor) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) ast.Node {

	className := ctx.IDENTIFIER().GetText()

	fieldsList := super.visitClassBody(ctx.ClassBody())

	return &ast.ClassNode{Name: className, Fields: fieldsList}
}

func (sv *StructVisitor) VisitFieldDeclaration(ctx *parser.FieldDeclarationContext) ast.Node {

	typ := ctx.TypeType().GetText()

	varDecls := ctx.VariableDeclarators().(*parser.VariableDeclaratorsContext).AllVariableDeclarator()

	fieldList := make([]*ast.FieldNode, len(varDecls), 0)
	for _, varDecl := range varDecls {
		name := varDecl.(*parser.VariableDeclaratorContext).VariableDeclaratorId().GetText()
		fieldList = append(fieldList, &ast.FieldNode{Name: name, Type: typ})
	}

	return ast.FieldListNode(fieldList)
}
