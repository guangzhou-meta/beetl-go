package core

import (
	"reflect"
)

type FormatI interface {
	Format(data interface{}, pattern string) interface{}

	ClassName() string

	GetClassType() reflect.Type
}

type ContextFormatI interface {
	Format2(data interface{}, pattern string, ctx *Context) interface{}
}

type ContextFormat struct {
}

func NewContextFormat() *ContextFormat {
	return &ContextFormat{}
}

func (cf *ContextFormat) Format(data interface{}, pattern string) interface{} {
	panic("unsupported")
}

func (cf *ContextFormat) Format2(data interface{}, pattern string, ctx *Context) interface{} {
	panic("not implement")
}

func (cf *ContextFormat) ClassName() string {
	panic("not implement")
}

func (cf *ContextFormat) GetClassType() reflect.Type {
	panic("not implement")
}
