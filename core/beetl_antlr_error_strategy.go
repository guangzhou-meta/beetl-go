package core

import (
	"fmt"
	"reflect"
)

import (
	"github.com/guangzhou-meta/beetl-go/util"
)

import (
	"github.com/guangzhou-meta/antlr4/runtime/Go/antlr"
)

var (
	keys = map[string]interface{}{
		"select":    nil,
		"if":        nil,
		"for":       nil,
		"elsefor":   nil,
		"while":     nil,
		"switch":    nil,
		"return":    nil,
		"break":     nil,
		"var":       nil,
		"continue":  nil,
		"directive": nil,
		"in":        nil,
		"case":      nil,
		"default":   nil,
		"try":       nil,
		"catch":     nil,
	}
	expects = map[string]string{
		"LEFT_PAR":  "(",
		"RIGHT_PAR": "]=nil",
	}
)

type BeetlAntlrErrorStrategy struct {
	*DefaultErrorStrategy
}

func NewBeetlAntlrErrorStrategy() *BeetlAntlrErrorStrategy {
	inst := &BeetlAntlrErrorStrategy{
		DefaultErrorStrategy: NewDefaultErrorStrategy(),
	}
	return inst
}

func (s *BeetlAntlrErrorStrategy) ReportError(recognizer antlr.Parser, e antlr.RecognitionException) {
	if s.InErrorRecoveryMode(recognizer) {
		return
	}
	s.beginErrorCondition(recognizer)
	switch e0 := e.(type) {
	case *antlr.NoViableAltException:
		s.reportNoViableAlternative(recognizer, e0)
	case *antlr.InputMisMatchException:
		s.reportInputMismatch(recognizer, e0)
	case *antlr.FailedPredicateException:
		s.reportFailedPredicate(recognizer, e0)
	default:
		emptyMsg := PARSER_UNKNOW_ERROR
		panic(NewBeetlParserException(PARSER_UNKNOW_ERROR, &emptyMsg).SetToken(getGrammarToken(e.GetOffendingToken())))
	}
}

func (s *BeetlAntlrErrorStrategy) beginErrorCondition(recognizer antlr.Parser) {
	s.errorRecoveryMode = true
}

func (s *BeetlAntlrErrorStrategy) reportNoViableAlternative(recognizer antlr.Parser, e *antlr.NoViableAltException) {
	tokens := recognizer.GetTokenStream()
	var input string
	tV := util.UnpackValue(reflect.ValueOf(e))
	startToken := tV.FieldByName("startToken").Interface().(antlr.Token)
	if tokens == nil {
		input = "<未知输入>"
	} else if antlr.TokenEOF == startToken.GetTokenType() {
		input = "<文件尾>"
	} else {
		input = tokens.GetTextFromTokens(startToken, e.GetOffendingToken())
	}
	var msg string
	if util.Contains(keys, e.GetOffendingToken().GetText()) {
		msg = fmt.Sprintf("不允许%s关键出现在这里:%s", e.GetOffendingToken().GetText(), s.escapeWSAndQuote(input))
	} else {
		msg = s.escapeWSAndQuote(input)
	}
	panic(NewBeetlParserException(PARSER_VIABLE_ERROR, &msg).SetToken(getGrammarToken(e.GetOffendingToken())))
}

func (s *BeetlAntlrErrorStrategy) reportInputMismatch(recognizer antlr.Parser, e *antlr.InputMisMatchException) {
	t1 := recognizer.GetTokenStream().LT(-1)
	msg := fmt.Sprintf("缺少输入在 %s 后面， 期望 %s", s.GetTokenErrorDisplay(t1))
	recognizer.NotifyErrorListeners(msg, e.GetOffendingToken(), e)
}

func (s *BeetlAntlrErrorStrategy) GetTokenErrorDisplay(t antlr.Token) string {
	if t == nil {
		return "<no token>"
	}
	str := t.GetText()
	if str == "" {
		if t.GetTokenType() == -1 {
			str = "<EOF>"
		} else {
			str = fmt.Sprintf("<%d>", t.GetTokenType())
		}
	}
	return s.escapeWSAndQuote(str)
}

func (s *BeetlAntlrErrorStrategy) reportFailedPredicate(recognizer antlr.Parser, e *antlr.FailedPredicateException) {
	ruleName := recognizer.GetRuleNames()[recognizer.GetParserRuleContext().GetRuleIndex()]
	msg := fmt.Sprintf("rule %s %s", ruleName, e.GetMessage())
	recognizer.NotifyErrorListeners(msg, e.GetOffendingToken(), e)
}

func getGrammarToken(token antlr.Token) *GrammarToken {
	return NewGrammarToken(token.GetText(), token.GetLine(), 0)
}

func (s *BeetlAntlrErrorStrategy) Recover(recognizer antlr.Parser, e antlr.RecognitionException) {
	s.DefaultErrorStrategy.Recover(recognizer, e)
}

func (d *BeetlAntlrErrorStrategy) RecoverInline(recognizer antlr.Parser) antlr.Token {
	matchedSymbol := d.SingleTokenDeletion(recognizer)
	if matchedSymbol != nil {
		recognizer.Consume()
		return matchedSymbol
	}

	if d.SingleTokenInsertion(recognizer) {
		return d.GetMissingSymbol(recognizer)
	}

	panic(fmt.Sprintf("input mismatch: %v", recognizer))
}

func (d *BeetlAntlrErrorStrategy) ReportMissingToken(recognizer antlr.Parser) {
	if d.InErrorRecoveryMode(recognizer) {
		return
	}

	d.beginErrorCondition(recognizer)

	t := recognizer.GetTokenStream().LT(-1)
	expecting := d.GetExpectedTokens(recognizer)
	expect := expecting.StringVerbose(recognizer.GetLiteralNames(), recognizer.GetSymbolicNames(), false)
	if util.Contains(expects, expect) {
		expect = expects[expect]
	}
	if expect == "'>>'" {
		expect = "'模板的占位结束符号'"
	}
	tokenStr := d.GetTokenErrorDisplay(t)
	var msg string
	if expect == "'}'" && tokenStr == "'>>'" {
		msg = "未找到 '{' 匹配的结束符号 '}'"
	} else {
		msg = "缺少输入 " + expect + " 在 " + tokenStr + " 后面"
	}
	panic(NewBeetlParserException(PARSER_MISS_ERROR, &msg).SetToken(getGrammarToken(t)))
}

func (d *BeetlAntlrErrorStrategy) ReportUnwantedToken(recognizer antlr.Parser) {
	if d.InErrorRecoveryMode(recognizer) {
		return
	}

	d.beginErrorCondition(recognizer)

	t := recognizer.GetCurrentToken()
	tokenName := d.GetTokenErrorDisplay(t)
	expecting := d.GetExpectedTokens(recognizer)
	key := d.getErrorKey(expecting.StringVerbose(recognizer.GetLiteralNames(), recognizer.GetSymbolicNames(), false))
	msg := "多余输入 " + tokenName + " 期望 " + key
	panic(NewBeetlParserException(PARSER_MISS_ERROR, &msg).SetToken(getGrammarToken(t)))
}

func (s *BeetlAntlrErrorStrategy) getErrorKey(key string) string {
	if key == "'>>'" {
		return "占位结束符号"
	}
	return key
}
