package core

import (
	"reflect"
)

import (
	"github.com/guangzhou-meta/beetl-go/core/om"
	"github.com/guangzhou-meta/beetl-go/util"
)

type VarRef struct {
	*Expression

	attributes []VarAttributeI
	safe       ExpressionI
	varIndex   int
	hasSafe    bool
	firstToken *GrammarToken
}

func NewVarRef(attributes []VarAttributeI, hasSafe bool, safe ExpressionI, token *GrammarToken, firstToken *GrammarToken) *VarRef {
	inst := &VarRef{
		Expression: NewExpression(token),
		attributes: attributes,
		safe:       safe,
		hasSafe:    hasSafe,
		firstToken: firstToken,
	}
	return inst
}

func (vf *VarRef) Evaluate(ctx *Context) interface{} {
	ret := vf.getValue(ctx)
	if ret.Safe {
		return ret.Value
	}

	value := ret.Value

	if len(vf.attributes) == 0 {
		return value
	}

	for i := range vf.attributes {
		attr := vf.attributes[i]
		if value == nil {
			if vf.hasSafe || ctx.safeOutput {
				if vf.safe == nil {
					return nil
				}
				return vf.safe.Evaluate(ctx)
			}
		}
		value = attr.Evaluate2(ctx, value)
	}

	if value == nil && (vf.hasSafe || ctx.safeOutput) {
		if vf.safe == nil {
			return nil
		}
		return vf.safe.Evaluate(ctx)
	}
	return value
}

func (vf *VarRef) getValue(ctx *Context) *VarRefResult {
	value := ctx.Vars[vf.varIndex]
	if NotExistObject == value {
		if _, ok := ctx.GlobalVar["_root"]; ok {
			root := ctx.GetGlobal("_root")
			attr := vf.firstToken.Text
			if root == nil {
				if vf.hasSafe || ctx.safeOutput {
					return vf.returnSafeValue(ctx)
				}
				msg := "_root为空指针，无" + attr + "值"
				panic(NewBeetlException(NULL, &msg).SetToken(vf.firstToken))
			}
			aa := om.BuildFiledAccessor(util.UnpackType(reflect.TypeOf(root)).Kind())
			value = aa.Value(root, attr)
			ctx.Vars[vf.varIndex] = value
		} else if vf.hasSafe || ctx.safeOutput {
			return vf.returnSafeValue(ctx)
		} else {
			panic(NewBeetlException(VAR_NOT_DEFINED, nil).SetToken(vf.firstToken))
		}
	} else if value == nil && (vf.hasSafe || ctx.safeOutput) {
		return vf.returnSafeValue(ctx)
	}
	return NewVarRefResult(value, false)
}

func (vf *VarRef) returnSafeValue(ctx *Context) *VarRefResult {
	var v interface{}
	if vf.safe == nil {
		v = nil
	} else {
		v = vf.safe.Evaluate(ctx)
	}
	return NewVarRefResult(v, true)
}

func (vf *VarRef) SetVarIndex(index int) {
	vf.varIndex = index
}

func (vf *VarRef) GetVarIndex() int {
	return vf.varIndex
}

func (vf *VarRef) EvaluateUntilLast(ctx *Context) interface{} {
	if len(vf.attributes) == 0 {
		panic("attributes is nil")
	}

	ret := vf.getValue(ctx)
	value := ret.Value
	if value == nil {
		panic(NewBeetlException(NULL, nil).SetToken(vf.firstToken))
	}

	for i := range vf.attributes {
		attr := vf.attributes[i]
		if value == nil {
			msg := "空指针"
			token := vf.firstToken
			if i != 0 {
				token = vf.attributes[i-1].GetASTNode().Token
			}
			panic(NewBeetlException(NULL, &msg).SetToken(token))
		}
		value = attr.Evaluate2(ctx, value)
	}

	return value
}

type VarRefResult struct {
	Value interface{}
	Safe  bool
}

func NewVarRefResult(value interface{}, safe bool) *VarRefResult {
	return &VarRefResult{
		Value: value,
		Safe:  safe,
	}
}
