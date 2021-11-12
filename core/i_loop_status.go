package core

import (
	"reflect"
)

import (
	"github.com/guangzhou-meta/beetl-go/util"
)

type ILoopStatus interface {
	HasNext() bool
	Next() interface{}
	GetIndex() int
	GetDataIndex() int
	IsFirst() bool
	IsLast() bool
	IsEven() bool
	IsOdd() bool
	HasSize() bool
	GetSize() int
	HasData() bool
}

type GeneralLoopType int

const (
	Unsupported GeneralLoopType = iota
	GeneralLoopTypeSlice
	GeneralLoopTypeMap
)

type GeneralLoopStatus struct {
	size    int
	index   int
	hasData bool

	data reflect.Value

	loopType GeneralLoopType
}

func NewGeneralLoopStatus() *GeneralLoopStatus {
	return &GeneralLoopStatus{
		size:     -1,
		index:    0,
		hasData:  false,
		loopType: Unsupported,
	}
}

func NewGeneralLoopStatusByList(list interface{}) *GeneralLoopStatus {
	inst := NewGeneralLoopStatus()
	data := util.UnpackValue(reflect.ValueOf(list))
	switch data.Kind() {
	case reflect.Slice, reflect.Array:
		inst.data = data
		inst.size = inst.data.Len()
		inst.loopType = GeneralLoopTypeSlice
		return inst
	}
	panic("unsupported this type: " + data.Kind().String())
}

func NewGeneralLoopStatusByMap(m interface{}) *GeneralLoopStatus {
	inst := NewGeneralLoopStatus()
	data := util.UnpackValue(reflect.ValueOf(m))
	switch data.Kind() {
	case reflect.Map:
		inst.data = data
		inst.size = len(data.MapKeys())
		inst.loopType = GeneralLoopTypeMap
		return inst
	}
	panic("unsupported this type: " + data.Kind().String())
}

func GetGeneralLoopIteratorStatus(o interface{}) *GeneralLoopStatus {
	t := util.UnpackType(reflect.TypeOf(o))
	switch t.Kind() {
	case reflect.Slice, reflect.Array:
		return NewGeneralLoopStatusByMap(o)
	case reflect.Map:
		return NewGeneralLoopStatusByMap(o)
	}
	return nil
}

func (gls *GeneralLoopStatus) HasNext() bool {
	return gls.index < gls.size-1
}

func (gls *GeneralLoopStatus) Next() interface{} {
	if !gls.hasData {
		gls.hasData = true
	}
	var res interface{}
	switch gls.loopType {
	case GeneralLoopTypeSlice:
		res = gls.data.Index(gls.index).Interface()
	case GeneralLoopTypeMap:
		k := gls.data.MapKeys()[gls.index]
		res = gls.data.MapIndex(k)
	}
	gls.index++
	return res
}
func (gls *GeneralLoopStatus) GetIndex() int {
	return gls.index
}
func (gls *GeneralLoopStatus) GetDataIndex() int {
	return gls.index - 1
}
func (gls *GeneralLoopStatus) IsFirst() bool {
	return gls.index == 1
}
func (gls *GeneralLoopStatus) IsLast() bool {
	return gls.index == gls.size
}
func (gls *GeneralLoopStatus) IsEven() bool {
	return gls.index%2 == 0
}
func (gls *GeneralLoopStatus) IsOdd() bool {
	return gls.index%2 == 1
}
func (gls *GeneralLoopStatus) HasSize() bool {
	return gls.size != -1
}
func (gls *GeneralLoopStatus) GetSize() int {
	if gls.size != -1 {
		return gls.size
	}
	panic("集合长度未知,请勿使用size")
}
func (gls *GeneralLoopStatus) HasData() bool {
	return gls.hasData
}
