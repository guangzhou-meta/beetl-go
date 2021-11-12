package core

import (
	"fmt"
)

var (
	placeholderSTOutput PlaceholderSTOutput
)

type PlaceholderSTI interface {
	StatementI

	GetFormat() *FormatExpression
}

type PlaceholderST struct {
	*Statement

	expression ExpressionI
	format     *FormatExpression
}

func NewPlaceholderST(exp ExpressionI, format *FormatExpression, token *GrammarToken) *PlaceholderST {
	return &PlaceholderST{
		Statement:  NewStatement(token),
		format:     format,
		expression: exp,
	}
}

func (p *PlaceholderST) Execute(ctx *Context) {
	value := p.expression.Evaluate(ctx)
	if p.format != nil {
		value = p.format.EvaluateValue(value, ctx)
	}
	if placeholderSTOutput != nil {
		placeholderSTOutput.Write(ctx, value)
		return
	}
	if value != nil {
		if s, ok := value.(string); ok {
			ctx.byteWriter.WriteString("'" + s + "'")
			//ctx.byteWriter.Append("'" + s + "'")
			return
		} else {
			switch v := value.(type) {
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				ctx.byteWriter.WriteString(fmt.Sprintf("%d", v))
				//ctx.byteWriter.Append(fmt.Sprintf("%d", v))
				return
			case float32, float64:
				ctx.byteWriter.WriteString(fmt.Sprintf("%f", v))
				//ctx.byteWriter.Append(fmt.Sprintf("%f", v))
				return
			}
		}
		ctx.byteWriter.WriteString(fmt.Sprintf("%s", value))
		//ctx.byteWriter.Append(fmt.Sprintf("%s", value))
	}
}

func (p *PlaceholderST) GetFormat() *FormatExpression {
	return p.format
}

type PlaceholderSTOutput interface {
	Write(ctx *Context, value interface{})
}

func SetPlaceholderSTOutput(writer PlaceholderSTOutput) {
	placeholderSTOutput = writer
}
