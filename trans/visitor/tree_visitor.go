package visitor

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/trans/ast"
	"github.com/dragonfax/java_converter/trans/node"
)

type TreeVisitor struct {
	*parser.BaseJavaParserVisitor[node.Node]
}

func NewTreeVisitor() *TreeVisitor {
	this := &TreeVisitor{}
	this.BaseJavaParserVisitor = parser.NewBaseJavaParserVisitor[node.Node](this)
	return this
}

func (gv *TreeVisitor) AggregateResult(aggregate, nextResult node.Node) node.Node {
	if aggregate == nil {
		return nextResult
	}

	if nextResult == nil {
		return aggregate
	}

	return gv.BaseJavaParserVisitor.AggregateResult(aggregate, nextResult)
}

func (gv *TreeVisitor) VisitCompilationUnit(ctx *parser.CompilationUnitContext) node.Node {

	packageName := ctx.PackageDeclaration().QualifiedName().GetText()

	imports := make([]*ast.Import, 0)
	for _, importCtx := range ctx.AllImportDeclaration() {
		importedPackageName := importCtx.QualifiedName().GetText()
		imports = append(imports, ast.NewImport(importedPackageName))
	}

	if len(ctx.AllTypeDeclaration()) == 0 {
		return nil
	}

	if len(ctx.AllTypeDeclaration()) > 1 {
		panic(fmt.Sprintf("wrong number of types in a file %d", 0))
	}

	node := gv.VisitTypeDeclaration(ctx.TypeDeclaration(0))
	if node == nil {
		return nil
	}

	class := node.(*ast.Class)
	class.PackageName = packageName
	class.Imports = imports

	return class
}

func (gv *TreeVisitor) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) node.Node {
	class := gv.VisitChildren(ctx).(*ast.Class)

	class.Name = ctx.IDENTIFIER().GetText()

	if ctx.TypeType() != nil {
		class.BaseClassName = ctx.TypeType().GetText()
	}

	if ctx.TypeList() != nil {
		for _, typeType := range ctx.TypeList().AllTypeType() {
			class.Interfaces = append(class.Interfaces, ast.NewTypeNodeFromContext(typeType))
		}
	}

	return class
}

func (gv *TreeVisitor) VisitClassBody(ctx *parser.ClassBodyContext) node.Node {

	class := ast.NewClass()

	for _, decl := range ctx.AllClassBodyDeclaration() {
		member := gv.VisitClassBodyDeclaration(decl)
		if member == nil {
			if decl.GetText() != ";" {
				fmt.Printf("WARNING: skipping class member: %s\n", decl.GetText())
			}
			continue
		}
		if subClass, ok := member.(*ast.Class); ok {
			class.OtherMembers = append(class.OtherMembers, subClass)
		} else if f, ok := member.(*ast.Field); ok {
			class.Fields = append(class.Fields, f)
		} else if fl, ok := member.(*ast.FieldList); ok {
			class.Fields = append(class.Fields, fl.Fields...)
		} else if m, ok := member.(*ast.Member); ok {
			class.Members = append(class.Members, m)
		} else {
			class.OtherMembers = append(class.OtherMembers, member)
		}
	}

	return class
}

func (gv *TreeVisitor) VisitClassBodyDeclaration(ctx *parser.ClassBodyDeclarationContext) node.Node {

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

func (gv *TreeVisitor) VisitGenericMethodDeclaration(ctx *parser.GenericMethodDeclarationContext) node.Node {

	method := gv.VisitMethodDeclaration(ctx.MethodDeclaration())
	method.(*ast.Member).TypeParameters = ast.NewTypeParameterList(ctx.TypeParameters())

	return method
}

func (gv *TreeVisitor) VisitGenericConstructorDeclaration(ctx *parser.GenericConstructorDeclarationContext) node.Node {
	panic("generic constructor declaration")
}

func (gv *TreeVisitor) VisitMethodDeclaration(ctx *parser.MethodDeclarationContext) node.Node {
	name := ctx.IDENTIFIER().GetText()

	body := ast.NewBlock(ctx.MethodBody().Block())
	m := ast.NewMethod(name, ast.FormalParameterListProcessor(ctx.FormalParameters().FormalParameterList()), ast.NewTypeOrVoid(ctx.TypeTypeOrVoid()), body)

	if ctx.THROWS() != nil {
		m.Throws = ctx.QualifiedNameList().GetText()
	}

	return m
}

func (v *TreeVisitor) VisitFieldDeclaration(ctx *parser.FieldDeclarationContext) node.Node {
	return ast.NewFieldList(ctx)
}

func (v *TreeVisitor) VisitConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) node.Node {
	c := ast.NewConstructor(ctx)
	return c
}

func (gv *TreeVisitor) VisitEnumDeclaration(ctx *parser.EnumDeclarationContext) node.Node {
	members := make([]*ast.Member, 0)
	fields := make([]*ast.Field, 0)
	otherMembers := make([]node.Node, 0)
	if ctx.EnumBodyDeclarations() != nil {
		for _, decl := range ctx.EnumBodyDeclarations().AllClassBodyDeclaration() {
			member := gv.VisitClassBodyDeclaration(decl)
			if member == nil {
				if decl.GetText() != ";" {
					fmt.Printf("WARNING: skipping class member: %s\n", decl.GetText())
				}
				continue
			}
			if subClass, ok := member.(*ast.Class); ok {
				// We don't do subclasses
				otherMembers = append(otherMembers, subClass)
			} else if f, ok := member.(*ast.Field); ok {
				fields = append(fields, f)
			} else if fl, ok := member.(*ast.FieldList); ok {
				fields = append(fields, fl.Fields...)
			} else if m, ok := member.(*ast.Member); ok {
				members = append(members, m)
			} else {
				otherMembers = append(otherMembers, member)
			}
		}
	}

	return ast.NewEnum(ctx, fields, members, otherMembers)
}

func (gv *TreeVisitor) VisitInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) node.Node {
	return ast.NewInterface(ctx)
}
