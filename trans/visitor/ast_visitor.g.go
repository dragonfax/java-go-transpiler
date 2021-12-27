/* AUTO-GENERATED: do not edit */
package visitor

import (
	"github.com/dragonfax/java_converter/trans/ast"
)

type GenASTVisitor[T comparable] interface {
	VisitArrayLiteral(tree *ast.ArrayLiteral) T
	VisitBlock(tree *ast.Block) T
	VisitChain(tree *ast.Chain) T
	VisitClass(tree *ast.Class) T
	VisitClassReference(tree *ast.ClassReference) T
	VisitEnumConstant(tree *ast.EnumConstant) T
	VisitEnumRef(tree *ast.EnumRef) T
	VisitField(tree *ast.Field) T
	VisitFieldList(tree *ast.FieldList) T
	VisitFieldReference(tree *ast.FieldReference) T
	VisitEnhancedFor(tree *ast.EnhancedFor) T
	VisitClassicFor(tree *ast.ClassicFor) T
	VisitHierarchy(tree *ast.Hierarchy) T
	VisitIf(tree *ast.If) T
	VisitImport(tree *ast.Import) T
	VisitInitializer(tree *ast.Initializer) T
	VisitLambda(tree *ast.Lambda) T
	VisitLiteral(tree *ast.Literal) T
	VisitLocalVarDecl(tree *ast.LocalVarDecl) T
	VisitMember(tree *ast.Member) T
	VisitMethodCall(tree *ast.MethodCall) T
	VisitMethodReference(tree *ast.MethodReference) T
	VisitPackage(tree *ast.Package) T
	VisitReturn(tree *ast.Return) T
	VisitThrow(tree *ast.Throw) T
	VisitBreak(tree *ast.Break) T
	VisitContinue(tree *ast.Continue) T
	VisitLabel(tree *ast.Label) T
	VisitIdentifier(tree *ast.Identifier) T
	VisitSwitch(tree *ast.Switch) T
	VisitSwitchCase(tree *ast.SwitchCase) T
	VisitSynchronizedBlock(tree *ast.SynchronizedBlock) T
	VisitTryCatch(tree *ast.TryCatch) T
	VisitCatchClause(tree *ast.CatchClause) T
	VisitTypePath(tree *ast.TypePath) T
	VisitTypeElement(tree *ast.TypeElement) T
	VisitTypeParameterList(tree *ast.TypeParameterList) T
	VisitTypeParameter(tree *ast.TypeParameter) T
	VisitUnimplemented(tree *ast.Unimplemented) T
	VisitVarRef(tree *ast.VarRef) T
}
