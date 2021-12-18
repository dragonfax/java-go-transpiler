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
    
        if ( aggregate instanceof FieldListNode && nextResult instanceof FieldListNode ) {
            return ((FieldListNode)aggregate).append((FieldListNode)nextResult);
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
    
        var type = ctx.typeType().getText();

        var varDecls = ctx.variableDeclarators().variableDeclarator();

        var fieldList = new ArrayList<FieldNode>(varDecls.size());
        for ( var varDecl : varDecls )  {
            var name = varDecl.variableDeclaratorId().getText();
    
            fieldList.add(new FieldNode( name, type ));
        }
    
        return new FieldListNode(fieldList.toArray(new FieldNode[]{}));
    }
    
}
