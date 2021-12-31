/* AUTO-GENERATED: do not edit */
package visitor

import (
	"fmt"

	"github.com/dragonfax/java_converter/trans/ast"
	"github.com/dragonfax/java_converter/trans/node"
)

func (av *BaseASTVisitor[T]) VisitNode(tree node.Node) T {
	if tree == nil {
		fmt.Printf("someone gave us a nil node to visit\n")
		return av.zero
	}

	if n, ok := tree.(*ast.ArrayLiteral); ok {
		return av.root.VisitArrayLiteral(n)

	} else if n, ok := tree.(*ast.Block); ok {
		return av.root.VisitBlock(n)

	} else if n, ok := tree.(*ast.Class); ok {
		return av.root.VisitClass(n)

	} else if n, ok := tree.(*ast.ClassRef); ok {
		return av.root.VisitClassRef(n)

	} else if n, ok := tree.(*ast.EnumConstant); ok {
		return av.root.VisitEnumConstant(n)

	} else if n, ok := tree.(*ast.EnumRef); ok {
		return av.root.VisitEnumRef(n)

	} else if n, ok := tree.(*ast.BaseExpression); ok {
		return av.root.VisitBaseExpression(n)

	} else if n, ok := tree.(*ast.FieldList); ok {
		return av.root.VisitFieldList(n)

	} else if n, ok := tree.(*ast.FieldRef); ok {
		return av.root.VisitFieldRef(n)

	} else if n, ok := tree.(*ast.EnhancedFor); ok {
		return av.root.VisitEnhancedFor(n)

	} else if n, ok := tree.(*ast.ClassicFor); ok {
		return av.root.VisitClassicFor(n)

	} else if n, ok := tree.(*ast.Hierarchy); ok {
		return av.root.VisitHierarchy(n)

	} else if n, ok := tree.(*ast.If); ok {
		return av.root.VisitIf(n)

	} else if n, ok := tree.(*ast.Import); ok {
		return av.root.VisitImport(n)

	} else if n, ok := tree.(*ast.Lambda); ok {
		return av.root.VisitLambda(n)

	} else if n, ok := tree.(*ast.Literal); ok {
		return av.root.VisitLiteral(n)

	} else if n, ok := tree.(*ast.LocalVarDecl); ok {
		return av.root.VisitLocalVarDecl(n)

	} else if n, ok := tree.(*ast.Method); ok {
		return av.root.VisitMethod(n)

	} else if n, ok := tree.(*ast.MethodCall); ok {
		return av.root.VisitMethodCall(n)

	} else if n, ok := tree.(*ast.MethodRef); ok {
		return av.root.VisitMethodRef(n)

	} else if n, ok := tree.(*ast.Package); ok {
		return av.root.VisitPackage(n)

	} else if n, ok := tree.(*ast.Return); ok {
		return av.root.VisitReturn(n)

	} else if n, ok := tree.(*ast.Throw); ok {
		return av.root.VisitThrow(n)

	} else if n, ok := tree.(*ast.Break); ok {
		return av.root.VisitBreak(n)

	} else if n, ok := tree.(*ast.Continue); ok {
		return av.root.VisitContinue(n)

	} else if n, ok := tree.(*ast.Label); ok {
		return av.root.VisitLabel(n)

	} else if n, ok := tree.(*ast.Switch); ok {
		return av.root.VisitSwitch(n)

	} else if n, ok := tree.(*ast.SwitchCase); ok {
		return av.root.VisitSwitchCase(n)

	} else if n, ok := tree.(*ast.SynchronizedBlock); ok {
		return av.root.VisitSynchronizedBlock(n)

	} else if n, ok := tree.(*ast.TryCatch); ok {
		return av.root.VisitTryCatch(n)

	} else if n, ok := tree.(*ast.CatchClause); ok {
		return av.root.VisitCatchClause(n)

	} else if n, ok := tree.(*ast.TypePath); ok {
		return av.root.VisitTypePath(n)

	} else if n, ok := tree.(*ast.TypeElement); ok {
		return av.root.VisitTypeElement(n)

	} else if n, ok := tree.(*ast.TypeParameterList); ok {
		return av.root.VisitTypeParameterList(n)

	} else if n, ok := tree.(*ast.TypeParameter); ok {
		return av.root.VisitTypeParameter(n)

	} else if n, ok := tree.(*ast.VarRef); ok {
		return av.root.VisitVarRef(n)

	} else {
		return av.root.VisitChildren(tree)
	}
}

func (av *BaseASTVisitor[T]) VisitArrayLiteral(tree *ast.ArrayLiteral) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitBlock(tree *ast.Block) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitClass(tree *ast.Class) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitClassRef(tree *ast.ClassRef) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitEnumConstant(tree *ast.EnumConstant) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitEnumRef(tree *ast.EnumRef) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitBaseExpression(tree *ast.BaseExpression) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitFieldList(tree *ast.FieldList) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitFieldRef(tree *ast.FieldRef) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitEnhancedFor(tree *ast.EnhancedFor) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitClassicFor(tree *ast.ClassicFor) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitHierarchy(tree *ast.Hierarchy) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitIf(tree *ast.If) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitImport(tree *ast.Import) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitLambda(tree *ast.Lambda) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitLiteral(tree *ast.Literal) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitLocalVarDecl(tree *ast.LocalVarDecl) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitMethod(tree *ast.Method) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitMethodCall(tree *ast.MethodCall) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitMethodRef(tree *ast.MethodRef) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitPackage(tree *ast.Package) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitReturn(tree *ast.Return) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitThrow(tree *ast.Throw) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitBreak(tree *ast.Break) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitContinue(tree *ast.Continue) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitLabel(tree *ast.Label) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitSwitch(tree *ast.Switch) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitSwitchCase(tree *ast.SwitchCase) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitSynchronizedBlock(tree *ast.SynchronizedBlock) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitTryCatch(tree *ast.TryCatch) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitCatchClause(tree *ast.CatchClause) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitTypePath(tree *ast.TypePath) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitTypeElement(tree *ast.TypeElement) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitTypeParameterList(tree *ast.TypeParameterList) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitTypeParameter(tree *ast.TypeParameter) T {
	return av.root.VisitChildren(tree)
}

func (av *BaseASTVisitor[T]) VisitVarRef(tree *ast.VarRef) T {
	return av.root.VisitChildren(tree)
}
