// Code generated from input/grammar/JavaParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // JavaParser

import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by JavaParser.
type JavaParserVisitor[T any] interface {
	antlr.ParseTreeVisitor[T]

	// Visit a parse tree produced by JavaParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) T

	// Visit a parse tree produced by JavaParser#packageDeclaration.
	VisitPackageDeclaration(ctx *PackageDeclarationContext) T

	// Visit a parse tree produced by JavaParser#importDeclaration.
	VisitImportDeclaration(ctx *ImportDeclarationContext) T

	// Visit a parse tree produced by JavaParser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) T

	// Visit a parse tree produced by JavaParser#modifier.
	VisitModifier(ctx *ModifierContext) T

	// Visit a parse tree produced by JavaParser#classOrInterfaceModifier.
	VisitClassOrInterfaceModifier(ctx *ClassOrInterfaceModifierContext) T

	// Visit a parse tree produced by JavaParser#variableModifier.
	VisitVariableModifier(ctx *VariableModifierContext) T

	// Visit a parse tree produced by JavaParser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) T

	// Visit a parse tree produced by JavaParser#typeParameters.
	VisitTypeParameters(ctx *TypeParametersContext) T

	// Visit a parse tree produced by JavaParser#typeParameter.
	VisitTypeParameter(ctx *TypeParameterContext) T

	// Visit a parse tree produced by JavaParser#typeBound.
	VisitTypeBound(ctx *TypeBoundContext) T

	// Visit a parse tree produced by JavaParser#enumDeclaration.
	VisitEnumDeclaration(ctx *EnumDeclarationContext) T

	// Visit a parse tree produced by JavaParser#enumConstants.
	VisitEnumConstants(ctx *EnumConstantsContext) T

	// Visit a parse tree produced by JavaParser#enumConstant.
	VisitEnumConstant(ctx *EnumConstantContext) T

	// Visit a parse tree produced by JavaParser#enumBodyDeclarations.
	VisitEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) T

	// Visit a parse tree produced by JavaParser#interfaceDeclaration.
	VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) T

	// Visit a parse tree produced by JavaParser#classBody.
	VisitClassBody(ctx *ClassBodyContext) T

	// Visit a parse tree produced by JavaParser#interfaceBody.
	VisitInterfaceBody(ctx *InterfaceBodyContext) T

	// Visit a parse tree produced by JavaParser#classBodyDeclaration.
	VisitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) T

	// Visit a parse tree produced by JavaParser#memberDeclaration.
	VisitMemberDeclaration(ctx *MemberDeclarationContext) T

	// Visit a parse tree produced by JavaParser#methodDeclaration.
	VisitMethodDeclaration(ctx *MethodDeclarationContext) T

	// Visit a parse tree produced by JavaParser#methodBody.
	VisitMethodBody(ctx *MethodBodyContext) T

	// Visit a parse tree produced by JavaParser#typeTypeOrVoid.
	VisitTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) T

	// Visit a parse tree produced by JavaParser#genericMethodDeclaration.
	VisitGenericMethodDeclaration(ctx *GenericMethodDeclarationContext) T

	// Visit a parse tree produced by JavaParser#genericConstructorDeclaration.
	VisitGenericConstructorDeclaration(ctx *GenericConstructorDeclarationContext) T

	// Visit a parse tree produced by JavaParser#constructorDeclaration.
	VisitConstructorDeclaration(ctx *ConstructorDeclarationContext) T

	// Visit a parse tree produced by JavaParser#fieldDeclaration.
	VisitFieldDeclaration(ctx *FieldDeclarationContext) T

	// Visit a parse tree produced by JavaParser#interfaceBodyDeclaration.
	VisitInterfaceBodyDeclaration(ctx *InterfaceBodyDeclarationContext) T

	// Visit a parse tree produced by JavaParser#interfaceMemberDeclaration.
	VisitInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) T

	// Visit a parse tree produced by JavaParser#constDeclaration.
	VisitConstDeclaration(ctx *ConstDeclarationContext) T

	// Visit a parse tree produced by JavaParser#constantDeclarator.
	VisitConstantDeclarator(ctx *ConstantDeclaratorContext) T

	// Visit a parse tree produced by JavaParser#interfaceMethodDeclaration.
	VisitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) T

	// Visit a parse tree produced by JavaParser#interfaceMethodModifier.
	VisitInterfaceMethodModifier(ctx *InterfaceMethodModifierContext) T

	// Visit a parse tree produced by JavaParser#genericInterfaceMethodDeclaration.
	VisitGenericInterfaceMethodDeclaration(ctx *GenericInterfaceMethodDeclarationContext) T

	// Visit a parse tree produced by JavaParser#variableDeclarators.
	VisitVariableDeclarators(ctx *VariableDeclaratorsContext) T

	// Visit a parse tree produced by JavaParser#variableDeclarator.
	VisitVariableDeclarator(ctx *VariableDeclaratorContext) T

	// Visit a parse tree produced by JavaParser#variableDeclaratorId.
	VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) T

	// Visit a parse tree produced by JavaParser#variableInitializer.
	VisitVariableInitializer(ctx *VariableInitializerContext) T

	// Visit a parse tree produced by JavaParser#arrayInitializer.
	VisitArrayInitializer(ctx *ArrayInitializerContext) T

	// Visit a parse tree produced by JavaParser#classOrInterfaceType.
	VisitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) T

	// Visit a parse tree produced by JavaParser#typeArgument.
	VisitTypeArgument(ctx *TypeArgumentContext) T

	// Visit a parse tree produced by JavaParser#qualifiedNameList.
	VisitQualifiedNameList(ctx *QualifiedNameListContext) T

	// Visit a parse tree produced by JavaParser#formalParameters.
	VisitFormalParameters(ctx *FormalParametersContext) T

	// Visit a parse tree produced by JavaParser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) T

	// Visit a parse tree produced by JavaParser#formalParameter.
	VisitFormalParameter(ctx *FormalParameterContext) T

	// Visit a parse tree produced by JavaParser#lastFormalParameter.
	VisitLastFormalParameter(ctx *LastFormalParameterContext) T

	// Visit a parse tree produced by JavaParser#qualifiedName.
	VisitQualifiedName(ctx *QualifiedNameContext) T

	// Visit a parse tree produced by JavaParser#literal.
	VisitLiteral(ctx *LiteralContext) T

	// Visit a parse tree produced by JavaParser#integerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) T

	// Visit a parse tree produced by JavaParser#floatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) T

	// Visit a parse tree produced by JavaParser#altAnnotationQualifiedName.
	VisitAltAnnotationQualifiedName(ctx *AltAnnotationQualifiedNameContext) T

	// Visit a parse tree produced by JavaParser#annotation.
	VisitAnnotation(ctx *AnnotationContext) T

	// Visit a parse tree produced by JavaParser#elementValuePairs.
	VisitElementValuePairs(ctx *ElementValuePairsContext) T

	// Visit a parse tree produced by JavaParser#elementValuePair.
	VisitElementValuePair(ctx *ElementValuePairContext) T

	// Visit a parse tree produced by JavaParser#elementValue.
	VisitElementValue(ctx *ElementValueContext) T

	// Visit a parse tree produced by JavaParser#elementValueArrayInitializer.
	VisitElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) T

	// Visit a parse tree produced by JavaParser#annotationTypeDeclaration.
	VisitAnnotationTypeDeclaration(ctx *AnnotationTypeDeclarationContext) T

	// Visit a parse tree produced by JavaParser#annotationTypeBody.
	VisitAnnotationTypeBody(ctx *AnnotationTypeBodyContext) T

	// Visit a parse tree produced by JavaParser#annotationTypeElementDeclaration.
	VisitAnnotationTypeElementDeclaration(ctx *AnnotationTypeElementDeclarationContext) T

	// Visit a parse tree produced by JavaParser#annotationTypeElementRest.
	VisitAnnotationTypeElementRest(ctx *AnnotationTypeElementRestContext) T

	// Visit a parse tree produced by JavaParser#annotationMethodOrConstantRest.
	VisitAnnotationMethodOrConstantRest(ctx *AnnotationMethodOrConstantRestContext) T

	// Visit a parse tree produced by JavaParser#annotationMethodRest.
	VisitAnnotationMethodRest(ctx *AnnotationMethodRestContext) T

	// Visit a parse tree produced by JavaParser#annotationConstantRest.
	VisitAnnotationConstantRest(ctx *AnnotationConstantRestContext) T

	// Visit a parse tree produced by JavaParser#defaultValue.
	VisitDefaultValue(ctx *DefaultValueContext) T

	// Visit a parse tree produced by JavaParser#block.
	VisitBlock(ctx *BlockContext) T

	// Visit a parse tree produced by JavaParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) T

	// Visit a parse tree produced by JavaParser#localVariableDeclaration.
	VisitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) T

	// Visit a parse tree produced by JavaParser#localTypeDeclaration.
	VisitLocalTypeDeclaration(ctx *LocalTypeDeclarationContext) T

	// Visit a parse tree produced by JavaParser#statement.
	VisitStatement(ctx *StatementContext) T

	// Visit a parse tree produced by JavaParser#catchClause.
	VisitCatchClause(ctx *CatchClauseContext) T

	// Visit a parse tree produced by JavaParser#catchType.
	VisitCatchType(ctx *CatchTypeContext) T

	// Visit a parse tree produced by JavaParser#finallyBlock.
	VisitFinallyBlock(ctx *FinallyBlockContext) T

	// Visit a parse tree produced by JavaParser#resourceSpecification.
	VisitResourceSpecification(ctx *ResourceSpecificationContext) T

	// Visit a parse tree produced by JavaParser#resources.
	VisitResources(ctx *ResourcesContext) T

	// Visit a parse tree produced by JavaParser#resource.
	VisitResource(ctx *ResourceContext) T

	// Visit a parse tree produced by JavaParser#switchBlockStatementGroup.
	VisitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) T

	// Visit a parse tree produced by JavaParser#switchLabel.
	VisitSwitchLabel(ctx *SwitchLabelContext) T

	// Visit a parse tree produced by JavaParser#forControl.
	VisitForControl(ctx *ForControlContext) T

	// Visit a parse tree produced by JavaParser#forInit.
	VisitForInit(ctx *ForInitContext) T

	// Visit a parse tree produced by JavaParser#enhancedForControl.
	VisitEnhancedForControl(ctx *EnhancedForControlContext) T

	// Visit a parse tree produced by JavaParser#parExpression.
	VisitParExpression(ctx *ParExpressionContext) T

	// Visit a parse tree produced by JavaParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) T

	// Visit a parse tree produced by JavaParser#methodCall.
	VisitMethodCall(ctx *MethodCallContext) T

	// Visit a parse tree produced by JavaParser#expression.
	VisitExpression(ctx *ExpressionContext) T

	// Visit a parse tree produced by JavaParser#lambdaExpression.
	VisitLambdaExpression(ctx *LambdaExpressionContext) T

	// Visit a parse tree produced by JavaParser#lambdaParameters.
	VisitLambdaParameters(ctx *LambdaParametersContext) T

	// Visit a parse tree produced by JavaParser#lambdaBody.
	VisitLambdaBody(ctx *LambdaBodyContext) T

	// Visit a parse tree produced by JavaParser#primary.
	VisitPrimary(ctx *PrimaryContext) T

	// Visit a parse tree produced by JavaParser#classType.
	VisitClassType(ctx *ClassTypeContext) T

	// Visit a parse tree produced by JavaParser#creator.
	VisitCreator(ctx *CreatorContext) T

	// Visit a parse tree produced by JavaParser#createdName.
	VisitCreatedName(ctx *CreatedNameContext) T

	// Visit a parse tree produced by JavaParser#innerCreator.
	VisitInnerCreator(ctx *InnerCreatorContext) T

	// Visit a parse tree produced by JavaParser#arrayCreatorRest.
	VisitArrayCreatorRest(ctx *ArrayCreatorRestContext) T

	// Visit a parse tree produced by JavaParser#classCreatorRest.
	VisitClassCreatorRest(ctx *ClassCreatorRestContext) T

	// Visit a parse tree produced by JavaParser#explicitGenericInvocation.
	VisitExplicitGenericInvocation(ctx *ExplicitGenericInvocationContext) T

	// Visit a parse tree produced by JavaParser#typeArgumentsOrDiamond.
	VisitTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) T

	// Visit a parse tree produced by JavaParser#nonWildcardTypeArgumentsOrDiamond.
	VisitNonWildcardTypeArgumentsOrDiamond(ctx *NonWildcardTypeArgumentsOrDiamondContext) T

	// Visit a parse tree produced by JavaParser#nonWildcardTypeArguments.
	VisitNonWildcardTypeArguments(ctx *NonWildcardTypeArgumentsContext) T

	// Visit a parse tree produced by JavaParser#typeList.
	VisitTypeList(ctx *TypeListContext) T

	// Visit a parse tree produced by JavaParser#typeType.
	VisitTypeType(ctx *TypeTypeContext) T

	// Visit a parse tree produced by JavaParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) T

	// Visit a parse tree produced by JavaParser#typeArguments.
	VisitTypeArguments(ctx *TypeArgumentsContext) T

	// Visit a parse tree produced by JavaParser#superSuffix.
	VisitSuperSuffix(ctx *SuperSuffixContext) T

	// Visit a parse tree produced by JavaParser#explicitGenericInvocationSuffix.
	VisitExplicitGenericInvocationSuffix(ctx *ExplicitGenericInvocationSuffixContext) T

	// Visit a parse tree produced by JavaParser#arguments.
	VisitArguments(ctx *ArgumentsContext) T

}