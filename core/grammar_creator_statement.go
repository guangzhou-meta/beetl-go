package core

import (
	"reflect"
)

import (
	"github.com/guangzhou-meta/beetl-go/core/om"
	"github.com/guangzhou-meta/beetl-go/util"
)

type VarAssignStatementI interface {
	StatementI

	GetVarAssignStatement() VarAssignStatementI
}

type VarAssignStatement struct {
	*Statement

	varIndex int
	exp      ExpressionI
}

func NewVarAssignStatement(exp ExpressionI, token *GrammarToken) *VarAssignStatement {
	return &VarAssignStatement{
		Statement: NewStatement(token),
		exp:       exp,
	}
}

func (s *VarAssignStatement) Execute(ctx *Context) {
	ctx.Vars[s.varIndex] = s.exp.Evaluate(ctx)
}

func (s *VarAssignStatement) SetVarIndex(index int) {
	s.varIndex = index
}

func (s *VarAssignStatement) GetVarIndex() int {
	return s.varIndex
}

func (s *VarAssignStatement) GetVarAssignStatement() VarAssignStatementI {
	return s
}

type VarAssignSeqStatement struct {
	*Statement

	sts []StatementI
}

func NewVarAssignSeqStatement(sts []StatementI, token *GrammarToken) *VarAssignSeqStatement {
	return &VarAssignSeqStatement{
		Statement: NewStatement(token),
		sts:       sts,
	}
}

func (s *VarAssignSeqStatement) Execute(ctx *Context) {
	for _, st := range s.sts {
		st.Execute(ctx)
	}
}

type VarRefAssignStatement struct {
	*VarAssignStatement

	varIndex         int
	exp              ExpressionI
	varRef           *VarRef
	lastVarAttribute VarAttributeI
}

func NewVarRefAssignStatement(exp ExpressionI, varRef *VarRef) *VarRefAssignStatement {
	return &VarRefAssignStatement{
		VarAssignStatement: NewVarAssignStatement(exp, exp.GetASTNode().Token),
		exp:                exp,
		varRef:             varRef,
		lastVarAttribute:   varRef.attributes[len(varRef.attributes)-1],
	}
}

func (s *VarRefAssignStatement) Execute(ctx *Context) {
	value := s.exp.Evaluate(ctx)
	obj := s.varRef.EvaluateUntilLast(ctx)
	var key interface{}
	if vsa, ok := s.lastVarAttribute.(*VarSquareAttribute); ok {
		key = vsa.exp.Evaluate(ctx)
	} else {
		key = s.lastVarAttribute.GetName()
	}

	if obj == nil {
		panic(NewBeetlException(NULL, nil).SetToken(s.varRef.Token))
	}

	aa := om.BuildFiledAccessor(util.UnpackType(reflect.TypeOf(obj)).Kind())
	aa.SetValue(obj, key, value)
}

type ReturnStatement struct {
	*Statement

	exp ExpressionI
}

func NewReturnStatement(exp ExpressionI, token *GrammarToken) *ReturnStatement {
	return &ReturnStatement{
		Statement: NewStatement(token),
		exp:       exp,
	}
}

func (rs *ReturnStatement) Execute(ctx *Context) {
	ctx.GotoFlag = IgotoReturn
	if rs.exp != nil {
		value := rs.exp.Evaluate(ctx)
		ctx.Vars[len(ctx.Vars)-1] = value
	}
}

type BreakStatement struct {
	*Statement
}

func NewBreakStatement(token *GrammarToken) *BreakStatement {
	return &BreakStatement{
		Statement: NewStatement(token),
	}
}

func (bs *BreakStatement) Execute(ctx *Context) {
	ctx.GotoFlag = IgotoBreak
}

type ContinueStatement struct {
	*Statement
}

func NewContinueStatement(token *GrammarToken) *ContinueStatement {
	return &ContinueStatement{
		Statement: NewStatement(token),
	}
}

func (bs *ContinueStatement) Execute(ctx *Context) {
	ctx.GotoFlag = IgotoContinue
}

type StaticTextASTNode struct {
	*Statement

	textIndex int
}

func NewStaticTextASTNode(textIndex int, token *GrammarToken) *StaticTextASTNode {
	return &StaticTextASTNode{
		Statement: NewStatement(token),
		textIndex: textIndex,
	}
}

func (st *StaticTextASTNode) Execute(ctx *Context) {
	t := ctx.StaticTextArray[st.textIndex]
	ctx.byteWriter.WriteString(string(t.([]rune)))
	//ctx.byteWriter.Append(string(t.([]rune)))
}

type StaticTextByteASTNode struct {
	*Statement

	textIndex int
}

func NewStaticTextByteASTNode(textIndex int, token *GrammarToken) *StaticTextByteASTNode {
	return &StaticTextByteASTNode{
		Statement: NewStatement(token),
		textIndex: textIndex,
	}
}

func (st *StaticTextByteASTNode) Execute(ctx *Context) {
	t := ctx.StaticTextArray[st.textIndex]
	ctx.byteWriter.WriteString(string(t.([]byte)))
	//ctx.byteWriter.Append(string(t.([]byte)))
}

type IfStatement struct {
	*Statement

	condition     ExpressionI
	ifStatement   StatementI
	elseStatement StatementI
}

func NewIfStatement(condition ExpressionI, ifStatement StatementI, elseStatement StatementI, token *GrammarToken) *IfStatement {
	return &IfStatement{
		Statement:     NewStatement(token),
		condition:     condition,
		ifStatement:   ifStatement,
		elseStatement: elseStatement,
	}
}

func (is *IfStatement) Execute(ctx *Context) {
	value := is.condition.Evaluate(ctx)
	if b, ok := value.(bool); ok && b {
		is.ifStatement.Execute(ctx)
	} else if is.elseStatement != nil {
		is.elseStatement.Execute(ctx)
	}
}

type StatementExpression struct {
	*Statement

	exp ExpressionI
}

func NewStatementExpression(exp ExpressionI, token *GrammarToken) *StatementExpression {
	return &StatementExpression{
		Statement: NewStatement(token),
		exp:       exp,
	}
}

func (se *StatementExpression) Execute(ctx *Context) {
	se.exp.Evaluate(ctx)
}

type TryCatchStatement struct {
	*Statement

	tryPart   *BlockStatement
	catchPart *BlockStatement
	error     *VarDefineNode
}

func NewTryCatchStatement(tryPart *BlockStatement, catchPart *BlockStatement, error *VarDefineNode, token *GrammarToken) *TryCatchStatement {
	return &TryCatchStatement{
		Statement: NewStatement(token),
		tryPart:   tryPart,
		catchPart: catchPart,
		error:     error,
	}
}

func (tcs *TryCatchStatement) Execute(ctx *Context) {
	defer func() {
		if err := recover(); err != nil {
			if tcs.catchPart == nil {
				if e, ok := err.(*BeetlException); ok {
					panic(e)
				} else {
					msg := err.(error).Error()
					panic(NewBeetlException(ERROR, &msg).SetToken(tcs.tryPart.Token))
				}
			}
			tcs.executeCatch(ctx, err.(error))
		}
	}()
	tcs.tryPart.Execute(ctx)
}

func (tcs *TryCatchStatement) executeCatch(ctx *Context, err error) {
	if tcs.error != nil {
		if e, ok := err.(*BeetlException); ok {
			ctx.Vars[tcs.error.varIndex] = e
		} else {
			msg := err.Error()
			ctx.Vars[tcs.error.varIndex] = NewBeetlException(ERROR, &msg)
		}
	}
	tcs.catchPart.Execute(ctx)
}

type TagStatement struct {
	*Statement

	tagName string
	paras   []ExpressionI
	block   StatementI
}

func NewTagStatement(tagName string, paras []ExpressionI, block StatementI, token *GrammarToken) *TagStatement {
	return &TagStatement{
		Statement: NewStatement(token),
		tagName:   tagName,
		paras:     paras,
		block:     block,
	}
}

func (ts *TagStatement) Execute(ctx *Context) {
	tagFactory := ctx.gt.GetTagFactory(ts.tagName)
	tag := tagFactory.CreateTag()
	var args []interface{}
	if len(ts.paras) == 0 {
		args = make([]interface{}, 0)
	} else {
		args = make([]interface{}, len(ts.paras))
		for i := range args {
			args[i] = ts.paras[i].Evaluate(ctx)
		}
	}
	tag.Init(ctx, args, ts.block)
	ts.runTag(tag, ctx)
}

func (ts *TagStatement) runTag(tag TagI, ctx *Context) {
	tag.Render()
	tag.AfterRender()
}

type TagVarBindingStatement struct {
	*TagStatement

	varIndexs []*VarDefineNode
}

func NewTagVarBindingStatement(tagName string, paras []ExpressionI, block StatementI, varDefine []*VarDefineNode, token *GrammarToken) *TagVarBindingStatement {
	return &TagVarBindingStatement{
		TagStatement: NewTagStatement(tagName, paras, block, token),
		varIndexs:    varDefine,
	}
}

func (ts *TagVarBindingStatement) runTag(tag TagI, ctx *Context) {
	if htmlTag, ok := tag.(*HTMLTagVarBindingWrapper); ok {
		indexMap := make(map[string]int)
		for _, node := range ts.varIndexs {
			indexMap[node.Token.Text] = node.varIndex
		}
		htmlTag.MapName2Index(indexMap)
		tag.Render()
		tag.AfterRender()
	} else {
		msg := "tag必须是HTMLTagVarBindingWrapper"
		panic(NewBeetlException(ERROR, &msg).SetToken(ts.Token))
	}
}

type SwitchStatement struct {
	*Statement

	value          ExpressionI
	m              map[ExpressionI]*BlockStatement
	conditionsList []ExpressionI
	blocks         []*BlockStatement
	defaultBlock   *BlockStatement
}

func NewSwitchStatement(value ExpressionI, m map[ExpressionI]*BlockStatement, defaultBlock *BlockStatement, token *GrammarToken) *SwitchStatement {
	var conditionsList []ExpressionI
	var blocks []*BlockStatement
	for k, v := range m {
		conditionsList = append(conditionsList, k)
		blocks = append(blocks, v)
	}
	return &SwitchStatement{
		Statement:      NewStatement(token),
		m:              m,
		value:          value,
		conditionsList: conditionsList,
		blocks:         blocks,
		defaultBlock:   defaultBlock,
	}
}

func (ss *SwitchStatement) Execute(ctx *Context) {
	o := ss.value.Evaluate(ctx)
	if o == nil {
		panic(NewBeetlException(NULL, nil).SetToken(ss.value.GetASTNode().Token))
	}

	isMatch := false
	for _, exp := range ss.conditionsList {
		if isMatch || util.ObjectsAreEqualValues(o, exp.Evaluate(ctx)) {
			isMatch = true
			block := ss.m[exp]
			if block != nil {
				block.Execute(ctx)
				switch ctx.GotoFlag {
				case IgotoNormal:
					break
				case IgotoReturn:
					return
				case IgotoBreak:
					ctx.GotoFlag = IgotoNormal
					return
				}
			}
		}
	}
	if !isMatch && ss.defaultBlock != nil {
		ss.defaultBlock.Execute(ctx)
		switch ctx.GotoFlag {
		case IgotoNormal:
			break
		case IgotoReturn:
			return
		case IgotoBreak:
			ctx.GotoFlag = IgotoNormal
			return
		}
	}
}

type SelectStatement struct {
	*Statement

	value        ExpressionI
	conditions   []ExpressionI
	blocks       []*BlockStatement
	defaultBlock *BlockStatement
}

func NewSelectStatement(value ExpressionI, conditions []ExpressionI, blocks []*BlockStatement, defaultBlock *BlockStatement, token *GrammarToken) *SelectStatement {
	return &SelectStatement{
		Statement:    NewStatement(token),
		value:        value,
		conditions:   conditions,
		blocks:       blocks,
		defaultBlock: defaultBlock,
	}
}

func (ss *SelectStatement) Execute(ctx *Context) {
	var base interface{}
	if ss.value != nil {
		base = ss.value.Evaluate(ctx)
		if base == nil {
			panic(NewBeetlException(NULL, nil).SetToken(ss.value.GetASTNode().Token))
		}
	}
	isMatch := false
	for i := range ss.conditions {
		exp := ss.conditions[i]
		condValue := exp.Evaluate(ctx)
		if ss.value == nil {
			if b, ok := condValue.(bool); ok {
				isMatch = b
			} else {
				panic(NewBeetlException(BOOLEAN_EXPECTED_ERROR, nil).SetToken(exp.GetASTNode().Token))
			}
		} else {
			if util.ObjectsAreEqualValues(base, condValue) {
				isMatch = true
			}
		}

		if isMatch {
			block := ss.blocks[i]
			if block != nil {
				block.Execute(ctx)
			}
			break
		}
	}

	if !isMatch && ss.defaultBlock != nil {
		ss.defaultBlock.Execute(ctx)
	}
}
