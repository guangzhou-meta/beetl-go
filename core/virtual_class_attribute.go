package core

import (
	"reflect"
)

type VirtualClassAttribute interface {
	Eval(o interface{}, attributeName string, ctx *Context) interface{}
}

type VirtualAttributeEval interface {
	VirtualClassAttribute
	IsSupport(subClass reflect.Type, attributeName string) bool
}
