package core

import (
	"strings"
)

type BeetlSQLTemplateEngine struct {
	*DefaultTemplateEngine
}

func NewBeetlSQLTemplateEngine() *BeetlSQLTemplateEngine {
	inst := &BeetlSQLTemplateEngine{
		DefaultTemplateEngine: NewDefaultTemplateEngine(),
	}
	inst.inject(inst)
	return inst
}

func (e *BeetlSQLTemplateEngine) getAntlrBuilder(groupTemplate *GroupTemplate) *AntlrProgramBuilder {
	return NewAntlrProgramBuilder(groupTemplate, e.getGrammarCreator(groupTemplate))
}

func (e *BeetlSQLTemplateEngine) getGrammarCreator(groupTemplate *GroupTemplate) GrammarCreatorI {
	return NewSQLGrammarCreator()
}

type SQLGrammarCreator struct {
	*GrammarCreator
}

func NewSQLGrammarCreator() *SQLGrammarCreator {
	return &SQLGrammarCreator{
		GrammarCreator: NewGrammarCreator(),
	}
}

func (gc *SQLGrammarCreator) CreateTextOutputSt(exp ExpressionI, format *FormatExpression) PlaceholderSTI {
	gc.disableSyntaxCheck("TextOutputSt")
	return NewSQLPlaceholderST(exp, format, nil)
}

func (gc *SQLGrammarCreator) CreateTextOutputSt2(exp ExpressionI, format *FormatExpression) *PlaceholderST {
	return NewPlaceholderST(exp, format, nil)
}

func (gc *SQLGrammarCreator) CreateFunction(name string, exps []ExpressionI, vas []VarAttributeI, hasSafe bool, safeExp ExpressionI, token *GrammarToken) FunctionExpressionI {
	gc.disableSyntaxCheck("Function")
	return NewSqlFunctionExpression(name, exps, vas, hasSafe, safeExp, token)
}

func (gc *SQLGrammarCreator) CreateFunctionExp(name string, exps []ExpressionI, vas []VarAttributeI, hasSafe bool, safeExp ExpressionI, token *GrammarToken) FunctionExpressionI {
	gc.disableSyntaxCheck("FunctionExp")
	return NewSqlFunctionExpression(name, exps, vas, hasSafe, safeExp, token)
}

type SqlFunctionExpression struct {
	*FunctionExpression
}

func NewSqlFunctionExpression(name string, exps []ExpressionI, vas []VarAttributeI, hasSafe bool, safeExp ExpressionI, token *GrammarToken) *SqlFunctionExpression {
	return &SqlFunctionExpression{
		FunctionExpression: NewFunctionExpression(name, exps, vas, hasSafe, safeExp, token),
	}
}

var (
	textFunList = map[string]interface{}{
		"text":      nil,
		"use":       nil,
		"globalUse": nil,
		"join":      nil,
		"page":      nil,
	}
)

type SQLPlaceholderST struct {
	*PlaceholderST
}

func NewSQLPlaceholderST(exp ExpressionI, format *FormatExpression, token *GrammarToken) *SQLPlaceholderST {
	return &SQLPlaceholderST{
		PlaceholderST: NewPlaceholderST(exp, format, token),
	}
}

func (st *SQLPlaceholderST) Execute(ctx *Context) {
	value := st.expression.Evaluate(ctx)

	if fe, ok := st.expression.(FunctionExpressionI); ok {
		funName := fe.GetToken().Text
		if strings.HasPrefix(funName, "db") {
			if value != nil {
				ctx.byteWriter.WriteString(value.(string))
				//ctx.byteWriter.Append(value.(string))
			}
			return
		} else if _, ok = textFunList[funName]; ok {
			if value != nil {
				ctx.byteWriter.WriteString(value.(string))
				//ctx.byteWriter.Append(value.(string))
			}
			return
		}
	}

	st.PlaceholderST.Execute(ctx)
}
