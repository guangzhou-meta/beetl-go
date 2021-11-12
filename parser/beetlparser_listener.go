package parser // BeetlParser

import (
	"github.com/guangzhou-meta/antlr4/runtime/Go/antlr"
)

// BeetlParserListener is a complete listener for a parse tree produced by BeetlParser.
type BeetlParserListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlockSt is called when entering the blockSt production.
	EnterBlockSt(c *BlockStContext)

	// EnterTextOutputSt is called when entering the textOutputSt production.
	EnterTextOutputSt(c *TextOutputStContext)

	// EnterStaticOutputSt is called when entering the staticOutputSt production.
	EnterStaticOutputSt(c *StaticOutputStContext)

	// EnterCommentTagSt is called when entering the commentTagSt production.
	EnterCommentTagSt(c *CommentTagStContext)

	// EnterIfSt is called when entering the ifSt production.
	EnterIfSt(c *IfStContext)

	// EnterForSt is called when entering the forSt production.
	EnterForSt(c *ForStContext)

	// EnterWhileSt is called when entering the whileSt production.
	EnterWhileSt(c *WhileStContext)

	// EnterSiwchSt is called when entering the siwchSt production.
	EnterSiwchSt(c *SiwchStContext)

	// EnterSelectSt is called when entering the selectSt production.
	EnterSelectSt(c *SelectStContext)

	// EnterTrySt is called when entering the trySt production.
	EnterTrySt(c *TryStContext)

	// EnterReturnSt is called when entering the returnSt production.
	EnterReturnSt(c *ReturnStContext)

	// EnterBreakSt is called when entering the breakSt production.
	EnterBreakSt(c *BreakStContext)

	// EnterContinueSt is called when entering the continueSt production.
	EnterContinueSt(c *ContinueStContext)

	// EnterVarSt is called when entering the varSt production.
	EnterVarSt(c *VarStContext)

	// EnterDirectiveSt is called when entering the directiveSt production.
	EnterDirectiveSt(c *DirectiveStContext)

	// EnterAssignSt is called when entering the assignSt production.
	EnterAssignSt(c *AssignStContext)

	// EnterFunctionTagSt is called when entering the functionTagSt production.
	EnterFunctionTagSt(c *FunctionTagStContext)

	// EnterStatmentExpSt is called when entering the statmentExpSt production.
	EnterStatmentExpSt(c *StatmentExpStContext)

	// EnterAjaxSt is called when entering the ajaxSt production.
	EnterAjaxSt(c *AjaxStContext)

	// EnterEnd is called when entering the end production.
	EnterEnd(c *EndContext)

	// EnterCommentTypeTag is called when entering the commentTypeTag production.
	EnterCommentTypeTag(c *CommentTypeTagContext)

	// EnterCommentTypeItemTag is called when entering the commentTypeItemTag production.
	EnterCommentTypeItemTag(c *CommentTypeItemTagContext)

	// EnterClassOrInterfaceType is called when entering the classOrInterfaceType production.
	EnterClassOrInterfaceType(c *ClassOrInterfaceTypeContext)

	// EnterTypeArguments is called when entering the typeArguments production.
	EnterTypeArguments(c *TypeArgumentsContext)

	// EnterTypeArgument is called when entering the typeArgument production.
	EnterTypeArgument(c *TypeArgumentContext)

	// EnterDirectiveExp is called when entering the directiveExp production.
	EnterDirectiveExp(c *DirectiveExpContext)

	// EnterDirectiveExpIDList is called when entering the directiveExpIDList production.
	EnterDirectiveExpIDList(c *DirectiveExpIDListContext)

	// EnterG_switchStatment is called when entering the g_switchStatment production.
	EnterG_switchStatment(c *G_switchStatmentContext)

	// EnterG_caseStatment is called when entering the g_caseStatment production.
	EnterG_caseStatment(c *G_caseStatmentContext)

	// EnterG_defaultStatment is called when entering the g_defaultStatment production.
	EnterG_defaultStatment(c *G_defaultStatmentContext)

	// EnterVarDeclareList is called when entering the varDeclareList production.
	EnterVarDeclareList(c *VarDeclareListContext)

	// EnterAssignId is called when entering the assignId production.
	EnterAssignId(c *AssignIdContext)

	// EnterAssignGeneralInSt is called when entering the assignGeneralInSt production.
	EnterAssignGeneralInSt(c *AssignGeneralInStContext)

	// EnterAssignTemplateVar is called when entering the assignTemplateVar production.
	EnterAssignTemplateVar(c *AssignTemplateVarContext)

	// EnterSwitchBlock is called when entering the switchBlock production.
	EnterSwitchBlock(c *SwitchBlockContext)

	// EnterSwitchBlockStatementGroup is called when entering the switchBlockStatementGroup production.
	EnterSwitchBlockStatementGroup(c *SwitchBlockStatementGroupContext)

	// EnterSwitchLabel is called when entering the switchLabel production.
	EnterSwitchLabel(c *SwitchLabelContext)

	// EnterForControl is called when entering the forControl production.
	EnterForControl(c *ForControlContext)

	// EnterForInControl is called when entering the forInControl production.
	EnterForInControl(c *ForInControlContext)

	// EnterGeneralForControl is called when entering the generalForControl production.
	EnterGeneralForControl(c *GeneralForControlContext)

	// EnterForInit is called when entering the forInit production.
	EnterForInit(c *ForInitContext)

	// EnterForUpdate is called when entering the forUpdate production.
	EnterForUpdate(c *ForUpdateContext)

	// EnterParExpression is called when entering the parExpression production.
	EnterParExpression(c *ParExpressionContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterStatementExpression is called when entering the statementExpression production.
	EnterStatementExpression(c *StatementExpressionContext)

	// EnterTextStatment is called when entering the textStatment production.
	EnterTextStatment(c *TextStatmentContext)

	// EnterTextVar is called when entering the textVar production.
	EnterTextVar(c *TextVarContext)

	// EnterTextformat is called when entering the textformat production.
	EnterTextformat(c *TextformatContext)

	// EnterConstantsTextStatment is called when entering the constantsTextStatment production.
	EnterConstantsTextStatment(c *ConstantsTextStatmentContext)

	// EnterConstantExpression is called when entering the constantExpression production.
	EnterConstantExpression(c *ConstantExpressionContext)

	// EnterIncDecOne is called when entering the incDecOne production.
	EnterIncDecOne(c *IncDecOneContext)

	// EnterAddminExp is called when entering the addminExp production.
	EnterAddminExp(c *AddminExpContext)

	// EnterNativeCallExp is called when entering the nativeCallExp production.
	EnterNativeCallExp(c *NativeCallExpContext)

	// EnterAndExp is called when entering the andExp production.
	EnterAndExp(c *AndExpContext)

	// EnterFunctionCallExp is called when entering the functionCallExp production.
	EnterFunctionCallExp(c *FunctionCallExpContext)

	// EnterAssignGeneralInExp is called when entering the assignGeneralInExp production.
	EnterAssignGeneralInExp(c *AssignGeneralInExpContext)

	// EnterOrExp is called when entering the orExp production.
	EnterOrExp(c *OrExpContext)

	// EnterNotExp is called when entering the notExp production.
	EnterNotExp(c *NotExpContext)

	// EnterMuldivmodExp is called when entering the muldivmodExp production.
	EnterMuldivmodExp(c *MuldivmodExpContext)

	// EnterCompareExp is called when entering the compareExp production.
	EnterCompareExp(c *CompareExpContext)

	// EnterLiteralExp is called when entering the literalExp production.
	EnterLiteralExp(c *LiteralExpContext)

	// EnterJsonExp is called when entering the jsonExp production.
	EnterJsonExp(c *JsonExpContext)

	// EnterParExp is called when entering the parExp production.
	EnterParExp(c *ParExpContext)

	// EnterNegExp is called when entering the negExp production.
	EnterNegExp(c *NegExpContext)

	// EnterOneIncDec is called when entering the oneIncDec production.
	EnterOneIncDec(c *OneIncDecContext)

	// EnterTernaryExp is called when entering the ternaryExp production.
	EnterTernaryExp(c *TernaryExpContext)

	// EnterVarRefExp is called when entering the varRefExp production.
	EnterVarRefExp(c *VarRefExpContext)

	// EnterAssingSelfExp is called when entering the assingSelfExp production.
	EnterAssingSelfExp(c *AssingSelfExpContext)

	// EnterGeneralAssignExp is called when entering the generalAssignExp production.
	EnterGeneralAssignExp(c *GeneralAssignExpContext)

	// EnterVarRef is called when entering the varRef production.
	EnterVarRef(c *VarRefContext)

	// EnterGeneralAssingSelfExp is called when entering the generalAssingSelfExp production.
	EnterGeneralAssingSelfExp(c *GeneralAssingSelfExpContext)

	// EnterVarAttributeGeneral is called when entering the varAttributeGeneral production.
	EnterVarAttributeGeneral(c *VarAttributeGeneralContext)

	// EnterVarAttributeVirtual is called when entering the varAttributeVirtual production.
	EnterVarAttributeVirtual(c *VarAttributeVirtualContext)

	// EnterVarAttributeArrayOrMap is called when entering the varAttributeArrayOrMap production.
	EnterVarAttributeArrayOrMap(c *VarAttributeArrayOrMapContext)

	// EnterSafe_output is called when entering the safe_output production.
	EnterSafe_output(c *Safe_outputContext)

	// EnterSafe_allow_exp is called when entering the safe_allow_exp production.
	EnterSafe_allow_exp(c *Safe_allow_expContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterFunctionTagCall is called when entering the functionTagCall production.
	EnterFunctionTagCall(c *FunctionTagCallContext)

	// EnterFunctionNs is called when entering the functionNs production.
	EnterFunctionNs(c *FunctionNsContext)

	// EnterNativeCall is called when entering the nativeCall production.
	EnterNativeCall(c *NativeCallContext)

	// EnterNativeMethod is called when entering the nativeMethod production.
	EnterNativeMethod(c *NativeMethodContext)

	// EnterNativeArray is called when entering the nativeArray production.
	EnterNativeArray(c *NativeArrayContext)

	// EnterNativeVarRefChain is called when entering the nativeVarRefChain production.
	EnterNativeVarRefChain(c *NativeVarRefChainContext)

	// EnterJson is called when entering the json production.
	EnterJson(c *JsonContext)

	// EnterJsonKeyValue is called when entering the jsonKeyValue production.
	EnterJsonKeyValue(c *JsonKeyValueContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlockSt is called when exiting the blockSt production.
	ExitBlockSt(c *BlockStContext)

	// ExitTextOutputSt is called when exiting the textOutputSt production.
	ExitTextOutputSt(c *TextOutputStContext)

	// ExitStaticOutputSt is called when exiting the staticOutputSt production.
	ExitStaticOutputSt(c *StaticOutputStContext)

	// ExitCommentTagSt is called when exiting the commentTagSt production.
	ExitCommentTagSt(c *CommentTagStContext)

	// ExitIfSt is called when exiting the ifSt production.
	ExitIfSt(c *IfStContext)

	// ExitForSt is called when exiting the forSt production.
	ExitForSt(c *ForStContext)

	// ExitWhileSt is called when exiting the whileSt production.
	ExitWhileSt(c *WhileStContext)

	// ExitSiwchSt is called when exiting the siwchSt production.
	ExitSiwchSt(c *SiwchStContext)

	// ExitSelectSt is called when exiting the selectSt production.
	ExitSelectSt(c *SelectStContext)

	// ExitTrySt is called when exiting the trySt production.
	ExitTrySt(c *TryStContext)

	// ExitReturnSt is called when exiting the returnSt production.
	ExitReturnSt(c *ReturnStContext)

	// ExitBreakSt is called when exiting the breakSt production.
	ExitBreakSt(c *BreakStContext)

	// ExitContinueSt is called when exiting the continueSt production.
	ExitContinueSt(c *ContinueStContext)

	// ExitVarSt is called when exiting the varSt production.
	ExitVarSt(c *VarStContext)

	// ExitDirectiveSt is called when exiting the directiveSt production.
	ExitDirectiveSt(c *DirectiveStContext)

	// ExitAssignSt is called when exiting the assignSt production.
	ExitAssignSt(c *AssignStContext)

	// ExitFunctionTagSt is called when exiting the functionTagSt production.
	ExitFunctionTagSt(c *FunctionTagStContext)

	// ExitStatmentExpSt is called when exiting the statmentExpSt production.
	ExitStatmentExpSt(c *StatmentExpStContext)

	// ExitAjaxSt is called when exiting the ajaxSt production.
	ExitAjaxSt(c *AjaxStContext)

	// ExitEnd is called when exiting the end production.
	ExitEnd(c *EndContext)

	// ExitCommentTypeTag is called when exiting the commentTypeTag production.
	ExitCommentTypeTag(c *CommentTypeTagContext)

	// ExitCommentTypeItemTag is called when exiting the commentTypeItemTag production.
	ExitCommentTypeItemTag(c *CommentTypeItemTagContext)

	// ExitClassOrInterfaceType is called when exiting the classOrInterfaceType production.
	ExitClassOrInterfaceType(c *ClassOrInterfaceTypeContext)

	// ExitTypeArguments is called when exiting the typeArguments production.
	ExitTypeArguments(c *TypeArgumentsContext)

	// ExitTypeArgument is called when exiting the typeArgument production.
	ExitTypeArgument(c *TypeArgumentContext)

	// ExitDirectiveExp is called when exiting the directiveExp production.
	ExitDirectiveExp(c *DirectiveExpContext)

	// ExitDirectiveExpIDList is called when exiting the directiveExpIDList production.
	ExitDirectiveExpIDList(c *DirectiveExpIDListContext)

	// ExitG_switchStatment is called when exiting the g_switchStatment production.
	ExitG_switchStatment(c *G_switchStatmentContext)

	// ExitG_caseStatment is called when exiting the g_caseStatment production.
	ExitG_caseStatment(c *G_caseStatmentContext)

	// ExitG_defaultStatment is called when exiting the g_defaultStatment production.
	ExitG_defaultStatment(c *G_defaultStatmentContext)

	// ExitVarDeclareList is called when exiting the varDeclareList production.
	ExitVarDeclareList(c *VarDeclareListContext)

	// ExitAssignId is called when exiting the assignId production.
	ExitAssignId(c *AssignIdContext)

	// ExitAssignGeneralInSt is called when exiting the assignGeneralInSt production.
	ExitAssignGeneralInSt(c *AssignGeneralInStContext)

	// ExitAssignTemplateVar is called when exiting the assignTemplateVar production.
	ExitAssignTemplateVar(c *AssignTemplateVarContext)

	// ExitSwitchBlock is called when exiting the switchBlock production.
	ExitSwitchBlock(c *SwitchBlockContext)

	// ExitSwitchBlockStatementGroup is called when exiting the switchBlockStatementGroup production.
	ExitSwitchBlockStatementGroup(c *SwitchBlockStatementGroupContext)

	// ExitSwitchLabel is called when exiting the switchLabel production.
	ExitSwitchLabel(c *SwitchLabelContext)

	// ExitForControl is called when exiting the forControl production.
	ExitForControl(c *ForControlContext)

	// ExitForInControl is called when exiting the forInControl production.
	ExitForInControl(c *ForInControlContext)

	// ExitGeneralForControl is called when exiting the generalForControl production.
	ExitGeneralForControl(c *GeneralForControlContext)

	// ExitForInit is called when exiting the forInit production.
	ExitForInit(c *ForInitContext)

	// ExitForUpdate is called when exiting the forUpdate production.
	ExitForUpdate(c *ForUpdateContext)

	// ExitParExpression is called when exiting the parExpression production.
	ExitParExpression(c *ParExpressionContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitStatementExpression is called when exiting the statementExpression production.
	ExitStatementExpression(c *StatementExpressionContext)

	// ExitTextStatment is called when exiting the textStatment production.
	ExitTextStatment(c *TextStatmentContext)

	// ExitTextVar is called when exiting the textVar production.
	ExitTextVar(c *TextVarContext)

	// ExitTextformat is called when exiting the textformat production.
	ExitTextformat(c *TextformatContext)

	// ExitConstantsTextStatment is called when exiting the constantsTextStatment production.
	ExitConstantsTextStatment(c *ConstantsTextStatmentContext)

	// ExitConstantExpression is called when exiting the constantExpression production.
	ExitConstantExpression(c *ConstantExpressionContext)

	// ExitIncDecOne is called when exiting the incDecOne production.
	ExitIncDecOne(c *IncDecOneContext)

	// ExitAddminExp is called when exiting the addminExp production.
	ExitAddminExp(c *AddminExpContext)

	// ExitNativeCallExp is called when exiting the nativeCallExp production.
	ExitNativeCallExp(c *NativeCallExpContext)

	// ExitAndExp is called when exiting the andExp production.
	ExitAndExp(c *AndExpContext)

	// ExitFunctionCallExp is called when exiting the functionCallExp production.
	ExitFunctionCallExp(c *FunctionCallExpContext)

	// ExitAssignGeneralInExp is called when exiting the assignGeneralInExp production.
	ExitAssignGeneralInExp(c *AssignGeneralInExpContext)

	// ExitOrExp is called when exiting the orExp production.
	ExitOrExp(c *OrExpContext)

	// ExitNotExp is called when exiting the notExp production.
	ExitNotExp(c *NotExpContext)

	// ExitMuldivmodExp is called when exiting the muldivmodExp production.
	ExitMuldivmodExp(c *MuldivmodExpContext)

	// ExitCompareExp is called when exiting the compareExp production.
	ExitCompareExp(c *CompareExpContext)

	// ExitLiteralExp is called when exiting the literalExp production.
	ExitLiteralExp(c *LiteralExpContext)

	// ExitJsonExp is called when exiting the jsonExp production.
	ExitJsonExp(c *JsonExpContext)

	// ExitParExp is called when exiting the parExp production.
	ExitParExp(c *ParExpContext)

	// ExitNegExp is called when exiting the negExp production.
	ExitNegExp(c *NegExpContext)

	// ExitOneIncDec is called when exiting the oneIncDec production.
	ExitOneIncDec(c *OneIncDecContext)

	// ExitTernaryExp is called when exiting the ternaryExp production.
	ExitTernaryExp(c *TernaryExpContext)

	// ExitVarRefExp is called when exiting the varRefExp production.
	ExitVarRefExp(c *VarRefExpContext)

	// ExitAssingSelfExp is called when exiting the assingSelfExp production.
	ExitAssingSelfExp(c *AssingSelfExpContext)

	// ExitGeneralAssignExp is called when exiting the generalAssignExp production.
	ExitGeneralAssignExp(c *GeneralAssignExpContext)

	// ExitVarRef is called when exiting the varRef production.
	ExitVarRef(c *VarRefContext)

	// ExitGeneralAssingSelfExp is called when exiting the generalAssingSelfExp production.
	ExitGeneralAssingSelfExp(c *GeneralAssingSelfExpContext)

	// ExitVarAttributeGeneral is called when exiting the varAttributeGeneral production.
	ExitVarAttributeGeneral(c *VarAttributeGeneralContext)

	// ExitVarAttributeVirtual is called when exiting the varAttributeVirtual production.
	ExitVarAttributeVirtual(c *VarAttributeVirtualContext)

	// ExitVarAttributeArrayOrMap is called when exiting the varAttributeArrayOrMap production.
	ExitVarAttributeArrayOrMap(c *VarAttributeArrayOrMapContext)

	// ExitSafe_output is called when exiting the safe_output production.
	ExitSafe_output(c *Safe_outputContext)

	// ExitSafe_allow_exp is called when exiting the safe_allow_exp production.
	ExitSafe_allow_exp(c *Safe_allow_expContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitFunctionTagCall is called when exiting the functionTagCall production.
	ExitFunctionTagCall(c *FunctionTagCallContext)

	// ExitFunctionNs is called when exiting the functionNs production.
	ExitFunctionNs(c *FunctionNsContext)

	// ExitNativeCall is called when exiting the nativeCall production.
	ExitNativeCall(c *NativeCallContext)

	// ExitNativeMethod is called when exiting the nativeMethod production.
	ExitNativeMethod(c *NativeMethodContext)

	// ExitNativeArray is called when exiting the nativeArray production.
	ExitNativeArray(c *NativeArrayContext)

	// ExitNativeVarRefChain is called when exiting the nativeVarRefChain production.
	ExitNativeVarRefChain(c *NativeVarRefChainContext)

	// ExitJson is called when exiting the json production.
	ExitJson(c *JsonContext)

	// ExitJsonKeyValue is called when exiting the jsonKeyValue production.
	ExitJsonKeyValue(c *JsonKeyValueContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)
}
