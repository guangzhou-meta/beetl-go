package core

import (
	"reflect"
)

import (
	"github.com/guangzhou-meta/beetl-go/util"
)

import (
	"github.com/guangzhou-meta/go-lib/arrays"
)

type FunctionWrapperI interface {
	CheckContextRequired(paras []reflect.Type) bool
	GetContextParas(paras []interface{}, ctx *Context) []interface{}
	GetMethodName() string
}

type FunctionWrapper struct {
	requiredContext bool
	functionName    string
	target          reflect.Type
	cls             string
	ctxType         reflect.Type
	method          reflect.Value
}

func NewFunctionWrapper(funName string) *FunctionWrapper {
	return &FunctionWrapper{
		functionName: funName,
		ctxType:      reflect.TypeOf(NewContext(nil, nil)),
	}
}

func (fw *FunctionWrapper) CheckContextRequired(paras []reflect.Type) bool {
	return len(paras) != 0 && paras[len(paras)-1] == fw.ctxType
}

func (fw *FunctionWrapper) GetContextParas(paras []interface{}, ctx *Context) []interface{} {
	newParas := make([]interface{}, len(paras)+1)
	arrays.CopyFrom(paras, 0, newParas, 0, len(paras))
	newParas[len(paras)] = ctx
	return newParas
}

func (fw *FunctionWrapper) GetMethodName() string {
	return fw.functionName
}

func (fw *FunctionWrapper) Call(params []interface{}, ctx *Context) interface{} {
	var newParas []interface{}
	if fw.requiredContext {
		newParas = fw.GetContextParas(params, ctx)
	} else {
		newParas = params
	}
	values := []reflect.Value{reflect.New(fw.target)}
	for _, v := range newParas {
		if vv, ok := v.(reflect.Value); ok {
			values = append(values, vv)
		} else {
			values = append(values, reflect.ValueOf(v))
		}
	}
	return fw.method.Call(values)[0].Interface()
}

func (fw *FunctionWrapper) ClassName() string {
	panic("not implement")
}

func (fw *FunctionWrapper) GetClassType() reflect.Type {
	return fw.target
}

func GetFunctionWrapper(ns string, t reflect.Type) []FunctionWrapperI {
	info := util.GetObjectInfo(t)
	namespace := ns + "."
	var fwList []FunctionWrapperI
	for k, v := range info.GetMethods() {
		fw := NewFunctionWrapper(namespace + k)
		fw.target = t
		fw.method = v.Func
		var paras []reflect.Type
		for in := 0; in < v.Type.NumIn(); in++ {
			paras = append(paras, v.Type.In(in))
		}
		fw.requiredContext = fw.CheckContextRequired(paras)
		fwList = append(fwList, fw)
	}
	return fwList
}
