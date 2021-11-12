package core

import (
	"github.com/guangzhou-meta/beetl-go/util"
)

import (
	"github.com/guangzhou-meta/antlr4/runtime/Go/antlr"
)

type SyntaxErrorListener struct {
	*DefaultErrorListener
}

func NewSyntaxErrorListener() *SyntaxErrorListener {
	inst := &SyntaxErrorListener{
		DefaultErrorListener: NewDefaultErrorListener(),
	}
	return inst
}

func (l *SyntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	panic(NewBeetlException(TOKEN_ERROR, nil).SetToken(NewGrammarToken(util.ReportChineseTokenError(msg), line, column)))
}
