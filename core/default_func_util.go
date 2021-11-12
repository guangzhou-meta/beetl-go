package core

import (
	"reflect"
)

type StringUtil struct {
	*Function
}

func (f *StringUtil) Call(params []interface{}, ctx *Context) interface{} {
	return nil
}

func (f *StringUtil) ClassName() string {
	return "org.beetl.ext.fn.StringUtil"
}

func (f *StringUtil) GetClassType() reflect.Type {
	return f.Function.GetType(f)
}
