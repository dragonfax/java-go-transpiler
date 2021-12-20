package visitor

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/ast"
	"github.com/dragonfax/java_converter/trans/ast/exp"
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
		if file.Class != nil {
			panic("more than one class per file")
		}
		file.Class = class.(*ast.Class)
		file.Class.Package = file.PackageName
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
		if setClass, ok := member.(interface{ SetClass(string) }); ok {
			setClass.SetClass(class.Name)
		}
	}

	return class
}

func (gv *GoVisitor) VisitClassBody(ctx *parser.ClassBodyContext) ast.Node {

	class := ast.NewClass()

	for _, decl := range ctx.AllClassBodyDeclaration() {
		member := gv.VisitClassBodyDeclaration(decl)
		if member == nil {
			fmt.Printf("WARNING: skipping class member: %s", decl.GetText())
			continue
		}
		if subClass, ok := member.(*ast.Class); ok {
			// We don't do subclasses
			class.Members = append(class.Members, ast.NewSubClassTODO(subClass.Name))
		} else if fl, ok := member.(ast.FieldList); ok {
			class.Fields = append(class.Fields, fl...)
		} else {
			class.Members = append(class.Members, member)
		}
	}

	return class
}

func (gv *GoVisitor) VisitClassBodyDeclaration(ctx *parser.ClassBodyDeclarationContext) ast.Node {

	/* collect children and notify them of their modifiers */

	member := gv.VisitChildren(ctx)
	if member == nil {
		return nil
	}

	isPublic := false // java default
	isTransient := false
	isStatic := false
	isAbstract := false
	for _, modifier := range ctx.AllModifier() {
		modifierText := modifier.GetText()
		if modifierText == "public" || modifierText == "protected" {
			isPublic = true
		}
		if modifierText == "transient" {
			isTransient = true
		}
		if modifierText == "status" {
			isStatic = true
		}
		if modifierText == "abstract" {
			isAbstract = true
		}
	}

	if set, ok := member.(interface{ SetPublic(bool) }); isPublic && ok {
		set.SetPublic(isPublic)
	}
	if set, ok := member.(interface{ SetTransient(bool) }); isTransient && ok {
		set.SetTransient(isTransient)
	}
	if set, ok := member.(interface{ SetStatic(bool) }); isStatic && ok {
		set.SetStatic(isStatic)
	}
	if set, ok := member.(interface{ SetAbstract(bool) }); isAbstract && ok {
		set.SetAbstract(isAbstract)
	}

	return member
}

func (gv *GoVisitor) VisitMethodDeclaration(ctx *parser.MethodDeclarationContext) ast.Node {
	name := ctx.IDENTIFIER().GetText()

	body := exp.NewBlockNode(ctx.MethodBody().Block())
	m := ast.NewMethod(name, "", exp.FormalParameterListProcessor(ctx.FormalParameters().FormalParameterList()), exp.NewTypeOrVoidNode(ctx.TypeTypeOrVoid()), body)

	if ctx.THROWS() != nil {
		m.Throws = ctx.QualifiedNameList().GetText()
	}

	return m
}

func (v *GoVisitor) VisitFieldDeclaration(ctx *parser.FieldDeclarationContext) ast.Node {
	return ast.NewFields(ctx)
}

func (v *GoVisitor) VisitConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) ast.Node {
	c := ast.NewConstructor(ctx)
	return c
}
