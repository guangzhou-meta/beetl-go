package core

import (
	"reflect"
)

import (
	"github.com/guangzhou-meta/beetl-go/util"
)

type FunctionI interface {
	Call(params []interface{}, ctx *Context) interface{}

	ClassName() string

	GetClassType() reflect.Type
}

type Function struct {
}

func (f *Function) Call(params []interface{}, ctx *Context) interface{} {
	return nil
}

func (f *Function) ClassName() string {
	panic("not implement")
}

func (f *Function) GetClassType() reflect.Type {
	return f.GetType(f)
}

func (f *Function) GetType(i interface{}) reflect.Type {
	return util.UnpackType(reflect.TypeOf(i))
}

type SimpleFunction struct {
	*Function
}

func (f *SimpleFunction) Call(params []interface{}, ctx *Context) interface{} {
	return nil
}

func (f *SimpleFunction) ClassName() string {
	return "org.beetl.ext.fn.SimpleFunction"
}

func (f *SimpleFunction) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}
