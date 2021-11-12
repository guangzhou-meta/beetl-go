package core

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

import (
	"github.com/guangzhou-meta/antlr4/runtime/Go/antlr"
)

type DefaultErrorStrategy struct {
	errorRecoveryMode bool
	lastErrorIndex    int
	lastErrorStates   *antlr.IntervalSet
}

func NewDefaultErrorStrategy() *DefaultErrorStrategy {

	d := new(DefaultErrorStrategy)

	// Indicates whether the error strategy is currently "recovering from an
	// error". This is used to suppress Reporting multiple error messages while
	// attempting to recover from a detected syntax error.
	//
	// @see //inErrorRecoveryMode
	//
	d.errorRecoveryMode = false

	// The index into the input stream where the last error occurred.
	// This is used to prevent infinite loops where an error is found
	// but no token is consumed during recovery...another error is found,
	// ad nauseum. This is a failsafe mechanism to guarantee that at least
	// one token/tree node is consumed for two errors.
	//
	d.lastErrorIndex = -1
	d.lastErrorStates = nil
	return d
}

func (d *DefaultErrorStrategy) Reset(recognizer antlr.Parser) {
	d.endErrorCondition(recognizer)
}

func (d *DefaultErrorStrategy) beginErrorCondition(recognizer antlr.Parser) {
	d.errorRecoveryMode = true
}

func (d *DefaultErrorStrategy) InErrorRecoveryMode(recognizer antlr.Parser) bool {
	return d.errorRecoveryMode
}

func (d *DefaultErrorStrategy) endErrorCondition(recognizer antlr.Parser) {
	d.errorRecoveryMode = false
	d.lastErrorStates = nil
	d.lastErrorIndex = -1
}

func (d *DefaultErrorStrategy) ReportMatch(recognizer antlr.Parser) {
	d.endErrorCondition(recognizer)
}

func (d *DefaultErrorStrategy) ReportError(recognizer antlr.Parser, e antlr.RecognitionException) {
	// if we've already Reported an error and have not Matched a token
	// yet successfully, don't Report any errors.
	if d.InErrorRecoveryMode(recognizer) {
		return // don't Report spurious errors
	}
	d.beginErrorCondition(recognizer)

	switch t := e.(type) {
	default:
		fmt.Println("unknown recognition error type: " + reflect.TypeOf(e).Name())
		//            fmt.Println(e.stack)
		recognizer.NotifyErrorListeners(e.GetMessage(), e.GetOffendingToken(), e)
	case *antlr.NoViableAltException:
		d.ReportNoViableAlternative(recognizer, t)
	case *antlr.InputMisMatchException:
		d.ReportInputMisMatch(recognizer, t)
	case *antlr.FailedPredicateException:
		d.ReportFailedPredicate(recognizer, t)
	}
}

func (d *DefaultErrorStrategy) Recover(recognizer antlr.Parser, e antlr.RecognitionException) {
	if d.lastErrorIndex == recognizer.GetInputStream().Index() &&
		d.lastErrorStates != nil && d.lastErrorStates.Contains(recognizer.GetState()) {
		// uh oh, another error at same token index and previously-Visited
		// state in ATN must be a case where LT(1) is in the recovery
		// token set so nothing got consumed. Consume a single token
		// at least to prevent an infinite loop d is a failsafe.
		recognizer.Consume()
	}
	d.lastErrorIndex = recognizer.GetInputStream().Index()
	if d.lastErrorStates == nil {
		d.lastErrorStates = antlr.NewIntervalSet()
	}
	d.lastErrorStates.AddOne(recognizer.GetState())
	followSet := d.getErrorRecoverySet(recognizer)
	d.consumeUntil(recognizer, followSet)
}

func (d *DefaultErrorStrategy) Sync(recognizer antlr.Parser) {
	// If already recovering, don't try to Sync
	if d.InErrorRecoveryMode(recognizer) {
		return
	}

	atn := recognizer.GetInterpreter().ATN()
	states := atn.States
	s := states[recognizer.GetState()]
	la := recognizer.GetTokenStream().LA(1)

	// try cheaper subset first might get lucky. seems to shave a wee bit off
	nextTokens := recognizer.GetATN().NextTokens(s, nil)
	if nextTokens.Contains(antlr.TokenEpsilon) || nextTokens.Contains(la) {
		return
	}

	switch s.GetStateType() {
	case antlr.ATNStateBlockStart, antlr.ATNStateStarBlockStart, antlr.ATNStatePlusBlockStart, antlr.ATNStateStarLoopEntry:
		// Report error and recover if possible
		if d.SingleTokenDeletion(recognizer) != nil {
			return
		}
		panic(antlr.NewInputMisMatchException(recognizer))
	case antlr.ATNStatePlusLoopBack, antlr.ATNStateStarLoopBack:
		d.ReportUnwantedToken(recognizer)
		expecting := antlr.NewIntervalSet()
		expecting.AddSet(recognizer.GetExpectedTokens())
		whatFollowsLoopIterationOrRule := expecting.AddSet(d.getErrorRecoverySet(recognizer))
		d.consumeUntil(recognizer, whatFollowsLoopIterationOrRule)
	default:
		// do nothing if we can't identify the exact kind of ATN state
	}
}

func (d *DefaultErrorStrategy) ReportNoViableAlternative(recognizer antlr.Parser, e *antlr.NoViableAltException) {
	tokens := recognizer.GetTokenStream()
	var input string
	eV := reflect.ValueOf(e).Elem()
	startToken := eV.FieldByName("startToken").Interface().(antlr.Token)
	if tokens != nil {
		if startToken.GetTokenType() == antlr.TokenEOF {
			input = "<EOF>"
		} else {
			input = tokens.GetTextFromTokens(startToken, e.GetOffendingToken())
		}
	} else {
		input = "<unknown input>"
	}
	msg := "no viable alternative at input " + d.escapeWSAndQuote(input)
	recognizer.NotifyErrorListeners(msg, e.GetOffendingToken(), e)
}

func (this *DefaultErrorStrategy) ReportInputMisMatch(recognizer antlr.Parser, e *antlr.InputMisMatchException) {
	msg := "mismatched input " + this.GetTokenErrorDisplay(e.GetOffendingToken()) +
		" expecting " + e.GetExpectedTokens().StringVerbose(recognizer.GetLiteralNames(), recognizer.GetSymbolicNames(), false)
	recognizer.NotifyErrorListeners(msg, e.GetOffendingToken(), e)
}

func (d *DefaultErrorStrategy) ReportFailedPredicate(recognizer antlr.Parser, e *antlr.FailedPredicateException) {
	ruleName := recognizer.GetRuleNames()[recognizer.GetParserRuleContext().GetRuleIndex()]
	msg := "rule " + ruleName + " " + e.GetMessage()
	recognizer.NotifyErrorListeners(msg, e.GetOffendingToken(), e)
}

func (d *DefaultErrorStrategy) ReportUnwantedToken(recognizer antlr.Parser) {
	if d.InErrorRecoveryMode(recognizer) {
		return
	}
	d.beginErrorCondition(recognizer)
	t := recognizer.GetCurrentToken()
	tokenName := d.GetTokenErrorDisplay(t)
	expecting := d.GetExpectedTokens(recognizer)
	msg := "extraneous input " + tokenName + " expecting " +
		expecting.StringVerbose(recognizer.GetLiteralNames(), recognizer.GetSymbolicNames(), false)
	recognizer.NotifyErrorListeners(msg, t, nil)
}

func (d *DefaultErrorStrategy) ReportMissingToken(recognizer antlr.Parser) {
	if d.InErrorRecoveryMode(recognizer) {
		return
	}
	d.beginErrorCondition(recognizer)
	t := recognizer.GetCurrentToken()
	expecting := d.GetExpectedTokens(recognizer)
	msg := "missing " + expecting.StringVerbose(recognizer.GetLiteralNames(), recognizer.GetSymbolicNames(), false) +
		" at " + d.GetTokenErrorDisplay(t)
	recognizer.NotifyErrorListeners(msg, t, nil)
}

func (d *DefaultErrorStrategy) RecoverInline(recognizer antlr.Parser) antlr.Token {
	// SINGLE TOKEN DELETION
	MatchedSymbol := d.SingleTokenDeletion(recognizer)
	if MatchedSymbol != nil {
		// we have deleted the extra token.
		// now, move past ttype token as if all were ok
		recognizer.Consume()
		return MatchedSymbol
	}
	// SINGLE TOKEN INSERTION
	if d.SingleTokenInsertion(recognizer) {
		return d.GetMissingSymbol(recognizer)
	}
	// even that didn't work must panic the exception
	panic(antlr.NewInputMisMatchException(recognizer))
}

func (d *DefaultErrorStrategy) SingleTokenInsertion(recognizer antlr.Parser) bool {
	currentSymbolType := recognizer.GetTokenStream().LA(1)
	// if current token is consistent with what could come after current
	// ATN state, then we know we're missing a token error recovery
	// is free to conjure up and insert the missing token
	atn := recognizer.GetInterpreter().ATN()
	states := reflect.ValueOf(atn).Elem().FieldByName("states").Interface().([]antlr.ATNState)
	currentState := states[recognizer.GetState()]
	next := currentState.GetTransitions()[0].GetTarget()
	expectingAtLL2 := atn.NextTokens(next, recognizer.GetParserRuleContext())
	if expectingAtLL2.Contains(currentSymbolType) {
		d.ReportMissingToken(recognizer)
		return true
	}

	return false
}

func (d *DefaultErrorStrategy) SingleTokenDeletion(recognizer antlr.Parser) antlr.Token {
	NextTokenType := recognizer.GetTokenStream().LA(2)
	expecting := d.GetExpectedTokens(recognizer)
	if expecting.Contains(NextTokenType) {
		d.ReportUnwantedToken(recognizer)
		// print("recoverFromMisMatchedToken deleting " \
		// + str(recognizer.GetTokenStream().LT(1)) \
		// + " since " + str(recognizer.GetTokenStream().LT(2)) \
		// + " is what we want", file=sys.stderr)
		recognizer.Consume() // simply delete extra token
		// we want to return the token we're actually Matching
		MatchedSymbol := recognizer.GetCurrentToken()
		d.ReportMatch(recognizer) // we know current token is correct
		return MatchedSymbol
	}

	return nil
}

func (d *DefaultErrorStrategy) GetMissingSymbol(recognizer antlr.Parser) antlr.Token {
	currentSymbol := recognizer.GetCurrentToken()
	expecting := d.GetExpectedTokens(recognizer)
	expectedTokenType := expecting.First()
	var tokenText string

	if expectedTokenType == antlr.TokenEOF {
		tokenText = "<missing EOF>"
	} else {
		ln := recognizer.GetLiteralNames()
		if expectedTokenType > 0 && expectedTokenType < len(ln) {
			tokenText = "<missing " + recognizer.GetLiteralNames()[expectedTokenType] + ">"
		} else {
			tokenText = "<missing undefined>" // TODO matches the JS impl
		}
	}
	current := currentSymbol
	lookback := recognizer.GetTokenStream().LT(-1)
	if current.GetTokenType() == antlr.TokenEOF && lookback != nil {
		current = lookback
	}

	tf := recognizer.GetTokenFactory()

	return tf.Create(current.GetSource(), expectedTokenType, tokenText, antlr.TokenDefaultChannel, -1, -1, current.GetLine(), current.GetColumn())
}

func (d *DefaultErrorStrategy) GetExpectedTokens(recognizer antlr.Parser) *antlr.IntervalSet {
	return recognizer.GetExpectedTokens()
}

func (d *DefaultErrorStrategy) GetTokenErrorDisplay(t antlr.Token) string {
	if t == nil {
		return "<no token>"
	}
	s := t.GetText()
	if s == "" {
		if t.GetTokenType() == antlr.TokenEOF {
			s = "<EOF>"
		} else {
			s = "<" + strconv.Itoa(t.GetTokenType()) + ">"
		}
	}
	return d.escapeWSAndQuote(s)
}

func (d *DefaultErrorStrategy) escapeWSAndQuote(s string) string {
	s = strings.Replace(s, "\t", "\\t", -1)
	s = strings.Replace(s, "\n", "\\n", -1)
	s = strings.Replace(s, "\r", "\\r", -1)
	return "'" + s + "'"
}

func (d *DefaultErrorStrategy) getErrorRecoverySet(recognizer antlr.Parser) *antlr.IntervalSet {
	atn := recognizer.GetInterpreter().ATN()
	ctx := recognizer.GetParserRuleContext()
	recoverSet := antlr.NewIntervalSet()
	for ctx != nil && ctx.GetInvokingState() >= 0 {
		// compute what follows who invoked us
		states := reflect.ValueOf(atn).Elem().FieldByName("states").Interface().([]antlr.ATNState)
		invokingState := states[ctx.GetInvokingState()]
		rt := invokingState.GetTransitions()[0]
		follow := atn.NextTokens(reflect.ValueOf(rt).Elem().FieldByName("followState").Interface().(antlr.ATNState), nil)
		recoverSet.AddSet(follow)
		ctx = ctx.GetParent().(antlr.ParserRuleContext)
	}
	recoverSet.RemoveOne(antlr.TokenEpsilon)
	return recoverSet
}

func (d *DefaultErrorStrategy) consumeUntil(recognizer antlr.Parser, set *antlr.IntervalSet) {
	ttype := recognizer.GetTokenStream().LA(1)
	for ttype != antlr.TokenEOF && !set.Contains(ttype) {
		recognizer.Consume()
		ttype = recognizer.GetTokenStream().LA(1)
	}
}
