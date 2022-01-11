package visitor

import (
	"fmt"

	"github.com/dragonfax/java-go-transpiler/input/parser"
	"github.com/dragonfax/java-go-transpiler/trans/ast"
	"github.com/dragonfax/java-go-transpiler/trans/node"
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

	packageName := "." // default package
	if ctx.PackageDeclaration() != nil {
		packageName = ctx.PackageDeclaration().QualifiedName().GetText()
	}
	class.PackageName = packageName

	imports := make([]*ast.Import, 0)
	for _, importCtx := range ctx.AllImportDeclaration() {
		importedPackageName := importCtx.QualifiedName().GetText()
		imports = append(imports, ast.NewImport(importedPackageName))
	}
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
			class.NestedClasses = append(class.NestedClasses, subClass)
		} else if f, ok := member.(*ast.Field); ok {
			class.Fields = append(class.Fields, f)
		} else if fl, ok := member.(*ast.FieldList); ok {
			class.Fields = append(class.Fields, fl.Fields...)
		} else if m, ok := member.(*ast.Method); ok {
			class.Methods = append(class.Methods, m)
		} else {
			panic(fmt.Sprintf("warning: other members: (%T)\n%s\n\n", member, decl.GetText()))
		}
	}

	return class
}

func (gv *TreeVisitor) VisitClassBodyDeclaration(ctx *parser.ClassBodyDeclarationContext) node.Node {

	// static and non-static initializers
	// wont' be processed by any of my other visit rules
	if ctx.Block() != nil {
		// acts as a method
		return ast.NewInitializerBlock(ctx)
	}

	/* collect children and notify them of their modifiers */

	method := gv.VisitChildren(ctx)
	if method == nil {
		return nil
	}

	for _, modifier := range ctx.AllModifier() {
		modifierText := modifier.GetText()
		switch modifierText {
		case "public", "protected":
			if set, ok := method.(interface{ SetPublic(bool) }); ok {
				set.SetPublic(true)
			}
		case "transient":
			if set, ok := method.(interface{ SetTransient(bool) }); ok {
				set.SetTransient(true)
			}
		case "static":
			if set, ok := method.(interface{ SetStatic(bool) }); ok {
				set.SetStatic(true)
			}
		case "abstract":
			if set, ok := method.(interface{ SetAbstract(bool) }); ok {
				set.SetAbstract(true)
			}
		case "synchronized":
			if set, ok := method.(interface{ SetSynchronized(bool) }); ok {
				set.SetSynchronized((true))
			}
		}
	}

	return method
}

func (gv *TreeVisitor) VisitGenericMethodDeclaration(ctx *parser.GenericMethodDeclarationContext) node.Node {

	method := gv.VisitMethodDeclaration(ctx.MethodDeclaration())
	method.(*ast.Method).TypeParameters = ast.NewTypeParameterList(ctx.TypeParameters())

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
	methods := make([]*ast.Method, 0)
	fields := make([]*ast.Field, 0)
	nestedClasses := make([]*ast.Class, 0)
	if ctx.EnumBodyDeclarations() != nil {
		for _, decl := range ctx.EnumBodyDeclarations().AllClassBodyDeclaration() {
			method := gv.VisitClassBodyDeclaration(decl)
			if method == nil {
				if decl.GetText() != ";" {
					fmt.Printf("WARNING: skipping class method: %s\n", decl.GetText())
				}
				continue
			}
			if subClass, ok := method.(*ast.Class); ok {
				// We don't do subclasses
				nestedClasses = append(nestedClasses, subClass)
			} else if f, ok := method.(*ast.Field); ok {
				fields = append(fields, f)
			} else if fl, ok := method.(*ast.FieldList); ok {
				fields = append(fields, fl.Fields...)
			} else if m, ok := method.(*ast.Method); ok {
				methods = append(methods, m)
			} else {
				panic("unknown member type adding to enum as a nested class")
				// nestedClasses = append(nestedClasses, method)
			}
		}
	}

	return ast.NewEnum(ctx, fields, methods, nestedClasses)
}

func (gv *TreeVisitor) VisitInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) node.Node {
	return ast.NewInterface(ctx)
}
