package core

const (
	IGrammarConstantsAjax          = "Ajax"
	IGrammarConstantsSelect        = "Select"
	IGrammarConstantsSwitch        = "Switch"
	IGrammarConstantsVarTag        = "VarTag"
	IGrammarConstantsTag           = "Tag"
	IGrammarConstantsTry           = "Try"
	IGrammarConstantsIf            = "If"
	IGrammarConstantsWhile         = "While"
	IGrammarConstantsFor           = "For"
	IGrammarConstantsForIn         = "ForIn"
	IGrammarConstantsContinue      = "Continue"
	IGrammarConstantsVarAssignSeq  = "VarAssignSeq"
	IGrammarConstantsVarAssign     = "VarAssign"
	IGrammarConstantsVarRefAssign  = "VarRefAssign"
	IGrammarConstantsTextOutputSt  = "TextOutputSt"
	IGrammarConstantsTextOutputSt2 = "TextOutputSt2"
	IGrammarConstantsReturn        = "Return"
	IGrammarConstantsBreak         = "Break"
	/* Express */
	IGrammarConstantsVarRefOptimal      = "VarRefOptimal"
	IGrammarConstantsFormat             = "Format"
	IGrammarConstantsTemplateContent    = "TemplateContent"
	IGrammarConstantsVarRefAssignExp    = "VarRefAssignExp"
	IGrammarConstantsIncDec             = "IncDec"
	IGrammarConstantsNeg                = "Neg"
	IGrammarConstantsNot                = "Not"
	IGrammarConstantsOr                 = "Or"
	IGrammarConstantsAnd                = "And"
	IGrammarConstantsInstanceNativeCall = "InstanceNativeCall"
	IGrammarConstantsClassNativeCall    = "ClassNativeCall"
	IGrammarConstantsFunctionExp        = "FunctionExp"
	IGrammarConstantsJsonMap            = "JsonMap"
	IGrammarConstantsJsonArray          = "JsonArray"
	IGrammarConstantsArth               = "Arth"
	IGrammarConstantsTernary            = "Ternary"
	IGrammarConstantsCompare            = "Compare"
	IGrammarConstantsFunction           = "Function"
)

type GrammarCreatorI interface {
	GetGrammarCreator() GrammarCreatorI

	AddDisableGrammar(disableGrammar string)

	CreateVarAssignSeq(assings []*VarAssignStatement) *VarAssignSeqStatement
	CreateVarAssign(exp ExpressionI, token *GrammarToken) *VarAssignStatement
	CreateVarRefAssign(exp ExpressionI, varRef *VarRef) *VarRefAssignStatement
	CreateBlock(nodes []StatementI, token *GrammarToken) *BlockStatement
	CreateTextOutputSt(exp ExpressionI, format *FormatExpression) PlaceholderSTI
	CreateTextOutputSt2(exp ExpressionI, format *FormatExpression) *PlaceholderST
	CreateReturn(exp ExpressionI) *ReturnStatement
	CreateBreak(token *GrammarToken) *BreakStatement
	CreateContinue(token *GrammarToken) *ContinueStatement
	CreateForIn(forVar *VarDefineNode, exp ExpressionI, hasSafe bool, forPart StatementI, elseForPart StatementI) *ForStatement
	CreateFor(varAssignSeq *VarAssignSeqStatement, expInit []ExpressionI, condition ExpressionI, expUpdate []ExpressionI, forPart StatementI, elseforPart StatementI, token *GrammarToken) *GeneralForStatement
	CreateWhile(exp ExpressionI, whileBody StatementI, token *GrammarToken) *WhileStatement
	CreateStaticText(textIndex int, token *GrammarToken) *StaticTextASTNode
	CreateStaticByteText(textIndex int, token *GrammarToken) *StaticTextByteASTNode
	CreateIf(condtion ExpressionI, ifStatement StatementI, elseStatement StatementI, token *GrammarToken) *IfStatement
	CreateStatementExpression(expression ExpressionI) *StatementExpression
	CreateTry(tryPart *BlockStatement, catchPart *BlockStatement, error *VarDefineNode, token *GrammarToken) *TryCatchStatement
	CreateTag(tagName string, expList []ExpressionI, block StatementI, token *GrammarToken) *TagStatement
	CreateVarTag(tagName string, expList []ExpressionI, block StatementI, varDefine []*VarDefineNode, token *GrammarToken) *TagVarBindingStatement
	CreateSwitch(value ExpressionI, ebMap map[ExpressionI]*BlockStatement, defaultBlock *BlockStatement, token *GrammarToken) *SwitchStatement
	CreateSelect(value ExpressionI, conditions []ExpressionI, blocks []*BlockStatement, defaultBlock *BlockStatement, token *GrammarToken) *SelectStatement
	CreateAjax(block *BlockStatement, token *GrammarToken, defaultRender bool) *AjaxStatement
	First(resource ResourceI) StatementI
	Last(resource ResourceI) StatementI

	CreateFunction(name string, exps []ExpressionI, vas []VarAttributeI, hasSafe bool, safeExp ExpressionI, token *GrammarToken) FunctionExpressionI
	CreateLiteral(value interface{}, token *GrammarToken) *Literal
	CreateCompare(a ExpressionI, b ExpressionI, mode int, token *GrammarToken) *CompareExpression
	CreateTernary(condtion ExpressionI, a ExpressionI, b ExpressionI, token *GrammarToken) *TernaryExpression
	CreateArth(a ExpressionI, b ExpressionI, mode int, token *GrammarToken) *ArthExpression
	CreateJasonArray(list []ExpressionI, token *GrammarToken) *JsonArrayExpression
	CreateJsonMap(m map[string]ExpressionI, token *GrammarToken) *JsonMapExpression
	CreateFunctionExp(name string, exps []ExpressionI, vas []VarAttributeI, hasSafe bool, safeExp ExpressionI, token *GrammarToken) FunctionExpressionI
	CreateClassNativeCall(clsNode *ClassNode, chain []NativeNodeI, token *GrammarToken) *NativeCallExpression
	CreateInstanceNativeCall(insNode *InstanceNode, chain []NativeNodeI, token *GrammarToken) *NativeCallExpression
	CreateAnd(exp1 ExpressionI, exp2 ExpressionI, token *GrammarToken) *AndExpression
	CreateOr(exp1 ExpressionI, exp2 ExpressionI, token *GrammarToken) *OrExpression
	CreateNot(exp ExpressionI, token *GrammarToken) *NotBooleanExpression
	CreateNeg(exp ExpressionI, token *GrammarToken) *NegExpression
	CreateIncDec(isInc bool, returnOrginal bool, token *GrammarToken) *IncDecExpression
	CreateVarRef(attributes []VarAttributeI, hasSafe bool, safe ExpressionI, token *GrammarToken, firstToken *GrammarToken) *VarRef
	CreateVarRefAssignExp(exp ExpressionI, varRef *VarRef) *VarRefAssignExpress
	CreateTemplateContent(block *BlockStatement, token *GrammarToken) *ContentBodyExpression
	CreateFormat(name string, pattern string, token *GrammarToken) *FormatExpression
}

type GrammarCreator struct {
	disableGrammarSet map[string]interface{}
}

func NewGrammarCreator() *GrammarCreator {
	return &GrammarCreator{
		disableGrammarSet: make(map[string]interface{}),
	}
}

func (gc *GrammarCreator) GetGrammarCreator() GrammarCreatorI {
	return gc
}

func (gc *GrammarCreator) AddDisableGrammar(disableGrammar string) {
	gc.disableGrammarSet[disableGrammar] = nil
}

func (gc *GrammarCreator) CreateVarAssignSeq(assings []*VarAssignStatement) *VarAssignSeqStatement {
	gc.disableSyntaxCheck(IGrammarConstantsVarAssignSeq)
	var sts []StatementI
	for _, i := range assings {
		sts = append(sts, i)
	}
	return NewVarAssignSeqStatement(sts, nil)
}

func (gc *GrammarCreator) CreateVarAssign(exp ExpressionI, token *GrammarToken) *VarAssignStatement {
	gc.disableSyntaxCheck(IGrammarConstantsVarAssign)
	return NewVarAssignStatement(exp, token)
}

func (gc *GrammarCreator) CreateVarRefAssign(exp ExpressionI, varRef *VarRef) *VarRefAssignStatement {
	gc.disableSyntaxCheck(IGrammarConstantsVarRefAssign)
	return NewVarRefAssignStatement(exp, varRef)
}

func (gc *GrammarCreator) CreateBlock(nodes []StatementI, token *GrammarToken) *BlockStatement {
	return NewBlockStatement(nodes, token)
}

func (gc *GrammarCreator) CreateTextOutputSt(exp ExpressionI, format *FormatExpression) PlaceholderSTI {
	gc.disableSyntaxCheck(IGrammarConstantsTextOutputSt)
	return NewPlaceholderST(exp, format, nil)
}

func (gc *GrammarCreator) CreateTextOutputSt2(exp ExpressionI, format *FormatExpression) *PlaceholderST {
	gc.disableSyntaxCheck(IGrammarConstantsTextOutputSt2)
	return NewPlaceholderST(exp, format, nil)
}

func (gc *GrammarCreator) CreateReturn(exp ExpressionI) *ReturnStatement {
	gc.disableSyntaxCheck(IGrammarConstantsReturn)
	return NewReturnStatement(exp, nil)
}

func (gc *GrammarCreator) CreateBreak(token *GrammarToken) *BreakStatement {
	gc.disableSyntaxCheck(IGrammarConstantsBreak)
	return NewBreakStatement(token)
}

func (gc *GrammarCreator) CreateContinue(token *GrammarToken) *ContinueStatement {
	gc.disableSyntaxCheck(IGrammarConstantsContinue)
	return NewContinueStatement(token)
}

func (gc *GrammarCreator) CreateForIn(forVar *VarDefineNode, exp ExpressionI, hasSafe bool, forPart StatementI, elseForPart StatementI) *ForStatement {
	gc.disableSyntaxCheck(IGrammarConstantsForIn)
	return NewForStatement(forVar, exp, hasSafe, forPart, elseForPart, forVar.Token)
}

func (gc *GrammarCreator) CreateFor(varAssignSeq *VarAssignSeqStatement, expInit []ExpressionI, condition ExpressionI, expUpdate []ExpressionI, forPart StatementI, elseforPart StatementI, token *GrammarToken) *GeneralForStatement {
	gc.disableSyntaxCheck(IGrammarConstantsFor)
	return NewGeneralForStatement(varAssignSeq, expInit, condition, expUpdate, forPart, elseforPart, token)
}

func (gc *GrammarCreator) CreateWhile(exp ExpressionI, whileBody StatementI, token *GrammarToken) *WhileStatement {
	gc.disableSyntaxCheck(IGrammarConstantsWhile)
	return NewWhileStatement(exp, whileBody, token)
}

func (gc *GrammarCreator) CreateStaticText(textIndex int, token *GrammarToken) *StaticTextASTNode {
	return NewStaticTextASTNode(textIndex, token)
}

func (gc *GrammarCreator) CreateStaticByteText(textIndex int, token *GrammarToken) *StaticTextByteASTNode {
	return NewStaticTextByteASTNode(textIndex, token)
}

func (gc *GrammarCreator) CreateIf(condtion ExpressionI, ifStatement StatementI, elseStatement StatementI, token *GrammarToken) *IfStatement {
	gc.disableSyntaxCheck(IGrammarConstantsIf)
	return NewIfStatement(condtion, ifStatement, elseStatement, token)
}

func (gc *GrammarCreator) CreateStatementExpression(expression ExpressionI) *StatementExpression {
	return NewStatementExpression(expression, nil)
}

func (gc *GrammarCreator) CreateTry(tryPart *BlockStatement, catchPart *BlockStatement, error *VarDefineNode, token *GrammarToken) *TryCatchStatement {
	gc.disableSyntaxCheck(IGrammarConstantsTry)
	return NewTryCatchStatement(tryPart, catchPart, error, token)
}

func (gc *GrammarCreator) CreateTag(tagName string, expList []ExpressionI, block StatementI, token *GrammarToken) *TagStatement {
	gc.disableSyntaxCheck(IGrammarConstantsTag)
	return NewTagStatement(tagName, expList, block, token)
}

func (gc *GrammarCreator) CreateVarTag(tagName string, expList []ExpressionI, block StatementI, varDefine []*VarDefineNode, token *GrammarToken) *TagVarBindingStatement {
	gc.disableSyntaxCheck(IGrammarConstantsVarTag)
	return NewTagVarBindingStatement(tagName, expList, block, varDefine, token)
}

func (gc *GrammarCreator) CreateSwitch(value ExpressionI, ebMap map[ExpressionI]*BlockStatement, defaultBlock *BlockStatement, token *GrammarToken) *SwitchStatement {
	gc.disableSyntaxCheck(IGrammarConstantsSwitch)
	return NewSwitchStatement(value, ebMap, defaultBlock, token)
}

func (gc *GrammarCreator) CreateSelect(value ExpressionI, conditions []ExpressionI, blocks []*BlockStatement, defaultBlock *BlockStatement, token *GrammarToken) *SelectStatement {
	gc.disableSyntaxCheck(IGrammarConstantsSelect)
	return NewSelectStatement(value, conditions, blocks, defaultBlock, token)
}

func (gc *GrammarCreator) CreateAjax(block *BlockStatement, token *GrammarToken, defaultRender bool) *AjaxStatement {
	gc.disableSyntaxCheck(IGrammarConstantsAjax)
	return NewAjaxStatement(block, token, defaultRender)
}

// Express //

func (gc *GrammarCreator) CreateFunction(name string, exps []ExpressionI, vas []VarAttributeI, hasSafe bool, safeExp ExpressionI, token *GrammarToken) FunctionExpressionI {
	gc.disableSyntaxCheck(IGrammarConstantsFunction)
	return NewFunctionExpression(name, exps, vas, hasSafe, safeExp, token)
}

func (gc *GrammarCreator) CreateLiteral(value interface{}, token *GrammarToken) *Literal {
	return NewLiteral(value, token)
}

func (gc *GrammarCreator) CreateCompare(a ExpressionI, b ExpressionI, mode int, token *GrammarToken) *CompareExpression {
	gc.disableSyntaxCheck(IGrammarConstantsCompare)
	return NewCompareExpression(a, b, mode, token)
}

func (gc *GrammarCreator) CreateTernary(condtion ExpressionI, a ExpressionI, b ExpressionI, token *GrammarToken) *TernaryExpression {
	gc.disableSyntaxCheck(IGrammarConstantsTernary)
	return NewTernaryExpression(condtion, a, b, token)
}

func (gc *GrammarCreator) CreateArth(a ExpressionI, b ExpressionI, mode int, token *GrammarToken) *ArthExpression {
	gc.disableSyntaxCheck(IGrammarConstantsArth)
	return NewArthExpression(a, b, mode, token)
}

func (gc *GrammarCreator) CreateJasonArray(list []ExpressionI, token *GrammarToken) *JsonArrayExpression {
	gc.disableSyntaxCheck(IGrammarConstantsJsonArray)
	return NewJsonArrayExpression(list, token)
}

func (gc *GrammarCreator) CreateJsonMap(m map[string]ExpressionI, token *GrammarToken) *JsonMapExpression {
	gc.disableSyntaxCheck(IGrammarConstantsJsonMap)
	return NewJsonMapExpression(m, token)
}

func (gc *GrammarCreator) CreateFunctionExp(name string, exps []ExpressionI, vas []VarAttributeI, hasSafe bool, safeExp ExpressionI, token *GrammarToken) FunctionExpressionI {
	gc.disableSyntaxCheck(IGrammarConstantsFunctionExp)
	return NewFunctionExpression(name, exps, vas, hasSafe, safeExp, token)
}

func (gc *GrammarCreator) CreateClassNativeCall(clsNode *ClassNode, chain []NativeNodeI, token *GrammarToken) *NativeCallExpression {
	gc.disableSyntaxCheck(IGrammarConstantsClassNativeCall)
	return NewNativeCallExpression(nil, clsNode, chain, token)
}

func (gc *GrammarCreator) CreateInstanceNativeCall(insNode *InstanceNode, chain []NativeNodeI, token *GrammarToken) *NativeCallExpression {
	gc.disableSyntaxCheck(IGrammarConstantsInstanceNativeCall)
	return NewNativeCallExpression(insNode, nil, chain, token)
}

func (gc *GrammarCreator) CreateAnd(exp1 ExpressionI, exp2 ExpressionI, token *GrammarToken) *AndExpression {
	gc.disableSyntaxCheck(IGrammarConstantsAnd)
	return NewAndExpression(exp1, exp2, token)
}

func (gc *GrammarCreator) CreateOr(exp1 ExpressionI, exp2 ExpressionI, token *GrammarToken) *OrExpression {
	gc.disableSyntaxCheck(IGrammarConstantsOr)
	return NewOrExpression(exp1, exp2, token)
}

func (gc *GrammarCreator) CreateNot(exp ExpressionI, token *GrammarToken) *NotBooleanExpression {
	gc.disableSyntaxCheck(IGrammarConstantsNot)
	return NewNotBooleanExpression(exp, token)
}

func (gc *GrammarCreator) CreateNeg(exp ExpressionI, token *GrammarToken) *NegExpression {
	gc.disableSyntaxCheck(IGrammarConstantsNeg)
	return NewNegExpression(exp, token)
}

func (gc *GrammarCreator) CreateIncDec(isInc bool, returnOrginal bool, token *GrammarToken) *IncDecExpression {
	gc.disableSyntaxCheck(IGrammarConstantsIncDec)
	return NewIncDecExpression(isInc, returnOrginal, token)
}

func (gc *GrammarCreator) CreateVarRef(attributes []VarAttributeI, hasSafe bool, safe ExpressionI, token *GrammarToken, firstToken *GrammarToken) *VarRef {
	gc.disableSyntaxCheck(IGrammarConstantsVarRefOptimal)
	return NewVarRef(attributes, hasSafe, safe, token, firstToken)
}

func (gc *GrammarCreator) CreateVarRefAssignExp(exp ExpressionI, varRef *VarRef) *VarRefAssignExpress {
	gc.disableSyntaxCheck(IGrammarConstantsVarRefAssignExp)
	return NewVarRefAssignExpress(exp, varRef)
}

func (gc *GrammarCreator) CreateTemplateContent(block *BlockStatement, token *GrammarToken) *ContentBodyExpression {
	gc.disableSyntaxCheck(IGrammarConstantsTemplateContent)
	return NewContentBodyExpression(block, token)
}

func (gc *GrammarCreator) CreateFormat(name string, pattern string, token *GrammarToken) *FormatExpression {
	gc.disableSyntaxCheck(IGrammarConstantsFormat)
	return NewFormatExpression(&name, &pattern, token)
}

func (gc *GrammarCreator) disableSyntaxCheck(ast string) {
	if _, ok := gc.disableGrammarSet[ast]; ok {
		msg := "语法 " + ast + "不允许"
		panic(NewBeetlException(GRAMMAR_NOT_ALLOWED, &msg))
	}
}

func (gc *GrammarCreator) First(resource ResourceI) StatementI {
	return nil
}

func (gc *GrammarCreator) Last(resource ResourceI) StatementI {
	return nil
}
