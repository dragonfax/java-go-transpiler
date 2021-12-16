/* visit a parse tree and build a golang struct from the java fields within the class in the file */
package main

import "github.com/dragonfax/java_converter/parser"

type StructVisitor struct {
}

func (sv *StructVisitor) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) Node {

	className := ctx.IDENTIFIER().GetText()

	fieldsList := sv.VisitClassBody(ctx.ClassBody()).(FieldListNode).([]*FieldNode)

	return ClassNode{Name: className, Fields: fields}

}

func (sv *StructVisitor) Aggregator(aggregate, next Node) Node {
	// drop nils.
	// send a single value up the line.
	// only merge and send FieldLists
	// anything else is a panic.

	if next == nil {
		return aggregate
	}

	if aggregate == nil && next != nil {
		return next
	}

	// with this design the only time we see multiple non-nil children is FieldLists

	aggFieldList, aggOk := aggregate.(FieldListNode)
	nextFieldList, nextOk := next.(FieldListNode)

	if aggOk && nextOk {
		return append(aggFieldList, nextFieldList...)
	}

	panic("unknown")
}

/* default node is just nil */

func (sv *StructVisitor) VisitFieldDeclaraction(ctx *parser.FieldDeclarationContext) Node {

	typ := sv.VisitTypeType(ctx.TypeType()).(*FieldNode).Type

	fieldList := make([]*FieldNode, 0)
	for _, varDecl := range ctx.VariableDeclarators().(*parser.VariableDeclaratorsContext).AllVariableDeclarator() {
		varDeclNode := sv.VisitVariableDeclarator(varDecl)

		fieldList = append(fieldList,
			&FieldNode{
				Type: typ,
				Name: varDeclNode.(*FieldNode).Name,
			})
	}

	return FieldListNode(fieldList)
}

func (sv *StructVisitor) VisitVariableDeclaratorId(ctx *parser.VariableDeclaratorIdContext) Node {
	// partial field node, just used to send part of the data up the line.
	return &FieldNode{Name: ctx.IDENTIFIER().GetText()}

}

func (sv *StructVisitor) VisitTypeType(ctx *parser.TypeTypeContext) Node {
	// send partial field node, they get combined up the line.

	if ctx.PrimitiveType() != nil {
		return &FieldNode{Type: ctx.PrimitiveType().GetText()}
	}

	if ctx.ClassOrInterfaceType() != nil {
		typ := ctx.ClassOrInterfaceType().(*parser.ClassOrInterfaceTypeContext).IDENTIFIER().GetText()
		return &FieldNode{Type: typ}
	}

	panic("unknown")
}
