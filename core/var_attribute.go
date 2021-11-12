package core

import (
	"reflect"
)

import (
	"github.com/guangzhou-meta/beetl-go/core/om"
	"github.com/guangzhou-meta/beetl-go/util"
)

type VarAttributeI interface {
	ExpressionI

	Evaluate2(ctx *Context, o interface{}) interface{}

	GetName() string
}

type VarAttribute struct {
	*Expression
	aaIndex int
	name    string
}

func NewVarAttribute(token *GrammarToken) *VarAttribute {
	return &VarAttribute{
		Expression: NewExpression(token),
		name:       token.Text,
		aaIndex:    -1,
	}
}

func (va *VarAttribute) GetName() string {
	return va.name
}

func (va *VarAttribute) Evaluate(ctx *Context) interface{} {
	panic("unsupported")
}

func (va *VarAttribute) Evaluate2(ctx *Context, o interface{}) interface{} {
	aa := om.BuildFiledAccessor(util.UnpackType(reflect.TypeOf(o)).Kind())
	return aa.Value(o, va.name)
}

type VarSquareAttribute struct {
	*VarAttribute
	exp ExpressionI
}

func NewVarSquareAttribute(exp ExpressionI, token *GrammarToken) *VarSquareAttribute {
	return &VarSquareAttribute{
		VarAttribute: NewVarAttribute(token),
		exp:          exp,
	}
}

func (vsa *VarSquareAttribute) Evaluate(ctx *Context) interface{} {
	panic("unsupported")
}

func (vsa *VarSquareAttribute) Evaluate2(ctx *Context, o interface{}) interface{} {
	value := vsa.exp.Evaluate(ctx)
	aa := om.BuildFiledAccessor(util.UnpackType(reflect.TypeOf(o)).Kind())
	return aa.Value(o, value)
}

type VarVirtualAttribute struct {
	*VarAttribute
	name string
}

func NewVarVirtualAttribute(token *GrammarToken) *VarVirtualAttribute {
	return &VarVirtualAttribute{
		VarAttribute: NewVarAttribute(token),
		name:         token.Text,
	}
}

func (vva *VarVirtualAttribute) Evaluate(ctx *Context) interface{} {
	panic("unsupported")
}

func (vva *VarVirtualAttribute) Evaluate2(ctx *Context, o interface{}) interface{} {
	vae := ctx.gt.GetVirtualAttributeEval(util.UnpackType(reflect.TypeOf(o)), vva.name)
	if vae == nil {
		panic(NewBeetlException(VIRTUAL_NOT_FOUND, nil).SetToken(vva.Token))
	}
	return vae.Eval(o, vva.name, ctx)
}
