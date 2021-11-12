package om

import (
	"reflect"
)

var (
	defaultAAFactory *DefaultAAFactory
)

func init() {
	defaultAAFactory = NewDefaultAAFactory()
}

func BuildFiledAccessor(kind reflect.Kind) AttributeAccessI {
	return defaultAAFactory.BuildFiledAccessor(kind)
}

type DefaultAAFactory struct {
	listAA        *ListAA
	mapAA         *MapAA
	reflectBeanAA *ReflectBeanAA
	classAttrs    map[reflect.Kind]AttributeAccessI
}

func NewDefaultAAFactory() *DefaultAAFactory {
	classAttrs := make(map[reflect.Kind]AttributeAccessI)
	inst := &DefaultAAFactory{
		classAttrs:    classAttrs,
		listAA:        NewListAA(),
		mapAA:         NewMapAA(),
		reflectBeanAA: NewReflectBeanAA(),
	}
	classAttrs[reflect.Map] = inst.mapAA
	classAttrs[reflect.Slice] = inst.listAA
	return inst
}

func (f *DefaultAAFactory) BuildFiledAccessor(c reflect.Kind) AttributeAccessI {
	aa := f.classAttrs[c]
	if aa != nil {
		return aa
	}
	aa = f.registerClass(c)
	return aa
}

func (f *DefaultAAFactory) registerClass(c reflect.Kind) AttributeAccessI {
	f.classAttrs[c] = f.reflectBeanAA
	return f.reflectBeanAA
}
