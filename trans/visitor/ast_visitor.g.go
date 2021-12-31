/* AUTO-GENERATED: do not edit */
package visitor

import (
	"github.com/dragonfax/java_converter/trans/ast"
)

type GenASTVisitor[T comparable] interface {
	VisitArrayLiteral(tree *ast.ArrayLiteral) T
	VisitBlock(tree *ast.Block) T
	VisitClass(tree *ast.Class) T
	VisitClassRef(tree *ast.ClassRef) T
	VisitEnumConstant(tree *ast.EnumConstant) T
	VisitEnumRef(tree *ast.EnumRef) T
	VisitField(tree *ast.Field) T
	VisitFieldList(tree *ast.FieldList) T
	VisitFieldRef(tree *ast.FieldRef) T
	VisitEnhancedFor(tree *ast.EnhancedFor) T
	VisitClassicFor(tree *ast.ClassicFor) T
	VisitHierarchy(tree *ast.Hierarchy) T
	VisitIf(tree *ast.If) T
	VisitImport(tree *ast.Import) T
	VisitLambda(tree *ast.Lambda) T
	VisitLiteral(tree *ast.Literal) T
	VisitLocalVarDecl(tree *ast.LocalVarDecl) T
	VisitMethod(tree *ast.Method) T
	VisitMethodCall(tree *ast.MethodCall) T
	VisitMethodRef(tree *ast.MethodRef) T
	VisitBinaryOperator(tree *ast.BinaryOperator) T
	VisitUnaryOperator(tree *ast.UnaryOperator) T
	VisitTernaryOperator(tree *ast.TernaryOperator) T
	VisitPackage(tree *ast.Package) T
	VisitReturn(tree *ast.Return) T
	VisitThrow(tree *ast.Throw) T
	VisitBreak(tree *ast.Break) T
	VisitContinue(tree *ast.Continue) T
	VisitLabel(tree *ast.Label) T
	VisitSwitch(tree *ast.Switch) T
	VisitSwitchCase(tree *ast.SwitchCase) T
	VisitSynchronizedBlock(tree *ast.SynchronizedBlock) T
	VisitTryCatch(tree *ast.TryCatch) T
	VisitCatchClause(tree *ast.CatchClause) T
	VisitTypePath(tree *ast.TypePath) T
	VisitTypeElement(tree *ast.TypeElement) T
	VisitTypeParameterList(tree *ast.TypeParameterList) T
	VisitTypeParameter(tree *ast.TypeParameter) T
	VisitIdentRef(tree *ast.IdentRef) T
}
