// Code generated from input/grammar/JavaParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // JavaParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseJavaParserVisitor[T any] struct {
	*antlr.BaseParseTreeVisitor[T]
}

func NewBaseJavaParserVisitor[T any](root JavaParserVisitor[T]) *BaseJavaParserVisitor[T] {
	return &BaseJavaParserVisitor[T]{
		BaseParseTreeVisitor: antlr.NewBaseParseTreeVisitor[T](root),
	}
}

func (v *BaseJavaParserVisitor[T]) Accept(tree antlr.ParseTree) T {

	switch ttree := tree.(type) {
	case *antlr.TerminalNodeImpl:
		return v.RootVisitor.VisitTerminal(ttree)
	case *antlr.ErrorNodeImpl:
		return v.RootVisitor.VisitErrorNode(ttree)

	case *CompilationUnitContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitCompilationUnit(ttree)

	case *PackageDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitPackageDeclaration(ttree)

	case *ImportDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitImportDeclaration(ttree)

	case *TypeDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitTypeDeclaration(ttree)

	case *ModifierContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitModifier(ttree)

	case *ClassOrInterfaceModifierContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitClassOrInterfaceModifier(ttree)

	case *VariableModifierContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitVariableModifier(ttree)

	case *ClassDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitClassDeclaration(ttree)

	case *TypeParametersContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitTypeParameters(ttree)

	case *TypeParameterContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitTypeParameter(ttree)

	case *TypeBoundContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitTypeBound(ttree)

	case *EnumDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitEnumDeclaration(ttree)

	case *EnumConstantsContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitEnumConstants(ttree)

	case *EnumConstantContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitEnumConstant(ttree)

	case *EnumBodyDeclarationsContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitEnumBodyDeclarations(ttree)

	case *InterfaceDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitInterfaceDeclaration(ttree)

	case *ClassBodyContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitClassBody(ttree)

	case *InterfaceBodyContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitInterfaceBody(ttree)

	case *ClassBodyDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitClassBodyDeclaration(ttree)

	case *MemberDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitMemberDeclaration(ttree)

	case *MethodDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitMethodDeclaration(ttree)

	case *MethodBodyContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitMethodBody(ttree)

	case *TypeTypeOrVoidContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitTypeTypeOrVoid(ttree)

	case *GenericMethodDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitGenericMethodDeclaration(ttree)

	case *GenericConstructorDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitGenericConstructorDeclaration(ttree)

	case *ConstructorDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitConstructorDeclaration(ttree)

	case *FieldDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitFieldDeclaration(ttree)

	case *InterfaceBodyDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitInterfaceBodyDeclaration(ttree)

	case *InterfaceMemberDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitInterfaceMemberDeclaration(ttree)

	case *ConstDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitConstDeclaration(ttree)

	case *ConstantDeclaratorContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitConstantDeclarator(ttree)

	case *InterfaceMethodDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitInterfaceMethodDeclaration(ttree)

	case *InterfaceMethodModifierContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitInterfaceMethodModifier(ttree)

	case *GenericInterfaceMethodDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitGenericInterfaceMethodDeclaration(ttree)

	case *VariableDeclaratorsContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitVariableDeclarators(ttree)

	case *VariableDeclaratorContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitVariableDeclarator(ttree)

	case *VariableDeclaratorIdContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitVariableDeclaratorId(ttree)

	case *VariableInitializerContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitVariableInitializer(ttree)

	case *ArrayInitializerContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitArrayInitializer(ttree)

	case *ClassOrInterfaceTypeContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitClassOrInterfaceType(ttree)

	case *TypeArgumentContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitTypeArgument(ttree)

	case *QualifiedNameListContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitQualifiedNameList(ttree)

	case *FormalParametersContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitFormalParameters(ttree)

	case *FormalParameterListContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitFormalParameterList(ttree)

	case *FormalParameterContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitFormalParameter(ttree)

	case *LastFormalParameterContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitLastFormalParameter(ttree)

	case *QualifiedNameContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitQualifiedName(ttree)

	case *LiteralContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitLiteral(ttree)

	case *IntegerLiteralContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitIntegerLiteral(ttree)

	case *FloatLiteralContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitFloatLiteral(ttree)

	case *AltAnnotationQualifiedNameContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitAltAnnotationQualifiedName(ttree)

	case *AnnotationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitAnnotation(ttree)

	case *ElementValuePairsContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitElementValuePairs(ttree)

	case *ElementValuePairContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitElementValuePair(ttree)

	case *ElementValueContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitElementValue(ttree)

	case *ElementValueArrayInitializerContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitElementValueArrayInitializer(ttree)

	case *AnnotationTypeDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitAnnotationTypeDeclaration(ttree)

	case *AnnotationTypeBodyContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitAnnotationTypeBody(ttree)

	case *AnnotationTypeElementDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitAnnotationTypeElementDeclaration(ttree)

	case *AnnotationTypeElementRestContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitAnnotationTypeElementRest(ttree)

	case *AnnotationMethodOrConstantRestContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitAnnotationMethodOrConstantRest(ttree)

	case *AnnotationMethodRestContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitAnnotationMethodRest(ttree)

	case *AnnotationConstantRestContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitAnnotationConstantRest(ttree)

	case *DefaultValueContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitDefaultValue(ttree)

	case *BlockContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitBlock(ttree)

	case *BlockStatementContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitBlockStatement(ttree)

	case *LocalVariableDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitLocalVariableDeclaration(ttree)

	case *LocalTypeDeclarationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitLocalTypeDeclaration(ttree)

	case *StatementContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitStatement(ttree)

	case *CatchClauseContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitCatchClause(ttree)

	case *CatchTypeContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitCatchType(ttree)

	case *FinallyBlockContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitFinallyBlock(ttree)

	case *ResourceSpecificationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitResourceSpecification(ttree)

	case *ResourcesContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitResources(ttree)

	case *ResourceContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitResource(ttree)

	case *SwitchBlockStatementGroupContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitSwitchBlockStatementGroup(ttree)

	case *SwitchLabelContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitSwitchLabel(ttree)

	case *ForControlContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitForControl(ttree)

	case *ForInitContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitForInit(ttree)

	case *EnhancedForControlContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitEnhancedForControl(ttree)

	case *ParExpressionContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitParExpression(ttree)

	case *ExpressionListContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitExpressionList(ttree)

	case *MethodCallContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitMethodCall(ttree)

	case *ExpressionContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitExpression(ttree)

	case *LambdaExpressionContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitLambdaExpression(ttree)

	case *LambdaParametersContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitLambdaParameters(ttree)

	case *LambdaBodyContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitLambdaBody(ttree)

	case *PrimaryContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitPrimary(ttree)

	case *ClassTypeContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitClassType(ttree)

	case *CreatorContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitCreator(ttree)

	case *CreatedNameContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitCreatedName(ttree)

	case *InnerCreatorContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitInnerCreator(ttree)

	case *ArrayCreatorRestContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitArrayCreatorRest(ttree)

	case *ClassCreatorRestContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitClassCreatorRest(ttree)

	case *ExplicitGenericInvocationContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitExplicitGenericInvocation(ttree)

	case *TypeArgumentsOrDiamondContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitTypeArgumentsOrDiamond(ttree)

	case *NonWildcardTypeArgumentsOrDiamondContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitNonWildcardTypeArgumentsOrDiamond(ttree)

	case *NonWildcardTypeArgumentsContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitNonWildcardTypeArguments(ttree)

	case *TypeListContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitTypeList(ttree)

	case *TypeTypeContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitTypeType(ttree)

	case *PrimitiveTypeContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitPrimitiveType(ttree)

	case *TypeArgumentsContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitTypeArguments(ttree)

	case *SuperSuffixContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitSuperSuffix(ttree)

	case *ExplicitGenericInvocationSuffixContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitExplicitGenericInvocationSuffix(ttree)

	case *ArgumentsContext:
		return v.BaseParseTreeVisitor.RootVisitor.(JavaParserVisitor[T]).VisitArguments(ttree)


	}
	var zero T
	return zero
}

func (v *BaseJavaParserVisitor[T]) VisitCompilationUnit(ctx *CompilationUnitContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitPackageDeclaration(ctx *PackageDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitImportDeclaration(ctx *ImportDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitTypeDeclaration(ctx *TypeDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitModifier(ctx *ModifierContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitClassOrInterfaceModifier(ctx *ClassOrInterfaceModifierContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitVariableModifier(ctx *VariableModifierContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitClassDeclaration(ctx *ClassDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitTypeParameters(ctx *TypeParametersContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitTypeParameter(ctx *TypeParameterContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitTypeBound(ctx *TypeBoundContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitEnumDeclaration(ctx *EnumDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitEnumConstants(ctx *EnumConstantsContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitEnumConstant(ctx *EnumConstantContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitClassBody(ctx *ClassBodyContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitInterfaceBody(ctx *InterfaceBodyContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitMemberDeclaration(ctx *MemberDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitMethodDeclaration(ctx *MethodDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitMethodBody(ctx *MethodBodyContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitGenericMethodDeclaration(ctx *GenericMethodDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitGenericConstructorDeclaration(ctx *GenericConstructorDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitConstructorDeclaration(ctx *ConstructorDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitFieldDeclaration(ctx *FieldDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitInterfaceBodyDeclaration(ctx *InterfaceBodyDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitConstDeclaration(ctx *ConstDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitConstantDeclarator(ctx *ConstantDeclaratorContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitInterfaceMethodModifier(ctx *InterfaceMethodModifierContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitGenericInterfaceMethodDeclaration(ctx *GenericInterfaceMethodDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitVariableDeclarators(ctx *VariableDeclaratorsContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitVariableDeclarator(ctx *VariableDeclaratorContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitVariableInitializer(ctx *VariableInitializerContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitArrayInitializer(ctx *ArrayInitializerContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitTypeArgument(ctx *TypeArgumentContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitQualifiedNameList(ctx *QualifiedNameListContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitFormalParameters(ctx *FormalParametersContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitFormalParameterList(ctx *FormalParameterListContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitFormalParameter(ctx *FormalParameterContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitLastFormalParameter(ctx *LastFormalParameterContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitQualifiedName(ctx *QualifiedNameContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitLiteral(ctx *LiteralContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitIntegerLiteral(ctx *IntegerLiteralContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitFloatLiteral(ctx *FloatLiteralContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitAltAnnotationQualifiedName(ctx *AltAnnotationQualifiedNameContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitAnnotation(ctx *AnnotationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitElementValuePairs(ctx *ElementValuePairsContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitElementValuePair(ctx *ElementValuePairContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitElementValue(ctx *ElementValueContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitAnnotationTypeDeclaration(ctx *AnnotationTypeDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitAnnotationTypeBody(ctx *AnnotationTypeBodyContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitAnnotationTypeElementDeclaration(ctx *AnnotationTypeElementDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitAnnotationTypeElementRest(ctx *AnnotationTypeElementRestContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitAnnotationMethodOrConstantRest(ctx *AnnotationMethodOrConstantRestContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitAnnotationMethodRest(ctx *AnnotationMethodRestContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitAnnotationConstantRest(ctx *AnnotationConstantRestContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitDefaultValue(ctx *DefaultValueContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitBlock(ctx *BlockContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitBlockStatement(ctx *BlockStatementContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitLocalTypeDeclaration(ctx *LocalTypeDeclarationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitStatement(ctx *StatementContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitCatchClause(ctx *CatchClauseContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitCatchType(ctx *CatchTypeContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitFinallyBlock(ctx *FinallyBlockContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitResourceSpecification(ctx *ResourceSpecificationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitResources(ctx *ResourcesContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitResource(ctx *ResourceContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitSwitchLabel(ctx *SwitchLabelContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitForControl(ctx *ForControlContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitForInit(ctx *ForInitContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitEnhancedForControl(ctx *EnhancedForControlContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitParExpression(ctx *ParExpressionContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitExpressionList(ctx *ExpressionListContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitMethodCall(ctx *MethodCallContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitExpression(ctx *ExpressionContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitLambdaExpression(ctx *LambdaExpressionContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitLambdaParameters(ctx *LambdaParametersContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitLambdaBody(ctx *LambdaBodyContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitPrimary(ctx *PrimaryContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitClassType(ctx *ClassTypeContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitCreator(ctx *CreatorContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitCreatedName(ctx *CreatedNameContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitInnerCreator(ctx *InnerCreatorContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitArrayCreatorRest(ctx *ArrayCreatorRestContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitClassCreatorRest(ctx *ClassCreatorRestContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitExplicitGenericInvocation(ctx *ExplicitGenericInvocationContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitNonWildcardTypeArgumentsOrDiamond(ctx *NonWildcardTypeArgumentsOrDiamondContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitNonWildcardTypeArguments(ctx *NonWildcardTypeArgumentsContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitTypeList(ctx *TypeListContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitTypeType(ctx *TypeTypeContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitPrimitiveType(ctx *PrimitiveTypeContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitTypeArguments(ctx *TypeArgumentsContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitSuperSuffix(ctx *SuperSuffixContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitExplicitGenericInvocationSuffix(ctx *ExplicitGenericInvocationSuffixContext) T {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaParserVisitor[T]) VisitArguments(ctx *ArgumentsContext) T {
	return v.VisitChildren(ctx)
}
