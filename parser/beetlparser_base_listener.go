package parser // BeetlParser

import (
	"github.com/guangzhou-meta/antlr4/runtime/Go/antlr"
)

// BaseBeetlParserListener is a complete listener for a parse tree produced by BeetlParser.
type BaseBeetlParserListener struct{}

var _ BeetlParserListener = &BaseBeetlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBeetlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBeetlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBeetlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBeetlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseBeetlParserListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseBeetlParserListener) ExitProg(ctx *ProgContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseBeetlParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseBeetlParserListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockSt is called when production blockSt is entered.
func (s *BaseBeetlParserListener) EnterBlockSt(ctx *BlockStContext) {}

// ExitBlockSt is called when production blockSt is exited.
func (s *BaseBeetlParserListener) ExitBlockSt(ctx *BlockStContext) {}

// EnterTextOutputSt is called when production textOutputSt is entered.
func (s *BaseBeetlParserListener) EnterTextOutputSt(ctx *TextOutputStContext) {}

// ExitTextOutputSt is called when production textOutputSt is exited.
func (s *BaseBeetlParserListener) ExitTextOutputSt(ctx *TextOutputStContext) {}

// EnterStaticOutputSt is called when production staticOutputSt is entered.
func (s *BaseBeetlParserListener) EnterStaticOutputSt(ctx *StaticOutputStContext) {}

// ExitStaticOutputSt is called when production staticOutputSt is exited.
func (s *BaseBeetlParserListener) ExitStaticOutputSt(ctx *StaticOutputStContext) {}

// EnterCommentTagSt is called when production commentTagSt is entered.
func (s *BaseBeetlParserListener) EnterCommentTagSt(ctx *CommentTagStContext) {}

// ExitCommentTagSt is called when production commentTagSt is exited.
func (s *BaseBeetlParserListener) ExitCommentTagSt(ctx *CommentTagStContext) {}

// EnterIfSt is called when production ifSt is entered.
func (s *BaseBeetlParserListener) EnterIfSt(ctx *IfStContext) {}

// ExitIfSt is called when production ifSt is exited.
func (s *BaseBeetlParserListener) ExitIfSt(ctx *IfStContext) {}

// EnterForSt is called when production forSt is entered.
func (s *BaseBeetlParserListener) EnterForSt(ctx *ForStContext) {}

// ExitForSt is called when production forSt is exited.
func (s *BaseBeetlParserListener) ExitForSt(ctx *ForStContext) {}

// EnterWhileSt is called when production whileSt is entered.
func (s *BaseBeetlParserListener) EnterWhileSt(ctx *WhileStContext) {}

// ExitWhileSt is called when production whileSt is exited.
func (s *BaseBeetlParserListener) ExitWhileSt(ctx *WhileStContext) {}

// EnterSiwchSt is called when production siwchSt is entered.
func (s *BaseBeetlParserListener) EnterSiwchSt(ctx *SiwchStContext) {}

// ExitSiwchSt is called when production siwchSt is exited.
func (s *BaseBeetlParserListener) ExitSiwchSt(ctx *SiwchStContext) {}

// EnterSelectSt is called when production selectSt is entered.
func (s *BaseBeetlParserListener) EnterSelectSt(ctx *SelectStContext) {}

// ExitSelectSt is called when production selectSt is exited.
func (s *BaseBeetlParserListener) ExitSelectSt(ctx *SelectStContext) {}

// EnterTrySt is called when production trySt is entered.
func (s *BaseBeetlParserListener) EnterTrySt(ctx *TryStContext) {}

// ExitTrySt is called when production trySt is exited.
func (s *BaseBeetlParserListener) ExitTrySt(ctx *TryStContext) {}

// EnterReturnSt is called when production returnSt is entered.
func (s *BaseBeetlParserListener) EnterReturnSt(ctx *ReturnStContext) {}

// ExitReturnSt is called when production returnSt is exited.
func (s *BaseBeetlParserListener) ExitReturnSt(ctx *ReturnStContext) {}

// EnterBreakSt is called when production breakSt is entered.
func (s *BaseBeetlParserListener) EnterBreakSt(ctx *BreakStContext) {}

// ExitBreakSt is called when production breakSt is exited.
func (s *BaseBeetlParserListener) ExitBreakSt(ctx *BreakStContext) {}

// EnterContinueSt is called when production continueSt is entered.
func (s *BaseBeetlParserListener) EnterContinueSt(ctx *ContinueStContext) {}

// ExitContinueSt is called when production continueSt is exited.
func (s *BaseBeetlParserListener) ExitContinueSt(ctx *ContinueStContext) {}

// EnterVarSt is called when production varSt is entered.
func (s *BaseBeetlParserListener) EnterVarSt(ctx *VarStContext) {}

// ExitVarSt is called when production varSt is exited.
func (s *BaseBeetlParserListener) ExitVarSt(ctx *VarStContext) {}

// EnterDirectiveSt is called when production directiveSt is entered.
func (s *BaseBeetlParserListener) EnterDirectiveSt(ctx *DirectiveStContext) {}

// ExitDirectiveSt is called when production directiveSt is exited.
func (s *BaseBeetlParserListener) ExitDirectiveSt(ctx *DirectiveStContext) {}

// EnterAssignSt is called when production assignSt is entered.
func (s *BaseBeetlParserListener) EnterAssignSt(ctx *AssignStContext) {}

// ExitAssignSt is called when production assignSt is exited.
func (s *BaseBeetlParserListener) ExitAssignSt(ctx *AssignStContext) {}

// EnterFunctionTagSt is called when production functionTagSt is entered.
func (s *BaseBeetlParserListener) EnterFunctionTagSt(ctx *FunctionTagStContext) {}

// ExitFunctionTagSt is called when production functionTagSt is exited.
func (s *BaseBeetlParserListener) ExitFunctionTagSt(ctx *FunctionTagStContext) {}

// EnterStatmentExpSt is called when production statmentExpSt is entered.
func (s *BaseBeetlParserListener) EnterStatmentExpSt(ctx *StatmentExpStContext) {}

// ExitStatmentExpSt is called when production statmentExpSt is exited.
func (s *BaseBeetlParserListener) ExitStatmentExpSt(ctx *StatmentExpStContext) {}

// EnterAjaxSt is called when production ajaxSt is entered.
func (s *BaseBeetlParserListener) EnterAjaxSt(ctx *AjaxStContext) {}

// ExitAjaxSt is called when production ajaxSt is exited.
func (s *BaseBeetlParserListener) ExitAjaxSt(ctx *AjaxStContext) {}

// EnterEnd is called when production end is entered.
func (s *BaseBeetlParserListener) EnterEnd(ctx *EndContext) {}

// ExitEnd is called when production end is exited.
func (s *BaseBeetlParserListener) ExitEnd(ctx *EndContext) {}

// EnterCommentTypeTag is called when production commentTypeTag is entered.
func (s *BaseBeetlParserListener) EnterCommentTypeTag(ctx *CommentTypeTagContext) {}

// ExitCommentTypeTag is called when production commentTypeTag is exited.
func (s *BaseBeetlParserListener) ExitCommentTypeTag(ctx *CommentTypeTagContext) {}

// EnterCommentTypeItemTag is called when production commentTypeItemTag is entered.
func (s *BaseBeetlParserListener) EnterCommentTypeItemTag(ctx *CommentTypeItemTagContext) {}

// ExitCommentTypeItemTag is called when production commentTypeItemTag is exited.
func (s *BaseBeetlParserListener) ExitCommentTypeItemTag(ctx *CommentTypeItemTagContext) {}

// EnterClassOrInterfaceType is called when production classOrInterfaceType is entered.
func (s *BaseBeetlParserListener) EnterClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// ExitClassOrInterfaceType is called when production classOrInterfaceType is exited.
func (s *BaseBeetlParserListener) ExitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// EnterTypeArguments is called when production typeArguments is entered.
func (s *BaseBeetlParserListener) EnterTypeArguments(ctx *TypeArgumentsContext) {}

// ExitTypeArguments is called when production typeArguments is exited.
func (s *BaseBeetlParserListener) ExitTypeArguments(ctx *TypeArgumentsContext) {}

// EnterTypeArgument is called when production typeArgument is entered.
func (s *BaseBeetlParserListener) EnterTypeArgument(ctx *TypeArgumentContext) {}

// ExitTypeArgument is called when production typeArgument is exited.
func (s *BaseBeetlParserListener) ExitTypeArgument(ctx *TypeArgumentContext) {}

// EnterDirectiveExp is called when production directiveExp is entered.
func (s *BaseBeetlParserListener) EnterDirectiveExp(ctx *DirectiveExpContext) {}

// ExitDirectiveExp is called when production directiveExp is exited.
func (s *BaseBeetlParserListener) ExitDirectiveExp(ctx *DirectiveExpContext) {}

// EnterDirectiveExpIDList is called when production directiveExpIDList is entered.
func (s *BaseBeetlParserListener) EnterDirectiveExpIDList(ctx *DirectiveExpIDListContext) {}

// ExitDirectiveExpIDList is called when production directiveExpIDList is exited.
func (s *BaseBeetlParserListener) ExitDirectiveExpIDList(ctx *DirectiveExpIDListContext) {}

// EnterG_switchStatment is called when production g_switchStatment is entered.
func (s *BaseBeetlParserListener) EnterG_switchStatment(ctx *G_switchStatmentContext) {}

// ExitG_switchStatment is called when production g_switchStatment is exited.
func (s *BaseBeetlParserListener) ExitG_switchStatment(ctx *G_switchStatmentContext) {}

// EnterG_caseStatment is called when production g_caseStatment is entered.
func (s *BaseBeetlParserListener) EnterG_caseStatment(ctx *G_caseStatmentContext) {}

// ExitG_caseStatment is called when production g_caseStatment is exited.
func (s *BaseBeetlParserListener) ExitG_caseStatment(ctx *G_caseStatmentContext) {}

// EnterG_defaultStatment is called when production g_defaultStatment is entered.
func (s *BaseBeetlParserListener) EnterG_defaultStatment(ctx *G_defaultStatmentContext) {}

// ExitG_defaultStatment is called when production g_defaultStatment is exited.
func (s *BaseBeetlParserListener) ExitG_defaultStatment(ctx *G_defaultStatmentContext) {}

// EnterVarDeclareList is called when production varDeclareList is entered.
func (s *BaseBeetlParserListener) EnterVarDeclareList(ctx *VarDeclareListContext) {}

// ExitVarDeclareList is called when production varDeclareList is exited.
func (s *BaseBeetlParserListener) ExitVarDeclareList(ctx *VarDeclareListContext) {}

// EnterAssignId is called when production assignId is entered.
func (s *BaseBeetlParserListener) EnterAssignId(ctx *AssignIdContext) {}

// ExitAssignId is called when production assignId is exited.
func (s *BaseBeetlParserListener) ExitAssignId(ctx *AssignIdContext) {}

// EnterAssignGeneralInSt is called when production assignGeneralInSt is entered.
func (s *BaseBeetlParserListener) EnterAssignGeneralInSt(ctx *AssignGeneralInStContext) {}

// ExitAssignGeneralInSt is called when production assignGeneralInSt is exited.
func (s *BaseBeetlParserListener) ExitAssignGeneralInSt(ctx *AssignGeneralInStContext) {}

// EnterAssignTemplateVar is called when production assignTemplateVar is entered.
func (s *BaseBeetlParserListener) EnterAssignTemplateVar(ctx *AssignTemplateVarContext) {}

// ExitAssignTemplateVar is called when production assignTemplateVar is exited.
func (s *BaseBeetlParserListener) ExitAssignTemplateVar(ctx *AssignTemplateVarContext) {}

// EnterSwitchBlock is called when production switchBlock is entered.
func (s *BaseBeetlParserListener) EnterSwitchBlock(ctx *SwitchBlockContext) {}

// ExitSwitchBlock is called when production switchBlock is exited.
func (s *BaseBeetlParserListener) ExitSwitchBlock(ctx *SwitchBlockContext) {}

// EnterSwitchBlockStatementGroup is called when production switchBlockStatementGroup is entered.
func (s *BaseBeetlParserListener) EnterSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {
}

// ExitSwitchBlockStatementGroup is called when production switchBlockStatementGroup is exited.
func (s *BaseBeetlParserListener) ExitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {
}

// EnterSwitchLabel is called when production switchLabel is entered.
func (s *BaseBeetlParserListener) EnterSwitchLabel(ctx *SwitchLabelContext) {}

// ExitSwitchLabel is called when production switchLabel is exited.
func (s *BaseBeetlParserListener) ExitSwitchLabel(ctx *SwitchLabelContext) {}

// EnterForControl is called when production forControl is entered.
func (s *BaseBeetlParserListener) EnterForControl(ctx *ForControlContext) {}

// ExitForControl is called when production forControl is exited.
func (s *BaseBeetlParserListener) ExitForControl(ctx *ForControlContext) {}

// EnterForInControl is called when production forInControl is entered.
func (s *BaseBeetlParserListener) EnterForInControl(ctx *ForInControlContext) {}

// ExitForInControl is called when production forInControl is exited.
func (s *BaseBeetlParserListener) ExitForInControl(ctx *ForInControlContext) {}

// EnterGeneralForControl is called when production generalForControl is entered.
func (s *BaseBeetlParserListener) EnterGeneralForControl(ctx *GeneralForControlContext) {}

// ExitGeneralForControl is called when production generalForControl is exited.
func (s *BaseBeetlParserListener) ExitGeneralForControl(ctx *GeneralForControlContext) {}

// EnterForInit is called when production forInit is entered.
func (s *BaseBeetlParserListener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BaseBeetlParserListener) ExitForInit(ctx *ForInitContext) {}

// EnterForUpdate is called when production forUpdate is entered.
func (s *BaseBeetlParserListener) EnterForUpdate(ctx *ForUpdateContext) {}

// ExitForUpdate is called when production forUpdate is exited.
func (s *BaseBeetlParserListener) ExitForUpdate(ctx *ForUpdateContext) {}

// EnterParExpression is called when production parExpression is entered.
func (s *BaseBeetlParserListener) EnterParExpression(ctx *ParExpressionContext) {}

// ExitParExpression is called when production parExpression is exited.
func (s *BaseBeetlParserListener) ExitParExpression(ctx *ParExpressionContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseBeetlParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseBeetlParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterStatementExpression is called when production statementExpression is entered.
func (s *BaseBeetlParserListener) EnterStatementExpression(ctx *StatementExpressionContext) {}

// ExitStatementExpression is called when production statementExpression is exited.
func (s *BaseBeetlParserListener) ExitStatementExpression(ctx *StatementExpressionContext) {}

// EnterTextStatment is called when production textStatment is entered.
func (s *BaseBeetlParserListener) EnterTextStatment(ctx *TextStatmentContext) {}

// ExitTextStatment is called when production textStatment is exited.
func (s *BaseBeetlParserListener) ExitTextStatment(ctx *TextStatmentContext) {}

// EnterTextVar is called when production textVar is entered.
func (s *BaseBeetlParserListener) EnterTextVar(ctx *TextVarContext) {}

// ExitTextVar is called when production textVar is exited.
func (s *BaseBeetlParserListener) ExitTextVar(ctx *TextVarContext) {}

// EnterTextformat is called when production textformat is entered.
func (s *BaseBeetlParserListener) EnterTextformat(ctx *TextformatContext) {}

// ExitTextformat is called when production textformat is exited.
func (s *BaseBeetlParserListener) ExitTextformat(ctx *TextformatContext) {}

// EnterConstantsTextStatment is called when production constantsTextStatment is entered.
func (s *BaseBeetlParserListener) EnterConstantsTextStatment(ctx *ConstantsTextStatmentContext) {}

// ExitConstantsTextStatment is called when production constantsTextStatment is exited.
func (s *BaseBeetlParserListener) ExitConstantsTextStatment(ctx *ConstantsTextStatmentContext) {}

// EnterConstantExpression is called when production constantExpression is entered.
func (s *BaseBeetlParserListener) EnterConstantExpression(ctx *ConstantExpressionContext) {}

// ExitConstantExpression is called when production constantExpression is exited.
func (s *BaseBeetlParserListener) ExitConstantExpression(ctx *ConstantExpressionContext) {}

// EnterIncDecOne is called when production incDecOne is entered.
func (s *BaseBeetlParserListener) EnterIncDecOne(ctx *IncDecOneContext) {}

// ExitIncDecOne is called when production incDecOne is exited.
func (s *BaseBeetlParserListener) ExitIncDecOne(ctx *IncDecOneContext) {}

// EnterAddminExp is called when production addminExp is entered.
func (s *BaseBeetlParserListener) EnterAddminExp(ctx *AddminExpContext) {}

// ExitAddminExp is called when production addminExp is exited.
func (s *BaseBeetlParserListener) ExitAddminExp(ctx *AddminExpContext) {}

// EnterNativeCallExp is called when production nativeCallExp is entered.
func (s *BaseBeetlParserListener) EnterNativeCallExp(ctx *NativeCallExpContext) {}

// ExitNativeCallExp is called when production nativeCallExp is exited.
func (s *BaseBeetlParserListener) ExitNativeCallExp(ctx *NativeCallExpContext) {}

// EnterAndExp is called when production andExp is entered.
func (s *BaseBeetlParserListener) EnterAndExp(ctx *AndExpContext) {}

// ExitAndExp is called when production andExp is exited.
func (s *BaseBeetlParserListener) ExitAndExp(ctx *AndExpContext) {}

// EnterFunctionCallExp is called when production functionCallExp is entered.
func (s *BaseBeetlParserListener) EnterFunctionCallExp(ctx *FunctionCallExpContext) {}

// ExitFunctionCallExp is called when production functionCallExp is exited.
func (s *BaseBeetlParserListener) ExitFunctionCallExp(ctx *FunctionCallExpContext) {}

// EnterAssignGeneralInExp is called when production assignGeneralInExp is entered.
func (s *BaseBeetlParserListener) EnterAssignGeneralInExp(ctx *AssignGeneralInExpContext) {}

// ExitAssignGeneralInExp is called when production assignGeneralInExp is exited.
func (s *BaseBeetlParserListener) ExitAssignGeneralInExp(ctx *AssignGeneralInExpContext) {}

// EnterOrExp is called when production orExp is entered.
func (s *BaseBeetlParserListener) EnterOrExp(ctx *OrExpContext) {}

// ExitOrExp is called when production orExp is exited.
func (s *BaseBeetlParserListener) ExitOrExp(ctx *OrExpContext) {}

// EnterNotExp is called when production notExp is entered.
func (s *BaseBeetlParserListener) EnterNotExp(ctx *NotExpContext) {}

// ExitNotExp is called when production notExp is exited.
func (s *BaseBeetlParserListener) ExitNotExp(ctx *NotExpContext) {}

// EnterMuldivmodExp is called when production muldivmodExp is entered.
func (s *BaseBeetlParserListener) EnterMuldivmodExp(ctx *MuldivmodExpContext) {}

// ExitMuldivmodExp is called when production muldivmodExp is exited.
func (s *BaseBeetlParserListener) ExitMuldivmodExp(ctx *MuldivmodExpContext) {}

// EnterCompareExp is called when production compareExp is entered.
func (s *BaseBeetlParserListener) EnterCompareExp(ctx *CompareExpContext) {}

// ExitCompareExp is called when production compareExp is exited.
func (s *BaseBeetlParserListener) ExitCompareExp(ctx *CompareExpContext) {}

// EnterLiteralExp is called when production literalExp is entered.
func (s *BaseBeetlParserListener) EnterLiteralExp(ctx *LiteralExpContext) {}

// ExitLiteralExp is called when production literalExp is exited.
func (s *BaseBeetlParserListener) ExitLiteralExp(ctx *LiteralExpContext) {}

// EnterJsonExp is called when production jsonExp is entered.
func (s *BaseBeetlParserListener) EnterJsonExp(ctx *JsonExpContext) {}

// ExitJsonExp is called when production jsonExp is exited.
func (s *BaseBeetlParserListener) ExitJsonExp(ctx *JsonExpContext) {}

// EnterParExp is called when production parExp is entered.
func (s *BaseBeetlParserListener) EnterParExp(ctx *ParExpContext) {}

// ExitParExp is called when production parExp is exited.
func (s *BaseBeetlParserListener) ExitParExp(ctx *ParExpContext) {}

// EnterNegExp is called when production negExp is entered.
func (s *BaseBeetlParserListener) EnterNegExp(ctx *NegExpContext) {}

// ExitNegExp is called when production negExp is exited.
func (s *BaseBeetlParserListener) ExitNegExp(ctx *NegExpContext) {}

// EnterOneIncDec is called when production oneIncDec is entered.
func (s *BaseBeetlParserListener) EnterOneIncDec(ctx *OneIncDecContext) {}

// ExitOneIncDec is called when production oneIncDec is exited.
func (s *BaseBeetlParserListener) ExitOneIncDec(ctx *OneIncDecContext) {}

// EnterTernaryExp is called when production ternaryExp is entered.
func (s *BaseBeetlParserListener) EnterTernaryExp(ctx *TernaryExpContext) {}

// ExitTernaryExp is called when production ternaryExp is exited.
func (s *BaseBeetlParserListener) ExitTernaryExp(ctx *TernaryExpContext) {}

// EnterVarRefExp is called when production varRefExp is entered.
func (s *BaseBeetlParserListener) EnterVarRefExp(ctx *VarRefExpContext) {}

// ExitVarRefExp is called when production varRefExp is exited.
func (s *BaseBeetlParserListener) ExitVarRefExp(ctx *VarRefExpContext) {}

// EnterAssingSelfExp is called when production assingSelfExp is entered.
func (s *BaseBeetlParserListener) EnterAssingSelfExp(ctx *AssingSelfExpContext) {}

// ExitAssingSelfExp is called when production assingSelfExp is exited.
func (s *BaseBeetlParserListener) ExitAssingSelfExp(ctx *AssingSelfExpContext) {}

// EnterGeneralAssignExp is called when production generalAssignExp is entered.
func (s *BaseBeetlParserListener) EnterGeneralAssignExp(ctx *GeneralAssignExpContext) {}

// ExitGeneralAssignExp is called when production generalAssignExp is exited.
func (s *BaseBeetlParserListener) ExitGeneralAssignExp(ctx *GeneralAssignExpContext) {}

// EnterVarRef is called when production varRef is entered.
func (s *BaseBeetlParserListener) EnterVarRef(ctx *VarRefContext) {}

// ExitVarRef is called when production varRef is exited.
func (s *BaseBeetlParserListener) ExitVarRef(ctx *VarRefContext) {}

// EnterGeneralAssingSelfExp is called when production generalAssingSelfExp is entered.
func (s *BaseBeetlParserListener) EnterGeneralAssingSelfExp(ctx *GeneralAssingSelfExpContext) {}

// ExitGeneralAssingSelfExp is called when production generalAssingSelfExp is exited.
func (s *BaseBeetlParserListener) ExitGeneralAssingSelfExp(ctx *GeneralAssingSelfExpContext) {}

// EnterVarAttributeGeneral is called when production varAttributeGeneral is entered.
func (s *BaseBeetlParserListener) EnterVarAttributeGeneral(ctx *VarAttributeGeneralContext) {}

// ExitVarAttributeGeneral is called when production varAttributeGeneral is exited.
func (s *BaseBeetlParserListener) ExitVarAttributeGeneral(ctx *VarAttributeGeneralContext) {}

// EnterVarAttributeVirtual is called when production varAttributeVirtual is entered.
func (s *BaseBeetlParserListener) EnterVarAttributeVirtual(ctx *VarAttributeVirtualContext) {}

// ExitVarAttributeVirtual is called when production varAttributeVirtual is exited.
func (s *BaseBeetlParserListener) ExitVarAttributeVirtual(ctx *VarAttributeVirtualContext) {}

// EnterVarAttributeArrayOrMap is called when production varAttributeArrayOrMap is entered.
func (s *BaseBeetlParserListener) EnterVarAttributeArrayOrMap(ctx *VarAttributeArrayOrMapContext) {}

// ExitVarAttributeArrayOrMap is called when production varAttributeArrayOrMap is exited.
func (s *BaseBeetlParserListener) ExitVarAttributeArrayOrMap(ctx *VarAttributeArrayOrMapContext) {}

// EnterSafe_output is called when production safe_output is entered.
func (s *BaseBeetlParserListener) EnterSafe_output(ctx *Safe_outputContext) {}

// ExitSafe_output is called when production safe_output is exited.
func (s *BaseBeetlParserListener) ExitSafe_output(ctx *Safe_outputContext) {}

// EnterSafe_allow_exp is called when production safe_allow_exp is entered.
func (s *BaseBeetlParserListener) EnterSafe_allow_exp(ctx *Safe_allow_expContext) {}

// ExitSafe_allow_exp is called when production safe_allow_exp is exited.
func (s *BaseBeetlParserListener) ExitSafe_allow_exp(ctx *Safe_allow_expContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseBeetlParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseBeetlParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterFunctionTagCall is called when production functionTagCall is entered.
func (s *BaseBeetlParserListener) EnterFunctionTagCall(ctx *FunctionTagCallContext) {}

// ExitFunctionTagCall is called when production functionTagCall is exited.
func (s *BaseBeetlParserListener) ExitFunctionTagCall(ctx *FunctionTagCallContext) {}

// EnterFunctionNs is called when production functionNs is entered.
func (s *BaseBeetlParserListener) EnterFunctionNs(ctx *FunctionNsContext) {}

// ExitFunctionNs is called when production functionNs is exited.
func (s *BaseBeetlParserListener) ExitFunctionNs(ctx *FunctionNsContext) {}

// EnterNativeCall is called when production nativeCall is entered.
func (s *BaseBeetlParserListener) EnterNativeCall(ctx *NativeCallContext) {}

// ExitNativeCall is called when production nativeCall is exited.
func (s *BaseBeetlParserListener) ExitNativeCall(ctx *NativeCallContext) {}

// EnterNativeMethod is called when production nativeMethod is entered.
func (s *BaseBeetlParserListener) EnterNativeMethod(ctx *NativeMethodContext) {}

// ExitNativeMethod is called when production nativeMethod is exited.
func (s *BaseBeetlParserListener) ExitNativeMethod(ctx *NativeMethodContext) {}

// EnterNativeArray is called when production nativeArray is entered.
func (s *BaseBeetlParserListener) EnterNativeArray(ctx *NativeArrayContext) {}

// ExitNativeArray is called when production nativeArray is exited.
func (s *BaseBeetlParserListener) ExitNativeArray(ctx *NativeArrayContext) {}

// EnterNativeVarRefChain is called when production nativeVarRefChain is entered.
func (s *BaseBeetlParserListener) EnterNativeVarRefChain(ctx *NativeVarRefChainContext) {}

// ExitNativeVarRefChain is called when production nativeVarRefChain is exited.
func (s *BaseBeetlParserListener) ExitNativeVarRefChain(ctx *NativeVarRefChainContext) {}

// EnterJson is called when production json is entered.
func (s *BaseBeetlParserListener) EnterJson(ctx *JsonContext) {}

// ExitJson is called when production json is exited.
func (s *BaseBeetlParserListener) ExitJson(ctx *JsonContext) {}

// EnterJsonKeyValue is called when production jsonKeyValue is entered.
func (s *BaseBeetlParserListener) EnterJsonKeyValue(ctx *JsonKeyValueContext) {}

// ExitJsonKeyValue is called when production jsonKeyValue is exited.
func (s *BaseBeetlParserListener) ExitJsonKeyValue(ctx *JsonKeyValueContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseBeetlParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseBeetlParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseBeetlParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseBeetlParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseBeetlParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseBeetlParserListener) ExitArguments(ctx *ArgumentsContext) {}
