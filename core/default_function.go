package core

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

import (
	"github.com/guangzhou-meta/beetl-go/common"
	"github.com/guangzhou-meta/beetl-go/util"
)

import (
	"github.com/guangzhou-meta/go-lib/arrays"
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

type DateFunction struct {
	*Function
}

func (f *DateFunction) Call(params []interface{}, ctx *Context) interface{} {
	if len(params) == 0 {
		return time.Now()
	} else if len(params) == 1 {
		if params[0] != nil {
			if i, ok := params[0].(int64); ok {
				return time.Unix(0, i*int64(time.Millisecond))
			}
		}
		panic("Parse date Error, Arg Long")
	} else if len(params) == 2 {
		p1, ok1 := params[0].(string)
		p2, ok2 := params[0].(string)
		if ok1 && ok2 {
			r, err := time.Parse(p2, p1)
			if err != nil {
				panic(err)
			}
			return r
		}
	}
	panic("Parse date Error,Args String, Sting")
}

func (f *DateFunction) ClassName() string {
	return "org.beetl.ext.fn.DateFunction"
}

func (f *DateFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type EmptyExpressionFunction struct {
	*Function
}

func (f *EmptyExpressionFunction) Call(params []interface{}, ctx *Context) interface{} {
	switch len(params) {
	case 0:
		return true
	case 1:
		return f.isEmpty(params[0])
	}
	for _, p := range params {
		if !f.isEmpty(p) {
			return false
		}
	}
	return true
}

func (f *EmptyExpressionFunction) ClassName() string {
	return "org.beetl.ext.fn.EmptyExpressionFunction"
}

func (f *EmptyExpressionFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

func (f *EmptyExpressionFunction) isEmpty(obj interface{}) bool {
	if obj == nil {
		return true
	}
	t := util.UnpackValue(reflect.ValueOf(obj))
	switch t.Kind() {
	case reflect.String:
		return len(t.Interface().(string)) == 0
	case reflect.Slice, reflect.Array:
		return t.Len() == 0
	case reflect.Map:
		return len(t.MapKeys()) == 0
	}
	return false
}

type NVLFunction struct {
	*Function
}

func (f *NVLFunction) Call(params []interface{}, ctx *Context) interface{} {
	if len(params) != 2 {
		panic("参数错误，期望Object, Object")
	}
	if params[0] == nil {
		return params[1]
	}
	return params[0]
}

func (f *NVLFunction) ClassName() string {
	return "org.beetl.ext.fn.NVLFunction"
}

func (f *NVLFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type DebugFunction struct {
	*Function
}

func (f *DebugFunction) Call(params []interface{}, ctx *Context) interface{} {
	if !ctx.gt.Function.debugEnable {
		return ""
	}
	expStrs := params[len(params)-2].([]interface{})
	sb := estrings.NewStringBuilder()
	for i := 0; i < len(params)-2; i++ {
		o := params[i]
		if expStrs[i] != nil {
			sb.Append(expStrs[i].(string)).Append("=")
		}
		if o != nil {
			if oo, ok := o.(string); ok {
				sb.Append("\"").Append(oo).Append("\"")
			} else {
				sb.Append(oo)
			}
		} else {
			sb.Append("null")
		}
		sb.Append(",")
	}
	line := params[len(params)-1].(string)
	resourceId := ctx.GetResourceId()
	sb.Append(" [在").Append(line).Append("行@").Append(resourceId).Append("]")
	fmt.Println(sb.String())
	return ""
}

func (f *DebugFunction) ClassName() string {
	return "org.beetl.ext.fn.DebugFunction"
}

func (f *DebugFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type CheckExistFunction struct {
	*Function
}

func (f *CheckExistFunction) Call(params []interface{}, ctx *Context) interface{} {
	if ctx.GlobalVar == nil {
		return false
	}
	var key string
	for _, o := range params {
		key = o.(string)
		if _, ok := ctx.GlobalVar[key]; !ok {
			return false
		}
	}
	return true
}

func (f *CheckExistFunction) ClassName() string {
	return "org.beetl.ext.fn.CheckExistFunction"
}

func (f *CheckExistFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type PrintfFunction struct {
	*Function
}

func (f *PrintfFunction) Call(params []interface{}, ctx *Context) interface{} {
	template := params[0].(string)
	args := make([]interface{}, len(params)-1)
	arrays.CopyFrom(params, 1, args, 0, len(args))
	sb := fmt.Sprintf(template, args...)
	ctx.GetByteWriter().WriteString(sb)
	//ctx.GetByteWriter().Append(sb)
	return ""
}

func (f *PrintfFunction) ClassName() string {
	return "org.beetl.ext.fn.Printf"
}

func (f *PrintfFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type DecodeFunction struct {
	*Function
}

func (f *DecodeFunction) Call(params []interface{}, ctx *Context) interface{} {
	var ret interface{}
	o := params[0]
	i := 1
	for {
		if f.same(o, params[i], ctx) {
			ret = params[i+1]
			break
		}
		if len(params)-1 == i+2 {
			ret = params[i+2]
			break
		} else {
			i = i + 2
		}
	}
	return f.unwrap(ret, ctx)
}

func (f *DecodeFunction) ClassName() string {
	return "org.beetl.ext.fn.DecodeFunction"
}

func (f *DecodeFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

func (f *DecodeFunction) same(a interface{}, b interface{}, ctx *Context) bool {
	return util.ObjectsAreEqualValues(a, f.unwrap(b, ctx))
}

func (f *DecodeFunction) unwrap(b interface{}, ctx *Context) interface{} {
	if bb, ok := b.(ExpressionRuntimeObject); ok {
		return bb.Get(ctx)
	}
	return b
}

type AssertFunction struct {
	*Function
}

func (f *AssertFunction) Call(params []interface{}, ctx *Context) interface{} {
	if ctx.gt.Function.assertEnable {
		result := params[0].(bool)
		var msg string
		if len(params) > 1 {
			msg = params[1].(string)
		} else {
			msg = "断言异常"
		}
		if !result {
			panic(msg)
		}
	}
	return ""
}

func (f *AssertFunction) ClassName() string {
	return "org.beetl.ext.fn.AssertFunction"
}

func (f *AssertFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type PrintFunction struct {
	*Function
}

func (f *PrintFunction) Call(params []interface{}, ctx *Context) interface{} {
	o := params[0]
	if o != nil {
		ctx.GetByteWriter().WriteString(o.(string))
		//ctx.GetByteWriter().Append(o.(string))
	}
	return ""
}

func (f *PrintFunction) ClassName() string {
	return "org.beetl.ext.fn.Print"
}

func (f *PrintFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type PrintlnFunction struct {
	*Function
}

func (f *PrintlnFunction) Call(params []interface{}, ctx *Context) interface{} {
	w := ctx.GetByteWriter()
	if len(params) == 0 {
		w.WriteString(ctx.Template.program.MetaData.LineSeparator)
		//w.Append(ctx.Template.program.MetaData.LineSeparator)
		return ""
	}
	o := params[0]
	if o != nil {
		w.WriteString(o.(string))
		w.WriteString(ctx.Template.program.MetaData.LineSeparator)
		//w.Append(o.(string))
		//w.Append(ctx.Template.program.MetaData.LineSeparator)
	}
	return ""
}

func (f *PrintlnFunction) ClassName() string {
	return "org.beetl.ext.fn.Println"
}

func (f *PrintlnFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type PrintFileFunction struct {
	*Function
}

func (f *PrintFileFunction) Call(params []interface{}, ctx *Context) interface{} {
	ctx.byteWriter.WriteString("not support PrintFile")
	//ctx.byteWriter.Append("not support PrintFile")
	return ""
}

func (f *PrintFileFunction) ClassName() string {
	return "org.beetl.ext.fn.PrintFile"
}

func (f *PrintFileFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type TruncFunction struct {
	*Function
}

func (f *TruncFunction) Call(params []interface{}, ctx *Context) interface{} {
	n := params[0]
	pos := 0
	if len(params) != 1 {
		pos = params[1].(int)
	}
	if pos == 0 {
		return n.(int64)
	}
	if nn, ok := n.(*big.Float); ok {
		return nn.SetPrec(uint(pos)).SetMode(big.ToPositiveInf)
	}
	nn := new(big.Float).SetFloat64(n.(float64))
	v, _ := nn.SetPrec(uint(pos)).SetMode(big.ToPositiveInf).Float64()
	return v
}

func (f *TruncFunction) ClassName() string {
	return "org.beetl.ext.fn.TruncFunction"
}

func (f *TruncFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type TruncFunction2 struct {
	*Function
}

func (f *TruncFunction2) Call(params []interface{}, ctx *Context) interface{} {
	obj := params[0]
	if obj == nil {
		return nil
	}
	if tt, ok := obj.(*time.Time); ok {
		var p *string
		if len(params) == 1 {
			p = nil
		} else {
			pp := params[1].(string)
			p = &pp
		}
		return f.truncateDate(tt, p)
	}
	nT := util.IsNumberType(obj)
	if nT != util.NumberTypeUnknown {
		var p interface{}
		if len(params) == 1 {
			p = nil
		} else {
			p = params[1]
		}
		return f.truncateNumber(obj, p)
	}
	panic("unsupported truncate: " + reflect.TypeOf(obj).String())
}

func (f *TruncFunction2) ClassName() string {
	return "org.beetl.ext.fn.TruncFunction2"
}

func (f *TruncFunction2) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

func (f *TruncFunction2) truncateDate(d *time.Time, p *string) string {
	if p == nil {
		pp := "2006-01-02"
		p = &pp
	}
	return d.Format(*p)
}

func (f *TruncFunction2) truncateNumber(obj interface{}, pos interface{}) string {
	str := obj.(string)
	index := strings.Index(str, ".")
	if index == -1 {
		return str
	}
	if pos == nil {
		return str[0:index]
	}
	p := pos.(int)
	if p == 0 {
		panic("参数不能为0")
	}
	dig := len(str) - index - 1
	if dig >= p {
		return str[0 : index+p+1]
	}
	return str
}

type QuestionMarkFunction struct {
	*Function
}

func (f *QuestionMarkFunction) Call(params []interface{}, ctx *Context) interface{} {
	if len(params) != 3 {
		panic("请输入问号表达式格式： qmark(a==\"a\",\"yes\",\"no\")")
	}
	b := params[0].(bool)
	if b {
		return params[1]
	}
	return params[2]
}

func (f *QuestionMarkFunction) ClassName() string {
	return "org.beetl.ext.fn.QuestionMark"
}

func (f *QuestionMarkFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type IsNotEmptyExpressionFunction struct {
	*Function
}

func (f *IsNotEmptyExpressionFunction) Call(params []interface{}, ctx *Context) interface{} {
	r := new(EmptyExpressionFunction).Call(params, ctx)
	return !(r.(bool))
}

func (f *IsNotEmptyExpressionFunction) ClassName() string {
	return "org.beetl.ext.fn.IsNotEmptyExpressionFunction"
}

func (f *IsNotEmptyExpressionFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type ParseIntFunction struct {
	*Function
}

func (f *ParseIntFunction) Call(params []interface{}, ctx *Context) interface{} {
	o := params[0]
	str := ""
	if o == nil {
		panic("Error:parseInt(nil)")
	}
	nT := util.IsNumberType(o)
	if nT == util.NumberTypeUnknown {
		str = o.(string)
	} else {
		str = fmt.Sprintf("%v", o)
	}
	if strings.Contains(str, ".") {
		str = strings.Split(str, "\\.")[0]
	}
	isNum := regexp.MustCompile(`-?[0-9]*`).Match([]byte(str))
	if "" == o || !isNum {
		panic("无法正确转换至int格式")
	}
	result, _ := strconv.ParseInt(str, 10, 32)
	return int(result)
}

func (f *ParseIntFunction) ClassName() string {
	return "org.beetl.ext.fn.ParseInt"
}

func (f *ParseIntFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type ParseLongFunction struct {
	*Function
}

func (f *ParseLongFunction) Call(params []interface{}, ctx *Context) interface{} {
	o := params[0]
	str := ""
	if o == nil {
		panic("Error:parseLong(nil)")
	}
	nT := util.IsNumberType(o)
	if nT == util.NumberTypeUnknown {
		str = o.(string)
	} else {
		str = fmt.Sprintf("%v", o)
	}
	if strings.Contains(str, ".") {
		str = strings.Split(str, "\\.")[0]
	}
	isNum := regexp.MustCompile(`-?[0-9]*`).Match([]byte(str))
	if "" == o || !isNum {
		panic("无法正确转换至int64格式")
	}
	result, _ := strconv.ParseInt(str, 10, 64)
	return result
}

func (f *ParseLongFunction) ClassName() string {
	return "org.beetl.ext.fn.ParseLong"
}

func (f *ParseLongFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type ParseDoubleFunction struct {
	*Function
}

func (f *ParseDoubleFunction) Call(params []interface{}, ctx *Context) interface{} {
	o := params[0]
	str := ""
	if o == nil {
		panic("Error:parseDouble(nil)")
	}
	if util.IsNumberType(o) == util.NumberTypeUnknown {
		str = o.(string)
	} else {
		str = fmt.Sprintf("%v", o)
	}
	result, _ := strconv.ParseFloat(str, 64)
	return result
}

func (f *ParseDoubleFunction) ClassName() string {
	return "org.beetl.ext.fn.ParseDouble"
}

func (f *ParseDoubleFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type RangeFunction struct {
	*Function
}

func (f *RangeFunction) Call(params []interface{}, ctx *Context) interface{} {
	panic("unsupported")
	return nil
}

func (f *RangeFunction) ClassName() string {
	return "org.beetl.ext.fn.Range"
}

func (f *RangeFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type FlushFunction struct {
	*Function
}

func (f *FlushFunction) Call(params []interface{}, ctx *Context) interface{} {
	//ctx.byteWriter.Flush()
	return nil
}

func (f *FlushFunction) ClassName() string {
	return "org.beetl.ext.fn.Flush"
}

func (f *FlushFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type JsonFunction struct {
	*Function
}

func (f *JsonFunction) Call(params []interface{}, ctx *Context) interface{} {
	obj := params[0]
	if obj == nil {
		return ""
	}
	r, _ := json.Marshal(obj)
	return string(r)
}

func (f *JsonFunction) ClassName() string {
	return "org.beetl.ext.fn.Json"
}

func (f *JsonFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type PageContextFunction struct {
	*Function
}

func (f *PageContextFunction) Call(params []interface{}, ctx *Context) interface{} {
	m := ctx.GetGlobal("$page")
	if m == nil {
		panic("没有找到$page全局变量")
	}
	mm := m.(map[string]interface{})
	key := params[0].(string)
	if len(params) == 1 {
		return mm[key]
	} else {
		v, ok := mm[key]
		mm[key] = params[1]
		if ok {
			return v
		}
		return nil
	}
}

func (f *PageContextFunction) ClassName() string {
	return "org.beetl.ext.fn.PageContextFunction"
}

func (f *PageContextFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type TypeNewFunction struct {
	*Function
}

func (f *TypeNewFunction) Call(params []interface{}, ctx *Context) interface{} {
	panic("unsupported")
}

func (f *TypeNewFunction) ClassName() string {
	return "org.beetl.ext.fn.TypeNewFunction"
}

func (f *TypeNewFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type TypeNameFunction struct {
	*Function
}

func (f *TypeNameFunction) Call(params []interface{}, ctx *Context) interface{} {
	panic("unsupported")
}

func (f *TypeNameFunction) ClassName() string {
	return "org.beetl.ext.fn.TypeNameFunction"
}

func (f *TypeNameFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type DynamicGlobalValueFunction struct {
	*Function
}

func (f *DynamicGlobalValueFunction) Call(params []interface{}, ctx *Context) interface{} {
	key := params[0].(string)
	return ctx.GetGlobal(key)
}

func (f *DynamicGlobalValueFunction) ClassName() string {
	return "org.beetl.ext.fn.DynamicGlobalValueFunction"
}

func (f *DynamicGlobalValueFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type AllGlobaAsJsonlFunction struct {
	*Function
}

func (f *AllGlobaAsJsonlFunction) Call(params []interface{}, ctx *Context) interface{} {
	res, _ := json.Marshal(ctx.GlobalVar)
	return string(res)
}

func (f *AllGlobaAsJsonlFunction) ClassName() string {
	return "org.beetl.ext.fn.AllGlobaAsJsonlFunction"
}

func (f *AllGlobaAsJsonlFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type HasAttributeFunction struct {
	*Function
}

func (f *HasAttributeFunction) Call(params []interface{}, ctx *Context) interface{} {
	o := params[0]
	if o == nil {
		panic("请传入参数")
	}
	if _, ok := o.(map[interface{}]interface{}); ok {
		return false
	}
	t := util.UnpackType(reflect.TypeOf(o))
	for i := 1; i < len(params); i++ {
		key := params[i].(string)
		if _, b := t.MethodByName(key); !b {
			return false
		}

	}
	return true
}

func (f *HasAttributeFunction) ClassName() string {
	return "org.beetl.ext.fn.HasAttributeFunction"
}

func (f *HasAttributeFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type EnvFunction struct {
	*Function
}

func (f *EnvFunction) Call(params []interface{}, ctx *Context) interface{} {
	cr := "\n"
	if _, ok := ctx.GlobalVar["$page"]; ok {
		cr = "<br/>"
	}
	cf := ctx.gt.conf
	var gvks []string
	for k := range ctx.GlobalVar {
		gvks = append(gvks, k)
	}
	sb := estrings.NewStringBuilder()
	sb.Append("resource:").
		Append(ctx.GetResourceId()).
		Append(cr).
		Append("global:").
		Append(gvks).
		Append(cr).
		Append("delimeter :").
		Append(cf.placeholderStart).
		Append(cf.placeholderEnd).
		Append(" , ").
		Append(cf.statementStart).
		Append(" ").
		Append(cf.statementEnd)
	return sb.String()
}

func (f *EnvFunction) ClassName() string {
	return "org.beetl.ext.fn.EnvFunction"
}

func (f *EnvFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type ParentTagFunction struct {
	*Function
}

func (f *ParentTagFunction) Call(params []interface{}, ctx *Context) interface{} {
	current := ctx.GetCurrentTag()
	if current != nil {
		return current.parent
	}
	return nil
}

func (f *ParentTagFunction) ClassName() string {
	return "org.beetl.ext.fn.ParentTagFunction"
}

func (f *ParentTagFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type PageQueryFunction struct {
	*Function
}

func (f *PageQueryFunction) Call(params []interface{}, ctx *Context) interface{} {
	o := ctx.GetGlobal(common.DBAutoGeneratedSqlPAGE)
	if o != nil && o.(bool) {
		return "count(1)"
	}
	if len(params) == 0 {
		return "*"
	}
	return params[0].(string)
}

func (f *PageQueryFunction) ClassName() string {
	return "org.beetl.ext.fn.PageQueryFunction"
}

func (f *PageQueryFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}

type IsBlank struct {
	*Function
}

func (f *IsBlank) Call(params []interface{}, ctx *Context) interface{} {
	o := params[0]
	if o == nil {
		return true
	}
	if cs, ok := o.(string); ok {
		return len(strings.TrimSpace(cs)) == 0
	}
	return false
}

func (f *IsBlank) ClassName() string {
	return "org.beetl.ext.fn.IsBlank"
}

func (f *IsBlank) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}
