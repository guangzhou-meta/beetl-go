package om

import (
	"reflect"
)

import (
	"github.com/guangzhou-meta/beetl-go/util"
)

type AttributeAccessI interface {
	Value(o interface{}, name interface{}) interface{}

	SetValue(o interface{}, name interface{}, value interface{})
}

type AttributeAccess struct {
}

func (aa *AttributeAccess) Value(o interface{}, name interface{}) interface{} {
	panic("not implement")
}
func (aa *AttributeAccess) SetValue(o interface{}, name interface{}, value interface{}) {
	aa.updateValue(o, name, value)
}

func (aa *AttributeAccess) updateValue(o interface{}, key interface{}, value interface{}) {
	t := util.UnpackType(reflect.TypeOf(o))
	v := util.UnpackValue(reflect.ValueOf(o))
	switch t.Kind() {
	case reflect.Map:
		v.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
	case reflect.Slice:
		v.Index(key.(int)).Set(reflect.ValueOf(value))
	default:
		v.FieldByName(key.(string)).Set(reflect.ValueOf(value))
	}
}

type ListAA struct {
	*AttributeAccess
}

func NewListAA() *ListAA {
	return &ListAA{}
}

func (aa *ListAA) Value(o interface{}, attr interface{}) interface{} {
	v := util.UnpackValue(reflect.ValueOf(o))
	return v.Index(attr.(int)).Interface()
}

type MapAA struct {
	*AttributeAccess
}

func NewMapAA() *MapAA {
	return &MapAA{}
}

func (aa *MapAA) Value(o interface{}, attr interface{}) interface{} {
	v := util.UnpackValue(reflect.ValueOf(o))
	return v.MapIndex(reflect.ValueOf(attr)).Interface()
}

type ReflectBeanAA struct {
	*AttributeAccess
}

func NewReflectBeanAA() *ReflectBeanAA {
	return &ReflectBeanAA{}
}

func (aa *ReflectBeanAA) Value(o interface{}, attr interface{}) interface{} {
	v := util.UnpackValue(reflect.ValueOf(o))
	return v.FieldByName(attr.(string)).Interface()
}
