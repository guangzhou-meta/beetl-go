package core

import (
	"reflect"
)

import (
	"github.com/guangzhou-meta/beetl-go/core/om"
	"github.com/guangzhou-meta/beetl-go/util"
)

type FunctionExpressionI interface {
	ExpressionI

	GetFunctionExpression() FunctionExpressionI

	GetToken() *GrammarToken
}

type FunctionExpression struct {
	*Expression

	name    string
	exps    []ExpressionI
	vas     []VarAttributeI
	safeExp ExpressionI
	hasSafe bool
}

func NewFunctionExpression(name string, exps []ExpressionI, vas []VarAttributeI, hasSafe bool, safeExp ExpressionI, token *GrammarToken) *FunctionExpression {
	return &FunctionExpression{
		Expression: NewExpression(token),
		name:       name,
		exps:       exps,
		vas:        vas,
		safeExp:    safeExp,
		hasSafe:    hasSafe,
	}
}

func (fe *FunctionExpression) GetFunctionExpression() FunctionExpressionI {
	return fe
}

func (fe *FunctionExpression) GetToken() *GrammarToken {
	return fe.Token
}

func (fe *FunctionExpression) Evaluate(ctx *Context) interface{} {
	fn := ctx.gt.GetFunction(fe.name)
	if fn == nil {
		return nil
	}
	paras := make([]interface{}, len(fe.exps))
	for i := range paras {
		paras[i] = fe.exps[i].Evaluate(ctx)
	}
	value := fn.Call(paras, ctx)
	if fe.vas != nil {
		for _, attr := range fe.vas {
			value = attr.Evaluate2(ctx, value)
			if value == nil {
				if fe.hasSafe {
					if fe.safeExp == nil {
						return nil
					}
					return fe.safeExp.Evaluate(ctx)
				} else {
					msg := "空指针"
					panic(NewBeetlException(ERROR, &msg).SetToken(attr.GetASTNode().Token))
				}
			}
		}
	}
	if value == nil && fe.hasSafe {
		if fe.safeExp == nil {
			return nil
		}
		return fe.safeExp.Evaluate(ctx)
	}
	return value
}

func (fe *FunctionExpression) GetResource(gt *GroupTemplate, name string) ResourceI {
	//resourceMap := gt.conf.resourceMap
	//unctionSuffix := resourceMap["functionSuffix"]
	//functionRoot := resourceMap["functionRoot"]
	//path:=strings.Replace(name,".","/",1)
	//return gt.
	return nil
}

var (
	NULLLiteral = NewLiteral(nil, nil)
)

type Literal struct {
	*Expression

	obj interface{}
}

func NewLiteral(value interface{}, token *GrammarToken) *Literal {
	return &Literal{
		Expression: NewExpression(token),
		obj:        value,
	}
}

func (ex *Literal) Evaluate(ctx *Context) interface{} {
	return ex.obj
}

const (
	CompareExpressionEQUAL       = 0
	CompareExpressionNOT_EQUAL   = 1
	CompareExpressionLARGE       = 2
	CompareExpressionLARGE_EQUAL = 3
	CompareExpressionLESS        = 4
	CompareExpressionLESS_EQUAL  = 5
)

type CompareExpression struct {
	*Expression

	compareMode int

	a ExpressionI
	b ExpressionI
}

func NewCompareExpression(a ExpressionI, b ExpressionI, mode int, token *GrammarToken) *CompareExpression {
	return &CompareExpression{
		Expression:  NewExpression(token),
		a:           a,
		b:           b,
		compareMode: mode,
	}
}

func (ex *CompareExpression) Evaluate(ctx *Context) interface{} {
	x := ex.a.Evaluate(ctx)
	y := ex.b.Evaluate(ctx)
	switch ex.compareMode {
	case CompareExpressionEQUAL:
		return util.ObjectsAreEqualValues(x, y)

	case CompareExpressionNOT_EQUAL:
		return !util.ObjectsAreEqualValues(x, y)
	case CompareExpressionLARGE_EQUAL:
		return !util.ObjectsValuesLess(x, y)

	case CompareExpressionLARGE:
		return !util.ObjectsValuesLessEquals(x, y)
	case CompareExpressionLESS:
		return util.ObjectsValuesLess(x, y)
	case CompareExpressionLESS_EQUAL:
		return util.ObjectsValuesLessEquals(x, y)
	default:
		panic("类型判断失败")
	}
}

type TernaryExpression struct {
	*Expression

	condition ExpressionI
	a         ExpressionI
	b         ExpressionI
}

func NewTernaryExpression(condition ExpressionI, a ExpressionI, b ExpressionI, token *GrammarToken) *TernaryExpression {
	return &TernaryExpression{
		Expression: NewExpression(token),
		condition:  condition,
		a:          a,
		b:          b,
	}
}

func (ex *TernaryExpression) Evaluate(ctx *Context) interface{} {
	value := ex.condition.Evaluate(ctx)
	if value == nil {
		panic(NewBeetlException(NULL, nil).SetToken(ex.condition.GetASTNode().Token))
	}
	var v bool
	var ok bool
	if v, ok = value.(bool); !ok {
		panic(NewBeetlException(BOOLEAN_EXPECTED_ERROR, nil).SetToken(ex.condition.GetASTNode().Token))
	}
	if v {
		return ex.a.Evaluate(ctx)
	}
	if ex.b != nil {
		return ex.b.Evaluate(ctx)
	}
	return nil
}

const (
	PLUS = 0
	MIN  = 1
	MUL  = 2
	DIV  = 3
	MOD  = 4
)

type ArthExpression struct {
	*Expression

	arthMode int
	a        ExpressionI
	b        ExpressionI
}

func NewArthExpression(a ExpressionI, b ExpressionI, mode int, token *GrammarToken) *ArthExpression {
	return &ArthExpression{
		Expression: NewExpression(token),
		a:          a,
		b:          b,
		arthMode:   mode,
	}
}

func (ex *ArthExpression) Evaluate(ctx *Context) interface{} {
	x := ex.a.Evaluate(ctx)
	y := ex.b.Evaluate(ctx)
	switch ex.arthMode {
	case PLUS:
		return util.Plus(x, y)
	case MIN:
		return util.Minus(x, y)
	case MUL:
		return util.Mult(x, y)
	case DIV:
		return util.Div(x, y)
	case MOD:
		return util.Mod(x, y)
	default:
		panic("unsupported")
	}
}

type JsonArrayExpression struct {
	*Expression
	array []ExpressionI
}

func NewJsonArrayExpression(array []ExpressionI, token *GrammarToken) *JsonArrayExpression {
	return &JsonArrayExpression{
		Expression: NewExpression(token),
		array:      array,
	}
}

func (ex *JsonArrayExpression) Evaluate(ctx *Context) interface{} {
	if len(ex.array) == 0 {
		return make([]interface{}, 0)
	}
	var values []interface{}
	for _, exp := range ex.array {
		values = append(values, exp.Evaluate(ctx))
	}
	return values
}

type JsonMapExpression struct {
	*Expression
	array []ExpressionI
	m     map[string]ExpressionI
}

func NewJsonMapExpression(m map[string]ExpressionI, token *GrammarToken) *JsonMapExpression {
	var array []ExpressionI
	for _, v := range m {
		array = append(array, v)
	}
	return &JsonMapExpression{
		Expression: NewExpression(token),
		m:          m,
		array:      array,
	}
}

func (ex *JsonMapExpression) Evaluate(ctx *Context) interface{} {
	if len(ex.m) == 0 {
		return make(map[string]interface{})
	}
	values := make(map[string]interface{})
	for k, v := range ex.m {
		values[k] = v.Evaluate(ctx)
	}
	return values
}

type NativeCallExpression struct {
	*Expression
	insNode *InstanceNode
	clsNode *ClassNode
	chain   []NativeNodeI
}

func NewNativeCallExpression(insNode *InstanceNode, clsNode *ClassNode, chain []NativeNodeI, token *GrammarToken) *NativeCallExpression {
	return &NativeCallExpression{
		Expression: NewExpression(token),
		insNode:    insNode,
		clsNode:    clsNode,
		chain:      chain,
	}
}

func (ex *NativeCallExpression) checkNull(o interface{}, node NativeNodeI) {
	if o == nil {
		panic(NewBeetlException(NULL, nil).SetToken(NewGrammarToken(node.GetName(), ex.GetASTNode().Token.Line, ex.GetASTNode().Token.Col)))
	}
}

func (ex *NativeCallExpression) Evaluate(ctx *Context) interface{} {
	panic("not implement")
}

type AndExpression struct {
	*Expression
	exp1 ExpressionI
	exp2 ExpressionI
}

func NewAndExpression(exp1 ExpressionI, exp2 ExpressionI, token *GrammarToken) *AndExpression {
	return &AndExpression{
		Expression: NewExpression(token),
		exp1:       exp1,
		exp2:       exp2,
	}
}

func (ex *AndExpression) Evaluate(ctx *Context) interface{} {
	if v1, ok := ex.exp1.Evaluate(ctx).(bool); ok && v1 {
		v2, ok := ex.exp2.Evaluate(ctx).(bool)
		return ok && v2
	}
	return false
}

type OrExpression struct {
	*Expression
	exp1 ExpressionI
	exp2 ExpressionI
}

func NewOrExpression(exp1 ExpressionI, exp2 ExpressionI, token *GrammarToken) *OrExpression {
	return &OrExpression{
		Expression: NewExpression(token),
		exp1:       exp1,
		exp2:       exp2,
	}
}

func (ex *OrExpression) Evaluate(ctx *Context) interface{} {
	if v1, ok := ex.exp1.Evaluate(ctx).(bool); ok && v1 {
		return true
	}
	v2, ok := ex.exp2.Evaluate(ctx).(bool)
	return ok && v2
}

type NotBooleanExpression struct {
	*Expression
	exp ExpressionI
}

func NewNotBooleanExpression(exp ExpressionI, token *GrammarToken) *NotBooleanExpression {
	return &NotBooleanExpression{
		Expression: NewExpression(token),
		exp:        exp,
	}
}

func (ex *NotBooleanExpression) Evaluate(ctx *Context) interface{} {
	v2, ok := ex.exp.Evaluate(ctx).(bool)
	return !(ok && v2)
}

type NegExpression struct {
	*Expression
	exp ExpressionI
}

func NewNegExpression(exp ExpressionI, token *GrammarToken) *NegExpression {
	return &NegExpression{
		Expression: NewExpression(token),
		exp:        exp,
	}
}

func (ex *NegExpression) Evaluate(ctx *Context) interface{} {
	return util.Negative(ex.exp.Evaluate(ctx))
}

type IncDecExpression struct {
	*Expression
	isInc          bool
	returnOriginal bool
	index          int
}

func NewIncDecExpression(isInc bool, returnOriginal bool, token *GrammarToken) *IncDecExpression {
	return &IncDecExpression{
		Expression:     NewExpression(token),
		isInc:          isInc,
		returnOriginal: returnOriginal,
	}
}

func (ex *IncDecExpression) Evaluate(ctx *Context) interface{} {
	c := ctx.Vars[ex.index]
	var newValue interface{}
	if ex.isInc {
		newValue = util.Plus(c, 1)
	} else {
		newValue = util.Minus(c, 1)
	}
	ctx.Vars[ex.index] = newValue
	if ex.returnOriginal {
		return c
	}
	return newValue
}

func (ex *IncDecExpression) SetVarIndex(index int) {
	ex.index = index
}

func (ex *IncDecExpression) GetVarIndex() int {
	return ex.index
}

type VarRefAssignExpress struct {
	*Expression
	exp              ExpressionI
	varRef           *VarRef
	lastVarAttribute VarAttributeI
	varIndex         int
}

func NewVarRefAssignExpress(exp ExpressionI, varRef *VarRef) *VarRefAssignExpress {
	inst := &VarRefAssignExpress{
		Expression: NewExpression(varRef.Token),
		exp:        exp,
		varRef:     varRef,
	}
	if len(varRef.attributes) == 0 {
		inst.lastVarAttribute = nil
	} else {
		inst.lastVarAttribute = varRef.attributes[len(varRef.attributes)-1]
	}
	return inst
}

func (ex *VarRefAssignExpress) Evaluate(ctx *Context) interface{} {
	value := ex.exp.Evaluate(ctx)
	if ex.lastVarAttribute == nil {
		ctx.Vars[ex.varIndex] = value
		return value
	}
	obj := ex.varRef.EvaluateUntilLast(ctx)
	var key interface{}
	if ev, ok := ex.lastVarAttribute.(*VarSquareAttribute); ok {
		key = ev.Evaluate(ctx)
	} else {
		key = ex.lastVarAttribute.GetName()
	}
	if obj == nil {
		panic(NewBeetlException(NULL, nil).SetToken(ex.varRef.Token))
	}
	aa := om.BuildFiledAccessor(util.UnpackType(reflect.TypeOf(obj)).Kind())
	aa.SetValue(obj, key, value)
	return value
}

func (ex *VarRefAssignExpress) SetVarIndex(index int) {
	ex.varIndex = index
}

func (ex *VarRefAssignExpress) GetVarIndex() int {
	return ex.varIndex
}

type ContentBodyExpression struct {
	*Expression
	block *BlockStatement
}

func NewContentBodyExpression(block *BlockStatement, token *GrammarToken) *ContentBodyExpression {
	return &ContentBodyExpression{
		Expression: NewExpression(token),
		block:      block,
	}
}

func (ex *ContentBodyExpression) Evaluate(ctx *Context) interface{} {
	real := ctx.byteWriter
	temp := real.GetTempWriter(real)
	//temp := estrings.NewStringBuilderStr(real.String())
	ctx.byteWriter = temp
	ex.block.Execute(ctx)
	ctx.byteWriter = real
	return temp.GetTempContent()
	//str := []byte(temp.String())
	//return NewByteBodyContent(str, len(str))
}

type ExpressionRuntime struct {
	*Expression
	express ExpressionI
}

func NewExpressionRuntime(express ExpressionI) *ExpressionRuntime {
	return &ExpressionRuntime{
		Expression: NewExpression(express.GetASTNode().Token),
		express:    express,
	}
}

func (ex *ExpressionRuntime) Evaluate(ctx *Context) interface{} {
	return NewExpressionRuntimeObject(ex)
}

type ExpressionRuntimeObject struct {
	runtime *ExpressionRuntime
}

func NewExpressionRuntimeObject(runtime *ExpressionRuntime) *ExpressionRuntimeObject {
	return &ExpressionRuntimeObject{
		runtime: runtime,
	}
}

func (ex *ExpressionRuntimeObject) Get(ctx *Context) interface{} {
	return ex.runtime.express.Evaluate(ctx)
}
