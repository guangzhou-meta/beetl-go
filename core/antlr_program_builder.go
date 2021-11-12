package core

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

import (
	"github.com/guangzhou-meta/beetl-go/parser"
)

import (
	"github.com/guangzhou-meta/antlr4/runtime/Go/antlr"
	"github.com/guangzhou-meta/go-lib/arrays"
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

var (
	safeParameters = map[string]interface{}{
		"isEmpty":    nil,
		"isNotEmpty": nil,
		"isBlank":    nil,
	}
)

type AntlrProgramBuilder struct {
	data            *ProgramMetaData
	pbCtx           *ProgramBuilderContext
	EmptyExpression []ExpressionI
	gt              *GroupTemplate
	endStatment     *EndStatement
	gc              GrammarCreatorI
}

func NewAntlrProgramBuilder(gt *GroupTemplate, gc GrammarCreatorI) *AntlrProgramBuilder {
	return &AntlrProgramBuilder{
		gt:              gt,
		gc:              gc,
		EmptyExpression: make([]ExpressionI, 0),
		endStatment:     NewEndStatement(),
		pbCtx:           NewProgramBuilderContext(),
		data:            NewProgramMetaData(),
	}
}

func (b *AntlrProgramBuilder) Build(tree antlr.ParseTree, resource ResourceI) *ProgramMetaData {
	size := tree.GetChildCount() - 1
	ls := make([]StatementI, 0)
	firstStatement := b.gc.First(resource)
	if firstStatement != nil {
		ls = append(ls, firstStatement)
	}

	for i := 0; i < size; i++ {
		st := b.parseStatement(tree.GetChild(i).(antlr.ParserRuleContext))
		if st != nil {
			ls = append(ls, st)
		}
	}

	lastStatement := b.gc.Last(resource)
	if lastStatement != nil {
		ls = append(ls, lastStatement)
	}

	switch b.pbCtx.current.gotoValue {
	case IgotoReturn, IgotoBreak:
		b.data.HasGoto = true
	}

	b.pbCtx.AnzlyszeGlobal()
	b.pbCtx.AnzlyszeLocal()
	b.data.varIndexSize = b.pbCtx.varIndexSize
	b.data.tempVarStartIndex = len(b.pbCtx.globalIndexMap)
	b.data.statements = ls
	b.data.globalIndexMap = b.pbCtx.globalIndexMap
	b.data.SetTemplateRootScopeIndexMap(b.pbCtx.rootIndexMap)
	return b.data
}

func (b *AntlrProgramBuilder) parseStatement(node antlr.ParserRuleContext) StatementI {
	if node == nil {
		return nil
	}
	switch nv := node.(type) {
	case *parser.VarStContext:
		return b.parseVarSt(nv)
	case *parser.BlockStContext:
		return b.parseBlock(nv.Block().(*parser.BlockContext).AllStatement(), node)
	case *parser.TextOutputStContext:
		return b.parseTextOutputSt(nv)
	case *parser.ReturnStContext:
		expCtx := nv.Expression().(*parser.ExpressionContext)
		var exp ExpressionI
		if expCtx != nil {
			exp = b.parseExpress(expCtx)
		}
		st := b.gc.CreateReturn(exp)
		b.pbCtx.current.gotoValue = IgotoReturn
		return st
	case *parser.BreakStContext:
		st := b.gc.CreateBreak(nil)
		if b.pbCtx.current.gotoValue != IgotoReturn {
			b.pbCtx.current.gotoValue = IgotoBreak
		}
		return st
	case *parser.ContinueStContext:
		st := b.gc.CreateContinue(nil)
		if b.pbCtx.current.gotoValue != IgotoReturn {
			b.pbCtx.current.gotoValue = IgotoContinue
		}
		return st
	case *parser.ForStContext:
		return b.parseForSt(nv)
	case *parser.StaticOutputStContext:
		cst := nv.ConstantsTextStatment().(*parser.ConstantsTextStatmentContext)
		str := cst.DecimalLiteral().GetSymbol().GetText()
		position, _ := strconv.ParseInt(str, 10, 32)
		if b.gt.conf.directByteOutput {
			return b.gc.CreateStaticByteText(int(position), nil)
		}
		return b.gc.CreateStaticText(int(position), nil)
	case *parser.IfStContext:
		return b.parseIf(nv)
	case *parser.StatmentExpStContext:
		return NewStatementExpression(b.parseExpress(nv.StatementExpression().(*parser.StatementExpressionContext).Expression()), nil)
	case *parser.DirectiveStContext:
		return b.parseDirectiveStatement(nv)
	case *parser.CommentTagStContext:
		return nil
	case *parser.TryStContext:
		return b.parseTryCatch(nv)
	case *parser.FunctionTagStContext:
		return b.parseTag(nv.FunctionTagCall().(*parser.FunctionTagCallContext))
	case *parser.WhileStContext:
		return b.parseWhile(nv)
	case *parser.AssignStContext:
		vas := b.parseAssign(nv.AssignMent())
		if _, ok := vas.(*VarRefAssignStatement); !ok {
			b.pbCtx.SetVarPosition(vas.(*VarAssignStatement).GetASTNode().Token.Text, vas.(ASTNodeI))
		}
	case *parser.SiwchStContext:
		return b.parseSwitch(nv)
	case *parser.SelectStContext:
		return b.parseSelect(nv)
	case *parser.AjaxStContext:
		return b.parseAjax(nv)
	case *parser.EndContext:
		return b.endStatment
	default:
		panic("未识别，确认模板书写是否正确")
	}
	return nil
}

func (b *AntlrProgramBuilder) parseVarSt(varSt *parser.VarStContext) *VarAssignSeqStatement {
	return b.parseVarDeclareList(varSt.VarDeclareList().(*parser.VarDeclareListContext))
}

func (b *AntlrProgramBuilder) parseVarDeclareList(ctx *parser.VarDeclareListContext) *VarAssignSeqStatement {
	list := ctx.AllAssignMent()
	var listNode []*VarAssignStatement
	for _, amc := range list {
		vas := b.parseAssign(amc)
		listNode = append(listNode, vas.(*VarAssignStatement))
		if _, ok := vas.(*VarRefAssignStatement); !ok {
			b.registerNewVar(vas.(ASTNodeI))
		}
	}
	seq := b.gc.CreateVarAssignSeq(listNode)
	return seq
}

func (b *AntlrProgramBuilder) parseAssign(amc parser.IAssignMentContext) StatementI {
	var vas VarAssignStatementI = nil
	switch v := amc.(type) {
	case *parser.AssignGeneralInStContext:
		agc := v
		agcv := agc.GeneralAssignExp().(*parser.GeneralAssignExpContext)
		expCtx := agcv.Expression()
		exp := b.parseExpress(expCtx)
		varRefCtx := agcv.VarRef().(*parser.VarRefContext)
		if len(varRefCtx.GetChildren()) == 1 {
			token := varRefCtx.Identifier().GetSymbol()
			vas = b.gc.CreateVarAssign(exp, b.getBTToken(token))
		} else {
			ref := b.parseVarRefInLeftExpression(varRefCtx)
			vas = b.gc.CreateVarRefAssign(exp, ref)
		}
	case *parser.AssignIdContext:
		vas = NewVarAssignStatement(NULLLiteral, b.getBTToken(v.Identifier().GetSymbol()))
	case *parser.AssignTemplateVarContext:
		blockCtx := v.Block().(*parser.BlockContext)
		block := b.parseBlock(blockCtx.AllStatement(), blockCtx)
		token := b.getBTToken(v.Identifier().GetSymbol())
		bodyExp := b.gc.CreateTemplateContent(block.(*BlockStatement), token)
		vas = b.gc.CreateVarAssign(bodyExp, token)
	default:
		panic(fmt.Sprintf("不支持在 %d", v.GetStart().GetLine()))
	}
	return vas
}

func (b *AntlrProgramBuilder) parseExpress(ctx parser.IExpressionContext) ExpressionI {
	if ctx == nil {
		return nil
	}
	switch v := ctx.(type) {
	case *parser.LiteralExpContext:
		return b.parseLiteralExpress(v.Literal())
	case *parser.VarRefExpContext:
		return b.parseVarRefExpression(v.VarRef())
	case *parser.CompareExpContext:
		return b.parseCompareExpression(v)
	case *parser.TernaryExpContext:
		return b.parseTernaryExpression(v)
	case *parser.MuldivmodExpContext:
		return b.parseMuldivmodExpression(v)
	case *parser.AddminExpContext:
		return b.parsePlusMins(v)
	case *parser.ParExpContext:
		return b.parseExpress(v.Expression())
	case *parser.FunctionCallExpContext:
		return b.parseFunExp(v.FunctionCall())
	case *parser.JsonExpContext:
		return b.parseJson(v.Json().(*parser.JsonContext))
	case *parser.NativeCallExpContext:
		panic("unsupported")
	case *parser.AndExpContext:
		return b.parseAndExpression(v)
	case *parser.OrExpContext:
		return b.parseOrExpression(v)
	case *parser.NotExpContext:
		return b.parseNotExpression(v)
	case *parser.NegExpContext:
		return b.parseNegExpression(v)
	case *parser.IncDecOneContext:
		return b.parseIncDecOneContext(v)
	case *parser.OneIncDecContext:
		return b.parseOneIncDecContext(v)
	case *parser.AssignGeneralInExpContext:
		return b.parseAssingInExp(v)
	case *parser.AssingSelfExpContext:
		return b.parseSelfAssingInExp(v)
	default:
		panic("unsupported")
	}
}

func (b *AntlrProgramBuilder) parseLiteralExpress(ctx parser.ILiteralContext) ExpressionI {
	tree := ctx.GetChild(0)
	var value interface{}
	switch t := tree.(type) {
	case antlr.TerminalNode:
		node := t.GetSymbol()
		strValue := node.GetText()
		tokenType := node.GetTokenType()
		switch tokenType {
		case parser.BeetlParserStringLiteral:
			value = b.getStringValue(strValue)
		case parser.BeetlParserFloatingPointLiteral:
			if b.isHighScaleNumber(strValue) {
				value, _ = big.NewFloat(0).SetString(strValue)
			} else {
				value, _ = strconv.ParseFloat(strValue, 64)
			}
		case parser.BeetlParserDecimalLiteral:
			if b.isHighScaleNumber(strValue) {
				value, _ = big.NewFloat(0).SetString(strValue)
			} else {
				value, _ = strconv.ParseInt(strValue, 10, 64)
			}
		case parser.BeetlParserNULL:
			value = nil
		}
	default:
		blc := tree.(*parser.BooleanLiteralContext)
		strValue := blc.GetChild(0).(antlr.ParserRuleContext).GetText()
		value, _ = strconv.ParseBool(strValue)
	}
	if value == nil {
		return NULLLiteral
	} else {
		return b.gc.CreateLiteral(value, b.getBTToken(ctx.GetStart()))
	}
}

func (b *AntlrProgramBuilder) getStringValue(strValue string) string {
	ch := []rune(strValue)
	sb := estrings.NewStringBuilderCap(len(ch))
	for i := 1; i < len(ch)-1; i++ {
		c := ch[i]
		if c == '\\' {
			i++
			spec := ch[i]
			var real rune
			switch spec {
			case '"', '\'', '\\':
				real = spec
			case 't':
				real = '\t'
			case 'r':
				real = '\r'
			case 'n':
				real = '\n'
			case 'b':
				real = '\b'
			case 'f':
				real = '\f'
			default:
				panic("不支持的转义符号" + strValue)
			}
			sb.Append(real)
		} else {
			sb.Append(c)
		}
	}
	return sb.String()
}

func (b *AntlrProgramBuilder) isHighScaleNumber(strValue string) bool {
	buf := []rune(strValue)
	c := buf[len(buf)-1]
	return c == 'h' || c == 'H'
}

func (b *AntlrProgramBuilder) getBTToken(start antlr.Token) *GrammarToken {
	return NewGrammarToken(start.GetText(), start.GetLine(), start.GetColumn())
}

func (b *AntlrProgramBuilder) getBTGToken(text string, line int) *GrammarToken {
	return NewGrammarToken(text, line, 0)
}

func (b *AntlrProgramBuilder) parseVarRefExpression(varRef parser.IVarRefContext) ExpressionI {
	varRefCtx := varRef.(*parser.VarRefContext)
	soctx := varRefCtx.Safe_output()
	hasSafe := false
	var safeExp ExpressionI
	if soctx != nil {
		safeExp = b.parseSafeOutput(soctx)
		hasSafe = true
	}
	if b.pbCtx.isSafeOutput {
		hasSafe = true
	}
	list := varRefCtx.AllVarAttribute()
	vas := b.parseVarAttribute(list)
	if len(vas) > 0 {
		first := vas[0]
		switch first.(type) {
		case *VarSquareAttribute, *VarVirtualAttribute:
			break
		default:
			b.pbCtx.SetVarAttr(varRefCtx.Identifier().GetText(), first.GetASTNode().Token.Text)
		}
	}
	varv := b.gc.CreateVarRef(vas, hasSafe, safeExp, b.getBTGToken(varRefCtx.GetText(), varRefCtx.Identifier().GetSymbol().GetLine()), b.getBTToken(varRefCtx.Identifier().GetSymbol()))
	b.pbCtx.SetVarPosition(varRefCtx.Identifier().GetText(), varv)
	return varv
}

func (b *AntlrProgramBuilder) parseSafeOutput(soctx parser.ISafe_outputContext) ExpressionI {
	list := soctx.GetChildren()
	var safeExp ExpressionI
	if len(list) == 1 {
		safeExp = nil
	} else {
		allowExp := list[1].(*parser.Safe_allow_expContext)
		if allowExp.Literal() != nil {
			safeExp = b.parseLiteralExpress(allowExp.Literal())
		} else if allowExp.NativeCall() != nil {
			panic("unsupported")
		} else if allowExp.FunctionCall() != nil {
			safeExp = b.parseFunExp(allowExp.FunctionCall())
		} else if allowExp.Expression() != nil {
			safeExp = b.parseExpress(allowExp.Expression())
		} else if allowExp.VarRef() != nil {
			safeExp = b.parseVarRefExpression(allowExp.VarRef())
		}
	}
	return safeExp
}

func (b *AntlrProgramBuilder) parseFunExp(ctx parser.IFunctionCallContext) ExpressionI {
	fCtx := ctx.(*parser.FunctionCallContext)
	expListCtx := fCtx.ExpressionList()
	exps := b.getExpressionList(expListCtx)
	vaListCtx := fCtx.AllVarAttribute()
	soctx := fCtx.Safe_output()
	var safeExp ExpressionI
	hasSafe := false
	if soctx != nil {
		safeExp = b.parseSafeOutput(soctx)
		hasSafe = true
	}

	if b.pbCtx.isSafeOutput {
		hasSafe = true
	}
	vs := b.parseVarAttribute(vaListCtx)
	idList := fCtx.FunctionNs().(*parser.FunctionNsContext).AllIdentifier()
	nsId := b.getID(idList)
	btToken := b.getBTGToken(nsId, fCtx.GetStart().GetLine())
	if _, ok := safeParameters[nsId]; ok {
		if len(exps) != 0 {
			for _, one := range exps {
				if ref, ok := one.(*VarRef); ok && !ref.hasSafe {
					ref.hasSafe = true
					ref.safe = nil
				}
			}
		}
	} else {
		switch nsId {
		case "has":
			if len(exps) != 0 {
				for i := range exps {
					one := exps[i]
					if ref, ok := one.(*VarRef); ok {
						if len(ref.attributes) != 0 {
							msg := "has函数用于判断全局变量是否存在，不能判断其属性是否有值，可以使用安全输出符号或者isEmpty函数"
							panic(NewBeetlException(HAS_CALL_ILLEGAL, &msg).SetToken(ref.Token))
						}
						name := ref.Token.Text
						newExp := b.gc.CreateLiteral(name, ref.Token)
						exps[i] = newExp
					} else {
						msg := "has函数用于判断全局变量是否存在,请传入一个全局变量名"
						panic(NewBeetlException(HAS_CALL_ILLEGAL, &msg).SetToken(one.GetASTNode().Token))
					}
				}
			}
		case "debug":
			l := b.gc.CreateLiteral(btToken.Line, btToken)
			length := len(exps)
			newExps := make([]ExpressionI, length+2)
			arrays.CopyFrom(exps, 0, newExps, 0, length)
			expStr := b.getExpressionString(expListCtx)
			newExps[length] = b.gc.CreateLiteral(expStr, btToken)
			newExps[length+1] = l
			for i := range exps {
				if _, ok := exps[i].(*VarRef); !ok {
					exps[i] = nil
				}
			}
			exps = newExps
		case "decode":
			if len(exps) >= 4 {
				newExps := make([]ExpressionI, len(exps))
				newExps[0] = exps[0]
				newExps[1] = exps[1]
				for i := 2; i < len(exps); i++ {
					newExps[i] = NewExpressionRuntime(exps[i])
				}
				exps = newExps
			}
		}
	}
	fe := b.gc.CreateFunctionExp(nsId, exps, vs, hasSafe, safeExp, btToken)
	return fe
}

func (b *AntlrProgramBuilder) getExpressionList(expListCtx parser.IExpressionListContext) []ExpressionI {
	if expListCtx == nil {
		return b.EmptyExpression
	}
	ctx := expListCtx.(*parser.ExpressionListContext)
	ecList := ctx.AllExpression()
	var exps []ExpressionI
	for _, ex := range ecList {
		exps = append(exps, b.parseExpress(ex))
	}
	return exps
}

func (b *AntlrProgramBuilder) getExpressionString(expListCtx parser.IExpressionListContext) []string {
	if expListCtx == nil {
		return make([]string, 0)
	}
	ctx := expListCtx.(*parser.ExpressionListContext)
	ecList := ctx.AllExpression()
	var exps []string
	for _, ex := range ecList {
		exps = append(exps, ex.GetText())
	}
	return exps
}

func (b *AntlrProgramBuilder) parseVarAttribute(list []parser.IVarAttributeContext) []VarAttributeI {
	listVarAttr := make([]VarAttributeI, 0)
	for _, vac := range list {
		switch v := vac.(type) {
		case *parser.VarAttributeGeneralContext:
			attr := NewVarAttribute(b.getBTToken(v.Identifier().GetSymbol()))
			listVarAttr = append(listVarAttr, attr)
		case *parser.VarAttributeArrayOrMapContext:
			exp := b.parseExpress(v.Expression())
			attr := NewVarSquareAttribute(exp, b.getBTGToken("[]", exp.GetASTNode().Token.Line))
			listVarAttr = append(listVarAttr, attr)
		case *parser.VarAttributeVirtualContext:
			attr := NewVarVirtualAttribute(b.getBTToken(v.Identifier().GetSymbol()))
			listVarAttr = append(listVarAttr, attr)
		}
	}
	return listVarAttr
}

func (b *AntlrProgramBuilder) getID(ids []antlr.TerminalNode) string {
	sb := estrings.NewStringBuilder()
	for _, n := range ids {
		sb.Append(n.GetSymbol().GetText()).Append(".")
	}
	sb.SetLength(sb.Length() - 1)
	return sb.String()
}

func (b *AntlrProgramBuilder) registerNewVar(vas ASTNodeI) {
	name := vas.GetASTNode().Token.Text
	eToken := b.pbCtx.HasDefined(name)
	if eToken != nil {
		msg := fmt.Sprintf("已经在第%d行定义", eToken.Line)
		panic(NewBeetlException(VAR_ALREADY_DEFINED, &msg).SetToken(vas.GetASTNode().Token))
	}
	b.pbCtx.AddVar(name)
	b.pbCtx.SetVarPosition(name, vas)
}

func (b *AntlrProgramBuilder) parseBlock(statement []parser.IStatementContext, node antlr.ParserRuleContext) StatementI {
	b.pbCtx.EnterBlock()
	var nodes []StatementI
	for _, i := range statement {
		nodes = append(nodes, b.parseStatement(i))
	}
	block := b.gc.CreateBlock(nodes, b.getBTToken(node.GetStart()))
	b.checkGoto(block)
	b.pbCtx.ExitBlock()
	return block
}

func (b *AntlrProgramBuilder) checkGoto(block IGotoI) {
	switch b.pbCtx.current.gotoValue {
	case IgotoNormal:
		return
	case IgotoContinue, IgotoBreak:
		if !b.pbCtx.current.canStopContinueBreakFlag {
			b.pbCtx.current.parent.gotoValue = b.pbCtx.current.gotoValue
		}
		block.SetGoto(true)
		return
	case IgotoReturn:
		if b.pbCtx.current.parent != b.pbCtx.root {
			b.pbCtx.current.parent.gotoValue = IgotoReturn
		} else {
			b.data.HasGoto = true
		}
		block.SetGoto(true)
	}
}

func (b *AntlrProgramBuilder) parseTextOutputSt(ctx *parser.TextOutputStContext) StatementI {
	tsc := ctx.TextStatment().(*parser.TextStatmentContext)
	tvc := tsc.TextVar().(*parser.TextVarContext)
	var format *FormatExpression
	if tvc.COMMA() != nil {
		var formatName string
		var pattern string
		var tokenName string
		line := 0
		tfc := tvc.Textformat().(*parser.TextformatContext)
		node := tfc.StringLiteral()
		if node != nil {
			pattern = b.getStringValue(node.GetText())
			tokenName = pattern
			line = node.GetSymbol().GetLine()
		}
		fnsc := tfc.FunctionNs().(*parser.FunctionNsContext)
		if fnsc != nil {
			listId := fnsc.AllIdentifier()
			formatName = b.getID(listId)
			tokenName = formatName
			line = listId[0].GetSymbol().GetLine()
		}
		format = b.gc.CreateFormat(formatName, pattern, b.getBTGToken(tokenName, line))
	}
	exp := b.parseExpress(tvc.Expression())
	if tsc.LEFT_TOKEN2() == nil {
		return b.gc.CreateTextOutputSt(exp, format)
	}
	return b.gc.CreateTextOutputSt2(exp, format)
}

func (b *AntlrProgramBuilder) parseForSt(ctx *parser.ForStContext) StatementI {
	pbCtx := b.pbCtx
	pbCtx.EnterBlock()
	pbCtx.current.canStopContinueBreakFlag = true

	forContext := ctx.Statement(0)
	var elseContext *parser.StatementContext

	forTypeCtx := ctx.ForControl().(*parser.ForControlContext)
	if forTypeCtx.ForInControl() != nil {
		forCtx := forTypeCtx.ForInControl().(*parser.ForInControlContext)
		exp := b.parseExpress(forCtx.Expression())
		forVar := NewVarDefineNode(b.getBTToken(forCtx.Identifier().GetSymbol()))
		if pbCtx.HasDefined(forVar.Token.Text) != nil {
			msg := fmt.Sprintf("已经在第%d行定义", pbCtx.HasDefined(forVar.Token.Text).Line)
			panic(NewBeetlException(VAR_ALREADY_DEFINED, &msg).SetToken(forVar.Token))
		}
		tt := forCtx.Identifier().GetSymbol()
		loopStatusVar := NewVarDefineNode(b.getBTGToken(tt.GetText(), tt.GetLine()))
		if pbCtx.HasDefined(loopStatusVar.Token.Text) != nil {
			msg := fmt.Sprintf("For循环隐含变量，已经在第%d行定义", pbCtx.HasDefined(loopStatusVar.Token.Text).Line)
			panic(NewBeetlException(VAR_ALREADY_DEFINED, &msg).SetToken(forVar.Token))
		}
		pbCtx.AddVarAndPosition(forVar)
		pbCtx.AddVarAndPosition(loopStatusVar)
		forPart := b.parseStatement(forContext)
		var elseForPart StatementI
		if ctx.Elsefor() != nil {
			elseContext = ctx.Statement(1).(*parser.StatementContext)
			elseForPart = b.parseStatement(elseContext)
		}
		hasSafe := false
		if varRef, ok := exp.(*VarRef); ok {
			hasSafe = varRef.hasSafe
		}
		if pbCtx.isSafeOutput {
			hasSafe = true
		}
		forStatement := b.gc.CreateForIn(forVar, exp, hasSafe, forPart, elseForPart)
		b.checkGoto(forStatement)
		pbCtx.ExitBlock()
		return forStatement
	} else {
		forCtx := forTypeCtx.GeneralForControl().(*parser.GeneralForControlContext)
		var initExp []ExpressionI
		var varInitSeq *VarAssignSeqStatement
		var condition ExpressionI
		var updateExp []ExpressionI
		if forInitCtx0 := forCtx.ForInit(); forInitCtx0 != nil {
			forInitCtx := forInitCtx0.(*parser.ForInitContext)
			if forInitCtx.Var() == nil {
				list := forInitCtx.ExpressionList().(*parser.ExpressionListContext).AllExpression()
				initExp = b.parseExpressionCtxList(list)
			} else {
				varDeclare := forInitCtx.VarDeclareList().(*parser.VarDeclareListContext)
				varInitSeq = b.parseVarDeclareList(varDeclare)
			}
		}
		if forCtx.Expression() != nil {
			condition = b.parseExpress(forCtx.Expression())
		}
		if forCtx.ForUpdate() != nil {
			updateCtx := forCtx.ForUpdate().(*parser.ForUpdateContext)
			list := updateCtx.ExpressionList().(*parser.ExpressionListContext).AllExpression()
			updateExp = b.parseExpressionCtxList(list)
		}
		forPart := b.parseStatement(forContext)
		var elseForPart StatementI
		if ctx.Elsefor() != nil {
			elseContext = ctx.Statement(1).(*parser.StatementContext)
			elseForPart = b.parseStatement(elseContext)
		}
		str := forTypeCtx.GetText()
		forStat := b.gc.CreateFor(varInitSeq, initExp, condition, updateExp, forPart, elseForPart, b.getBTGToken(str, forTypeCtx.GetStart().GetLine()))
		pbCtx.ExitBlock()
		return forStat
	}
}

func (b *AntlrProgramBuilder) parseExpressionCtxList(list []parser.IExpressionContext) []ExpressionI {
	if list == nil || len(list) == 0 {
		return make([]ExpressionI, 0)
	}
	var expList []ExpressionI
	for _, ec := range list {
		expList = append(expList, b.parseExpress(ec))
	}
	return expList
}

func (b *AntlrProgramBuilder) parseIf(ctx *parser.IfStContext) StatementI {
	pe := ctx.ParExpression().(*parser.ParExpressionContext)
	//expCtx := pe.Expression().(*parser.ExpressionContext)
	expCtx := pe.Expression().(*parser.NotExpContext)
	exp := b.parseExpress(expCtx)
	//ifStatCtx := ctx.Statement(0).(*parser.StatementContext)
	ifStatCtx := ctx.Statement(0).(*parser.BlockStContext)
	ifStat := b.parseStatement(ifStatCtx)
	elseStateCtxi := ctx.Statement(1)
	var elseStat StatementI
	if elseStateCtxi != nil {
		elseStateCtx := elseStateCtxi.(*parser.StatementContext)
		elseStat = b.parseStatement(elseStateCtx)
	}
	return b.gc.CreateIf(exp, ifStat, elseStat, b.getBTToken(ctx.If().GetSymbol()))
}

func (b *AntlrProgramBuilder) parseDirectiveStatement(stContext *parser.DirectiveStContext) StatementI {
	direExp := stContext.DirectiveExp().(*parser.DirectiveExpContext)
	token := direExp.Identifier().GetSymbol()
	directive := strings.ToLower(token.GetText())
	isSafeOutput := "safe_output_open" == directive
	b.pbCtx.isSafeOutput = isSafeOutput
	return nil
}

func (b *AntlrProgramBuilder) parseTryCatch(tryStCtx *parser.TryStContext) StatementI {
	tryBlockCtx := tryStCtx.Block(0).(*parser.BlockContext)
	tryPart := b.parseBlock(tryBlockCtx.AllStatement(), tryBlockCtx)
	var catchPart StatementI
	var errorNode *VarDefineNode
	pbCtx := b.pbCtx
	if tryStCtx.Catch() != nil {
		pbCtx.EnterBlock()
		if tryStCtx.Identifier() != nil {
			errorToken := tryStCtx.Identifier().GetSymbol()
			errorNode = NewVarDefineNode(b.getBTToken(errorToken))
			pbCtx.AddVarAndPosition(errorNode)
		}
		catchBlockCtx := tryStCtx.Block(1).(*parser.BlockContext)
		catchPart = b.parseBlock(catchBlockCtx.AllStatement(), catchBlockCtx)
		pbCtx.ExitBlock()
	}
	statement := b.gc.CreateTry(tryPart.(*BlockStatement), catchPart.(*BlockStatement), errorNode, b.getBTToken(tryStCtx.Try().GetSymbol()))
	return statement
}

func (b *AntlrProgramBuilder) parseTag(fc *parser.FunctionTagCallContext) StatementI {
	ids := fc.FunctionNs().(*parser.FunctionNsContext).AllIdentifier()
	id := b.getID(ids)
	expListCtx := fc.ExpressionList()
	var list []parser.IExpressionContext
	if expListCtx == nil {
		list = make([]parser.IExpressionContext, 0)
	} else {
		list = expListCtx.(*parser.ExpressionListContext).AllExpression()
	}
	expList := b.parseExpressionCtxList(list)
	switch id {
	case "htmltagvar", "htmltagexport", "htmltagRootExport":
		panic("unsupported")
	default:
		blockCtx := fc.Block().(*parser.BlockContext)
		block := b.parseBlock(blockCtx.AllStatement(), blockCtx)
		tf := b.gt.GetTagFactory(id)
		token := b.getBTGToken(id, fc.FunctionNs().GetStart().GetLine())
		if tf == nil {
			panic(NewBeetlException(TAG_NOT_FOUND, nil).SetToken(token))
		}
		return b.gc.CreateTag(id, expList, block, token)
	}
}

func (b *AntlrProgramBuilder) parseWhile(wc *parser.WhileStContext) StatementI {
	pbCtx := b.pbCtx
	pbCtx.EnterBlock()
	pbCtx.current.canStopContinueBreakFlag = true
	conditionCtx := wc.ParExpression().(*parser.ParExpressionContext).Expression().(*parser.ExpressionContext)
	bodyCtx := wc.Statement().(*parser.StatementContext)
	condition := b.parseExpress(conditionCtx)
	body := b.parseStatement(bodyCtx)
	whileStat := b.gc.CreateWhile(condition, body, b.getBTToken(wc.GetStart()))
	pbCtx.ExitBlock()
	return whileStat
}

func (b *AntlrProgramBuilder) parseSwitch(sctx *parser.SiwchStContext) StatementI {
	ect := sctx.ParExpression().(*parser.ParExpressionContext).Expression().(*parser.ExpressionContext)
	exp := b.parseExpress(ect)
	list := sctx.SwitchBlock().(*parser.SwitchBlockContext).AllSwitchBlockStatementGroup()
	conditionsStatementsMap := make(map[ExpressionI]*BlockStatement)
	var conditionList []ExpressionI
	var defaultBlock *BlockStatement
	for _, gp := range list {
		group := gp.(*parser.SwitchBlockStatementGroupContext)
		labels := group.AllSwitchLabel()
		stats := group.AllStatement()
		var block *BlockStatement
		if stats == nil {
			block = nil
		} else {
			b.parseBlock(stats, group)
		}
		for _, lb := range labels {
			label := lb.(*parser.SwitchLabelContext)
			caseExp := b.parseExpress(label.Expression())
			if caseExp == nil {
				defaultBlock = block
				break
			} else {
				conditionList = append(conditionList, caseExp)
				conditionsStatementsMap[caseExp] = block
			}
		}
	}
	return b.gc.CreateSwitch(exp, conditionsStatementsMap, defaultBlock, b.getBTToken(sctx.GetStart()))
}

func (b *AntlrProgramBuilder) parseSelect(selectCtx *parser.SelectStContext) StatementI {
	ctx := selectCtx.G_switchStatment().(*parser.G_switchStatmentContext)
	exp := ctx.Expression().(*parser.ExpressionContext)
	var base ExpressionI
	if exp == nil {
		base = nil
	} else {
		base = b.parseExpress(exp)
	}
	caseCtxList := ctx.AllG_caseStatment()
	var conditionList []ExpressionI
	var blockList []*BlockStatement
	for _, cctx := range caseCtxList {
		caseCtx := cctx.(*parser.G_caseStatmentContext)
		expCtxList := caseCtx.AllExpression()
		statCtxList := caseCtx.AllStatement()
		block := b.parseBlock(statCtxList, caseCtx)
		for _, expC := range expCtxList {
			expCtx := expC.(*parser.ExpressionContext)
			condition := b.parseExpress(expCtx)
			conditionList = append(conditionList, condition)
			blockList = append(blockList, block.(*BlockStatement))
		}
	}
	var defaultStatement StatementI
	defaultCtx := ctx.G_defaultStatment().(*parser.G_defaultStatmentContext)
	if defaultCtx != nil {
		defaultCtxList := defaultCtx.AllStatement()
		defaultStatement = b.parseBlock(defaultCtxList, ctx)
	}
	return b.gc.CreateSelect(base, conditionList, blockList, defaultStatement.(*BlockStatement), b.getBTToken(selectCtx.Select().GetSymbol()))
}

func (b *AntlrProgramBuilder) parseAjax(ajaxCtx *parser.AjaxStContext) StatementI {
	var token *GrammarToken
	flag := "render"
	nodes := ajaxCtx.AllIdentifier()
	if len(nodes) == 1 {
		token = b.getBTToken(nodes[0].GetSymbol())
	} else {
		token = b.getBTToken(nodes[1].GetSymbol())
		flag = nodes[0].GetSymbol().GetText()
		if !("render" == flag || "norender" == flag) {
			msg := "expect render or norender ,but " + flag
			panic(NewBeetlException(AJAX_PROPERTY_ERROR, &msg).SetToken(token))
		}
	}
	blockCtx := ajaxCtx.Block().(*parser.BlockContext)
	block := b.parseBlock(blockCtx.AllStatement(), blockCtx)
	ajaxStat := b.gc.CreateAjax(block.(*BlockStatement), token, "render" == flag)
	if b.data.ajaxs == nil {
		b.data.ajaxs = make(map[string]*AjaxStatement)
	}
	anchor := ajaxStat.Token.Text
	if ll, ok := b.data.ajaxs[anchor]; ok {
		lastToken := ll.Token
		msg := fmt.Sprintf("已经在第%d行定义", lastToken.Line)
		panic(NewBeetlException(AJAX_ALREADY_DEFINED, &msg).SetToken(token))
	}
	b.data.ajaxs[anchor] = ajaxStat
	b.reParseAjax(ajaxStat, blockCtx)
	return ajaxStat
}

func (b *AntlrProgramBuilder) reParseAjax(ajaxStat *AjaxStatement, blockCtx *parser.BlockContext) {
	templateCtx := b.pbCtx
	b.pbCtx = NewProgramBuilderContext()
	templateBlock := b.parseBlock(blockCtx.AllStatement(), blockCtx)
	localMetaData := b.buildAjaxBlock(b.pbCtx, templateBlock)
	b.pbCtx = templateCtx
	ajaxStat.SetLocalProgramMetaData(localMetaData)
}

func (b *AntlrProgramBuilder) buildAjaxBlock(pbCtx *ProgramBuilderContext, templateBlock StatementI) *ProgramMetaData {
	local := NewProgramMetaData()
	pbCtx.AnzlyszeGlobal()
	pbCtx.AnzlyszeLocal()
	local.varIndexSize = pbCtx.varIndexSize
	local.tempVarStartIndex = len(pbCtx.globalIndexMap)
	local.statements = []StatementI{templateBlock}
	local.globalIndexMap = pbCtx.globalIndexMap
	local.SetTemplateRootScopeIndexMap(pbCtx.rootIndexMap)
	return local
}

func (b *AntlrProgramBuilder) parseVarRefInLeftExpression(varRef *parser.VarRefContext) *VarRef {
	soctx := varRef.Safe_output().(*parser.Safe_outputContext)
	if soctx != nil {
		msg := "语法错误，赋值表达式不能使用安全输出"
		panic(NewBeetlException(ERROR, &msg))
	}
	list := varRef.AllVarAttribute()
	vas := b.parseVarAttribute(list)
	if len(vas) > 0 {
		first := vas[0]
		switch first.(type) {
		case *VarSquareAttribute, *VarVirtualAttribute:
			break
		default:
			b.pbCtx.SetVarAttr(varRef.Identifier().GetText(), first.GetASTNode().Token.Text)
		}
	}
	varr := NewVarRef(vas, false, nil, b.getBTGToken(varRef.GetText(), varRef.Identifier().GetSymbol().GetLine()), b.getBTToken(varRef.Identifier().GetSymbol()))
	b.pbCtx.SetVarPosition(varRef.Identifier().GetText(), varr)
	return varr
}

func (b *AntlrProgramBuilder) parseCompareExpression(ctx *parser.CompareExpContext) ExpressionI {
	left := b.parseExpress(ctx.Expression(0))
	right := b.parseExpress(ctx.Expression(1))
	tn := ctx.GetChild(1).(antlr.TerminalNode)
	mode := 0
	switch {
	case ctx.EQUAL() != nil:
		mode = 0
	case ctx.NOT_EQUAL() != nil:
		mode = 1
	case ctx.LARGE() != nil:
		mode = 2
	case ctx.LARGE_EQUAL() != nil:
		mode = 3
	case ctx.LESS() != nil:
		mode = 4
	case ctx.LESS_EQUAL() != nil:
		mode = 5
	}
	return b.gc.CreateCompare(left, right, mode, b.getBTToken(tn.GetSymbol()))
}

func (b *AntlrProgramBuilder) parseTernaryExpression(ctx *parser.TernaryExpContext) ExpressionI {
	cond := b.parseExpress(ctx.Expression(0))
	left := b.parseExpress(ctx.Expression(1))
	var right ExpressionI = nil
	if ctx.COLON() != nil {
		right = b.parseExpress(ctx.Expression(2))
	}
	tn := ctx.GetChild(1).(antlr.TerminalNode)
	return b.gc.CreateTernary(cond, left, right, b.getBTToken(tn.GetSymbol()))
}

func (b *AntlrProgramBuilder) parseMuldivmodExpression(ctx *parser.MuldivmodExpContext) ExpressionI {
	left := b.parseExpress(ctx.Expression(0))
	right := b.parseExpress(ctx.Expression(1))
	tn := ctx.GetChild(1).(antlr.TerminalNode)
	mode := 0
	switch {
	case ctx.MUlTIP() != nil:
		mode = MUL
	case ctx.DIV() != nil:
		mode = DIV
	case ctx.MOD() != nil:
		mode = MOD
	}
	return b.gc.CreateArth(left, right, mode, b.getBTToken(tn.GetSymbol()))
}

func (b *AntlrProgramBuilder) parsePlusMins(ctx *parser.AddminExpContext) ExpressionI {
	left := b.parseExpress(ctx.Expression(0))
	right := b.parseExpress(ctx.Expression(1))
	tn := ctx.GetChild(1).(antlr.TerminalNode)
	mode := 0
	switch {
	case ctx.ADD() != nil:
		mode = PLUS
	case ctx.MIN() != nil:
		mode = MIN
	}
	return b.gc.CreateArth(left, right, mode, b.getBTToken(tn.GetSymbol()))
}

func (b *AntlrProgramBuilder) parseJson(ctx *parser.JsonContext) ExpressionI {
	if ctx.LEFT_SQBR() != nil {
		var jsonn *JsonArrayExpression
		listCtx := ctx.AllExpression()
		token := b.getBTToken(ctx.LEFT_SQBR().GetSymbol())
		if len(listCtx) == 0 {
			jsonn = b.gc.CreateJasonArray(make([]ExpressionI, 0), token)
		} else {
			var list []ExpressionI
			for _, expc := range listCtx {
				expCtx := expc.(*parser.ExpressionContext)
				list = append(list, b.parseExpress(expCtx))
			}
			jsonn = b.gc.CreateJasonArray(list, token)
		}
		return jsonn
	} else {
		var jsonn *JsonMapExpression
		listCtx := ctx.AllJsonKeyValue()
		token := b.getBTToken(ctx.LEFT_BRACE().GetSymbol())
		if len(listCtx) == 0 {
			jsonn = b.gc.CreateJsonMap(make(map[string]ExpressionI), token)
		} else {
			m := make(map[string]ExpressionI)
			for _, kv := range listCtx {
				kvCtx := kv.(*parser.JsonKeyValueContext)
				var key string
				if kvCtx.StringLiteral() == nil {
					key = kvCtx.Identifier().GetSymbol().GetText()
				} else {
					key = b.getStringValue(kvCtx.StringLiteral().GetText())
				}
				exp := b.parseExpress(kvCtx.Expression())
				m[key] = exp
			}
			jsonn = b.gc.CreateJsonMap(m, token)
		}
		return jsonn
	}
}

func (b *AntlrProgramBuilder) parseAndExpression(ctx *parser.AndExpContext) ExpressionI {
	left := b.parseExpress(ctx.Expression(0))
	right := b.parseExpress(ctx.Expression(1))
	return b.gc.CreateAnd(left, right, b.getBTToken(ctx.AND().GetSymbol()))
}

func (b *AntlrProgramBuilder) parseOrExpression(ctx *parser.OrExpContext) ExpressionI {
	left := b.parseExpress(ctx.Expression(0))
	right := b.parseExpress(ctx.Expression(1))
	return b.gc.CreateOr(left, right, b.getBTToken(ctx.OR().GetSymbol()))
}

func (b *AntlrProgramBuilder) parseNotExpression(ctx *parser.NotExpContext) ExpressionI {
	exp := b.parseExpress(ctx.Expression())
	return b.gc.CreateNot(exp, b.getBTToken(ctx.NOT().GetSymbol()))
}

func (b *AntlrProgramBuilder) parseNegExpression(ctx *parser.NegExpContext) ExpressionI {
	exp := b.parseExpress(ctx.Expression())
	if ctx.MIN() == nil {
		return exp
	} else {
		return b.gc.CreateNeg(exp, b.getBTToken(ctx.MIN().GetSymbol()))
	}
}

func (b *AntlrProgramBuilder) parseIncDecOneContext(ctx *parser.IncDecOneContext) ExpressionI {
	isInc := ctx.INCREASE() != nil
	t := b.getBTToken(ctx.Identifier().GetSymbol())
	exp := b.gc.CreateIncDec(isInc, false, t)
	b.pbCtx.SetVarPosition(t.Text, exp)
	return exp
}

func (b *AntlrProgramBuilder) parseOneIncDecContext(ctx *parser.OneIncDecContext) ExpressionI {
	isInc := ctx.INCREASE() != nil
	t := b.getBTToken(ctx.Identifier().GetSymbol())
	exp := b.gc.CreateIncDec(isInc, true, t)
	b.pbCtx.SetVarPosition(t.Text, exp)
	return exp
}

func (b *AntlrProgramBuilder) parseAssingInExp(agc *parser.AssignGeneralInExpContext) ExpressionI {
	agcc := agc.GeneralAssignExp().(*parser.GeneralAssignExpContext)
	expCtx := agcc.Expression()
	exp := b.parseExpress(expCtx)
	varRefCtx := agcc.VarRef().(*parser.VarRefContext)
	ref := b.parseVarRefInLeftExpression(varRefCtx)
	vas := b.gc.CreateVarRefAssignExp(exp, ref)
	if len(ref.attributes) == 0 {
		token := varRefCtx.Identifier().GetSymbol()
		if b.pbCtx.HasDefined(token.GetText()) == nil {
			panic(NewBeetlException(VAR_NOT_DEFINED, nil).SetToken(b.getBTToken(token)))
		}
		b.registerVar(vas)
	}
	return vas
}

func (b *AntlrProgramBuilder) registerVar(vas ASTNodeI) {
	b.pbCtx.SetVarPosition(vas.GetASTNode().Token.Text, vas)
}

func (b *AntlrProgramBuilder) parseSelfAssingInExp(selfExpContext *parser.AssingSelfExpContext) ExpressionI {
	selfExp := selfExpContext.GeneralAssingSelfExp().(*parser.GeneralAssingSelfExpContext)
	varRefCtx := selfExp.VarRef().(*parser.VarRefContext)
	ref := b.parseVarRefInLeftExpression(varRefCtx)
	exp := b.parseExpress(selfExp.Expression())
	tn := selfExp.GetChild(1).(antlr.TerminalNode)
	mode := 0
	switch {
	case selfExp.MUlTIP() != nil:
		mode = MUL
	case selfExp.DIV() != nil:
		mode = DIV
	case selfExp.ADD() != nil:
		mode = PLUS
	case selfExp.MIN() != nil:
		mode = MIN
	}
	arthExpress := b.gc.CreateArth(ref, exp, mode, b.getBTToken(tn.GetSymbol()))
	vas := b.gc.CreateVarRefAssignExp(arthExpress, ref)
	if len(ref.attributes) == 0 {
		token := varRefCtx.Identifier().GetSymbol()
		if b.pbCtx.HasDefined(token.GetText()) == nil {
			panic(NewBeetlException(VAR_NOT_DEFINED, nil).SetToken(b.getBTToken(token)))
		}
		b.registerVar(vas)
	}
	return vas
}
