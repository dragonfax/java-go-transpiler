package convert.visitor;

import convert.ast.*;
import parser.*;
import java.util.ArrayList;

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

        var aggFieldOk = aggregate instanceof FieldNode;
        var nextFieldOk = nextResult instanceof FieldNode;
        if (aggFieldOk &&  nextFieldOk ) {
            return new FieldNode(((FieldNode)nextResult).name, ((FieldNode)aggregate).type);
        }
    
        FieldListNode aggFieldList = null;
        FieldListNode nextFieldList = null;
        if (aggregate instanceof FieldListNode) {
            aggFieldList = (FieldListNode)aggregate;
        }
        if ( nextResult instanceof FieldListNode) {
            nextFieldList = (FieldListNode)nextResult;
        }
    
        if ( aggFieldList != null &&  nextFieldList != null )  {
            return aggFieldList.append(nextFieldList);
        }

        if ( aggFieldList != null && nextFieldOk ) {
            return aggFieldList.append(new FieldListNode((FieldNode)nextResult));
        }
    
        throw new RuntimeException(String.format("unknown aggregation situation: %s(%s) and %s(%s)", aggregate.getClass(), aggregate, nextResult.getClass(), nextResult));
    }


    @Override
    public Node visitClassDeclaration(JavaParser.ClassDeclarationContext ctx) {

        var className = ctx.IDENTIFIER().getText();
    
        var fieldsList = (FieldListNode)visitClassBody(ctx.classBody());
    
        return new ClassNode(className, fieldsList);
    }
    
    /* defaultResult is just null */
    

    @Override
    public Node visitFieldDeclaration(JavaParser.FieldDeclarationContext ctx) {
    
        var type = ((FieldNode)visitTypeType(ctx.typeType())).type;

        var varDecls = ctx.variableDeclarators().variableDeclarator();

        var fieldList = new ArrayList<FieldNode>(varDecls.size());
    
        for ( var varDecl : varDecls )  {
            var varDeclNode = visitVariableDeclarator(varDecl);
    
            fieldList.add(new FieldNode( ((FieldNode)varDeclNode).name, type ));
        }
    
        return new FieldListNode(fieldList.toArray(new FieldNode[]{}));
    }
    
    @Override
    public Node visitVariableDeclaratorId(JavaParser.VariableDeclaratorIdContext ctx) {
        // partial field node, just used to send part of the data up the line.
        return new FieldNode(ctx.IDENTIFIER().getText(), null);
    }
    
    @Override
    public Node visitTypeType(JavaParser.TypeTypeContext ctx) {
        // send partial field node, they get combined up the line.
    
        if ( ctx.primitiveType() != null ) {
            return new FieldNode(null, ctx.primitiveType().getText());
        }
    
        if ( ctx.classOrInterfaceType() != null ) {
            var type = ctx.classOrInterfaceType().IDENTIFIER().get(0).getText();
            return new FieldNode(null, type);
        }
    
        throw new RuntimeException("unknown");
    }
    
}
