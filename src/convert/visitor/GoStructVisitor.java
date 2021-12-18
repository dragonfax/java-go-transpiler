package convert.visitor;

import convert.ast.*;
import parser.*;

public class GoStructVisitor extends JavaParserBaseVisitor<Node> {

    @Override
    protected  Node aggregateResult(Node aggregate, Node nextResult) {
        // drop nils.
        // send a single value up the line.
        // only merge and send FieldLists
        // anything else is a panic.
    
        if (nextResult == null) {
            return aggregate;
        }
    
        if (aggregate == null && nextResult != null ) {
            return nextResult;
        }
    
        // with this design the only time we see multiple non-nil children is FieldLists
    
        FieldListNode aggFieldList;
        FieldListNode nextFieldList;
        if (aggregate instanceof FieldListNode) {
            aggFieldList = (FieldListNode)aggregate;
        }
        if ( nextResult instanceof FieldListNode) {
            nextFieldList = (FieldListNode)nextResult;
        }
    
        if ( aggFieldList != null &&  nextFieldList != null )  {
            return aggFieldList.append(nextFieldList);
        }
    
        throw new RuntimeException("unknown aggregation situation");
    }


    @Override
    protected boolean shouldVisitNextChild(org.antlr.v4.runtime.tree.RuleNode node, Node currentResult) {
        return true;
    }

    @Override
    public Node visitClassDeclaration(JavaParser.ClassDeclarationContext ctx) {

        var className = ctx.IDENTIFIER().getText();
    
        var fieldsList = (FieldListNode)visitClassBody(ctx.classBody());
    
        return new ClassNode(className, fieldsList);
    }
    
    /* defaultResult is just null */
    
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
    
}
