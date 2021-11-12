package core

import (
	"reflect"
)

import (
	"github.com/guangzhou-meta/beetl-go/util"
)

type FormatExpression struct {
	*Expression
	name    *string
	pattern *string
}

func NewFormatExpression(name *string, pattern *string, token *GrammarToken) *FormatExpression {
	return &FormatExpression{
		Expression: NewExpression(token),
		name:       name,
		pattern:    pattern,
	}
}

func (fe *FormatExpression) Evaluate(ctx *Context) interface{} {
	panic("unsupported")
}

func (fe *FormatExpression) EvaluateValue(o interface{}, ctx *Context) interface{} {
	var format FormatI
	if fe.name != nil {
		format = ctx.gt.GetFormat(*fe.name)
	} else {
		if o == nil {
			return nil
		}
		format = ctx.gt.GetDefaultFormat(util.UnpackType(reflect.TypeOf(o)))
	}

	if format == nil {
		panic(NewBeetlException(FORMAT_NOT_FOUND, nil).SetToken(fe.Token))
	}

	if f, ok := format.(ContextFormatI); ok {
		return f.Format2(o, *fe.pattern, ctx)
	}
	return format.Format(o, *fe.pattern)
}
