package visitor

import (
	"github.com/dragonfax/java_converter/input/ast"
	"github.com/dragonfax/java_converter/input/ast/exp"
	"github.com/dragonfax/java_converter/input/parser"
)

type GoVisitor struct {
	*parser.BaseJavaParserVisitor[ast.Node]
}

func NewGoVisitor() *GoVisitor {
	this := &GoVisitor{}
	this.BaseJavaParserVisitor = parser.NewBaseJavaParserVisitor[ast.Node](this)
	return this
}

func (gv *GoVisitor) AggregateResult(aggregate, nextResult ast.Node) ast.Node {
	if aggregate == nil {
		return nextResult
	}

	if nextResult == nil {
		return aggregate
	}

	return gv.BaseJavaParserVisitor.AggregateResult(aggregate, nextResult)
}

func (gv *GoVisitor) VisitCompilationUnit(ctx *parser.CompilationUnitContext) ast.Node {
	file := ast.NewFile()

	file.PackageName = ctx.PackageDeclaration().QualifiedName().GetText()

	for _, importCtx := range ctx.AllImportDeclaration() {
		file.Imports = append(file.Imports, ast.NewImport(importCtx.QualifiedName().GetText()))
	}

	class := gv.VisitChildren(ctx)
	if class != nil {
		file.Classes = append(file.Classes, class.(*ast.Class))
	}

	return file
}

func (gv *GoVisitor) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) ast.Node {
	class := gv.VisitChildren(ctx).(*ast.Class)

	class.Name = ctx.IDENTIFIER().GetText()

	if ctx.TypeType() != nil {
		class.BaseClass = ctx.TypeType().GetText()
	}

	if ctx.TypeList() != nil {
		for _, typeType := range ctx.TypeList().AllTypeType() {
			class.Interfaces = append(class.Interfaces, exp.NewTypeNode(typeType))
		}
	}

	for _, member := range class.Members {
		if setClass, ok := member.(ast.HasSetClass); ok {
			setClass.SetClass(class.Name)
		}
	}

	return class
}

func (gv *GoVisitor) VisitClassBody(ctx *parser.ClassBodyContext) ast.Node {

	class := ast.NewClass()

	for _, decl := range ctx.AllClassBodyDeclaration() {
		member := gv.VisitClassBodyDeclaration(decl)
		if fl, ok := member.(ast.FieldList); ok {
			class.Fields = append(class.Fields, fl...)
		} else {
			class.Members = append(class.Members, member)
		}
	}

	return class
}

func (gv *GoVisitor) VisitClassBodyDeclaration(ctx *parser.ClassBodyDeclarationContext) ast.Node {

	member := gv.VisitChildren(ctx)
	if member == nil {
		return nil
	}

	lastModifier := "private" // java default
	for _, modifier := range ctx.AllModifier() {
		modifierText := modifier.GetText()
		if modifierText == "public" || modifierText == "private" {
			// these are all we care about right now.
			// for each class member
			lastModifier = modifierText
			break
		}
	}

	if setModifier, ok := member.(ast.HasSetModifier); ok {
		setModifier.SetModifier(lastModifier)
	}

	return member
}

func (gv *GoVisitor) VisitMethodDeclaration(ctx *parser.MethodDeclarationContext) ast.Node {
	name := ctx.IDENTIFIER().GetText()

	body := exp.NewBlockNode(ctx.MethodBody().Block())
	m := ast.NewMethod("", name, "", exp.FormalParameterListProcessor(ctx.FormalParameters().FormalParameterList()), ctx.TypeTypeOrVoid().GetText(), body)

	// TODO notify the method of its class name (or give it a back ref or something)

	return m
}

func (v *GoVisitor) VisitFieldDeclaration(ctx *parser.FieldDeclarationContext) ast.Node {
	return ast.NewFields(ctx)
}

func (v *GoVisitor) VisitConstantDeclaration(ctx *parser.ConstructorDeclarationContext) ast.Node {

	c := ast.NewConstructor()
	c.Name = ctx.IDENTIFIER().GetText()
	c.Body = exp.NewBlockNode(ctx.Block())

	return c
}
