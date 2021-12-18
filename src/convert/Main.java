package convert;

import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.*;

import convert.ast.Node;
import convert.visitor.GoStructVisitor;
import parser.*;


public class Main {
    
    private static String exampleFilename = "example/DrawableMesh.java";

    public static void main(String[] args) throws Exception {


        CharStream input = CharStreams.fromFileName(exampleFilename);
        JavaLexer lexer = new JavaLexer(input);
        CommonTokenStream tokens = new CommonTokenStream(lexer);
        JavaParser parser = new JavaParser(tokens);
        JavaParser.CompilationUnitContext tree = parser.compilationUnit(); // parse a compilationUnit
        GoStructVisitor visitor = new GoStructVisitor();

        Node ast = visitor.visit(tree);
        System.out.println(ast);

    }
}
