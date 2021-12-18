package convert.visitor;

import convert.ast.*;
import parser.*;
import java.util.ArrayList;

public class GoStructVisitor extends JavaParserBaseVisitor<Node> {

    @Override
    protected  Node aggregateResult(Node aggregate, Node nextResult) {
        /* 1. drop nils
         * 2. merge FieldLists and Fields
         */

        if (aggregate == null) {
            return nextResult;
        }

        if (nextResult == null) {
            return aggregate;
        }
    
        if (! (aggregate instanceof FieldListNode )) {
            return super.aggregateResult(aggregate, nextResult);
        }

        var aggFieldList = (FieldListNode)aggregate;

        if ( nextResult instanceof FieldListNode) {
            return aggFieldList.append((FieldListNode)nextResult);
        } else if ( nextResult instanceof FieldNode ) {
            return aggFieldList.append(new FieldListNode((FieldNode)nextResult));
        }
    
        return super.aggregateResult(aggregate, nextResult);
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
