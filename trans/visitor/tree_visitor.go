package visitor

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/ast"
	"github.com/dragonfax/java_converter/trans/ast/exp"
	"github.com/dragonfax/java_converter/trans/hier"
)

type TreeVisitor struct {
	*parser.BaseJavaParserVisitor[ast.Node]

	Hierarchy *hier.Hierarchy
}

func NewTreeVisitor(h *hier.Hierarchy) *TreeVisitor {
	this := &TreeVisitor{Hierarchy: h}
	this.BaseJavaParserVisitor = parser.NewBaseJavaParserVisitor[ast.Node](this)
	return this
}

func (gv *TreeVisitor) AggregateResult(aggregate, nextResult ast.Node) ast.Node {
	if aggregate == nil {
		return nextResult
	}

	if nextResult == nil {
		return aggregate
	}

	return gv.BaseJavaParserVisitor.AggregateResult(aggregate, nextResult)
}

func (gv *TreeVisitor) VisitCompilationUnit(ctx *parser.CompilationUnitContext) ast.Node {

	node := gv.VisitChildren(ctx)
	if node == nil {
		return nil
	}

	if _, ok := node.(*ast.Enum); ok {
		panic("top level enum detected")
	}

	class := node.(*ast.Class)
	class.Package = ctx.PackageDeclaration().QualifiedName().GetText()

	for _, importCtx := range ctx.AllImportDeclaration() {
		importedPackageName := importCtx.QualifiedName().GetText()
		class.Imports = append(class.Imports, ast.NewImport(importedPackageName))
		gv.Hierarchy.AddClass(importedPackageName, class)
	}

	return class
}

func (gv *TreeVisitor) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) ast.Node {
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

func (gv *TreeVisitor) VisitClassBody(ctx *parser.ClassBodyContext) ast.Node {

	class := ast.NewClass()

	for _, decl := range ctx.AllClassBodyDeclaration() {
		member := gv.VisitClassBodyDeclaration(decl)
		if member == nil && decl.GetText() != ";" {
			fmt.Printf("WARNING: skipping class member: %s\n", decl.GetText())
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

func (gv *TreeVisitor) VisitClassBodyDeclaration(ctx *parser.ClassBodyDeclarationContext) ast.Node {

	// static and non-static initializers
	// wont' be processed by any of my other visit rules

	if ctx.Block() != nil {
		// acts as a member
		return ast.NewInitializerBlock(ctx)
	}

	/* collect children and notify them of their modifiers */

	member := gv.VisitChildren(ctx)
	if member == nil {
		return nil
	}

	for _, modifier := range ctx.AllModifier() {
		modifierText := modifier.GetText()
		switch modifierText {
		case "public", "protected":
			if set, ok := member.(interface{ SetPublic(bool) }); ok {
				set.SetPublic(true)
			}
		case "transient":
			if set, ok := member.(interface{ SetTransient(bool) }); ok {
				set.SetTransient(true)
			}
		case "static":
			if set, ok := member.(interface{ SetStatic(bool) }); ok {
				set.SetStatic(true)
			}
		case "abstract":
			if set, ok := member.(interface{ SetAbstract(bool) }); ok {
				set.SetAbstract(true)
			}
		case "synchronized":
			if set, ok := member.(interface{ SetSynchronized(bool) }); ok {
				set.SetSynchronized((true))
			}
		}
	}

	return member
}

func (gv *TreeVisitor) VisitMethodDeclaration(ctx *parser.MethodDeclarationContext) ast.Node {
	name := ctx.IDENTIFIER().GetText()

	body := exp.NewBlockNode(ctx.MethodBody().Block())
	m := ast.NewMethod(name, "", exp.FormalParameterListProcessor(ctx.FormalParameters().FormalParameterList()), exp.NewTypeOrVoidNode(ctx.TypeTypeOrVoid()), body)

	if ctx.THROWS() != nil {
		m.Throws = ctx.QualifiedNameList().GetText()
	}

	return m
}

func (v *TreeVisitor) VisitFieldDeclaration(ctx *parser.FieldDeclarationContext) ast.Node {
	return ast.NewFields(ctx)
}

func (v *TreeVisitor) VisitConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) ast.Node {
	c := ast.NewConstructor(ctx)
	return c
}

func (gv *TreeVisitor) VisitEnumDeclaration(ctx *parser.EnumDeclarationContext) ast.Node {
	return ast.NewEnum(ctx)
}
