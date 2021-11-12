package util

import (
	"bytes"
	"math/big"
	"reflect"
	"unsafe"
)

type ObjectInfo struct {
	allMethod map[string]reflect.Method
	c         reflect.Type
}

func NewObjectInfo(t reflect.Type) *ObjectInfo {
	inst := &ObjectInfo{
		c:         t,
		allMethod: make(map[string]reflect.Method),
	}
	inst.init()
	return inst
}

func (o *ObjectInfo) init() {
	for i := 0; i < o.c.NumMethod(); i++ {
		m := o.c.Method(i)
		o.allMethod[m.Name] = m
	}
}

func (o *ObjectInfo) GetMethods() map[string]reflect.Method {
	return o.allMethod
}

func GetObjectInfo(c reflect.Type) *ObjectInfo {
	return NewObjectInfo(c)
}

func ObjectsAreEqual(expected interface{}, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}

func ObjectsAreEqualValues(expected interface{}, actual interface{}) bool {
	if ObjectsAreEqual(expected, actual) {
		return true
	}

	actualType := reflect.TypeOf(actual)
	if actualType == nil {
		return false
	}
	expectedValue := reflect.ValueOf(expected)
	if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
		// Attempt comparison after type conversion
		return reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
	}

	return false
}

func GetAluType(left interface{}, right interface{}) int {
	if left == nil || right == nil {
		return AluNULL
	}

	switch left.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		switch right.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			return AluINTEGER
		}
	case float32, float64:
		switch right.(type) {
		case float32, float64:
			return AluFLOAT
		}
	}
	return AluSTRING
}

func GetAluTypeSingle(left interface{}) int {
	if left == nil {
		return AluNULL
	}
	switch left.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return AluINTEGER
	case float32, float64:
		return AluFLOAT
	}
	return AluSTRING
}

func CompareObject(o1 interface{}, o2 interface{}) int {
	l, lok := o1.(ComparableI)
	r, rok := o2.(ComparableI)
	if lok && rok {
		result := l.CompareTo(r)
		if result > 0 {
			return 1
		}
		return -1
	}
	return -2
}

func ObjectsValuesLess(o1 interface{}, o2 interface{}) bool {
	if !(o1 == nil || o2 == nil) {
		switch GetAluType(o1, o2) {
		case AluINTEGER:
			return o1.(int64) < o2.(int64)
		case AluFLOAT:
			return o1.(float64) < o2.(float64)
		default:
			result := CompareObject(o1, o2)
			if result != -2 {
				return result < 0
			}
		}
	}
	return false
}

func ObjectsValuesLessEquals(o1 interface{}, o2 interface{}) bool {
	if !(o1 == nil || o2 == nil) {
		switch GetAluType(o1, o2) {
		case AluINTEGER:
			return o1.(int64) <= o2.(int64)
		case AluFLOAT:
			return o1.(float64) <= o2.(float64)
		default:
			result := CompareObject(o1, o2)
			if result != -2 {
				return result <= 0
			}
		}
	}
	return false
}

func Plus(o1 interface{}, o2 interface{}) interface{} {
	if !(o1 == nil || o2 == nil) {
		switch GetAluType(o1, o2) {
		case AluINTEGER:
			return o1.(int64) + o2.(int64)
		case AluFLOAT:
			return o1.(float64) + o2.(float64)
		case AluSTRING:
			return o1.(string) + o2.(string)
		default:
			panic("unsupported")
		}
	}
	return nil
}

func Minus(o1 interface{}, o2 interface{}) interface{} {
	if !(o1 == nil || o2 == nil) {
		switch GetAluType(o1, o2) {
		case AluINTEGER:
			return o1.(int64) - o2.(int64)
		case AluFLOAT:
			return o1.(float64) - o2.(float64)
		default:
			panic("unsupported")
		}
	}
	return nil
}

func Mult(o1 interface{}, o2 interface{}) interface{} {
	if !(o1 == nil || o2 == nil) {
		switch GetAluType(o1, o2) {
		case AluINTEGER:
			return o1.(int64) * o2.(int64)
		case AluFLOAT:
			return o1.(float64) * o2.(float64)
		default:
			panic("unsupported")
		}
	}
	return nil
}

func Div(o1 interface{}, o2 interface{}) interface{} {
	if !(o1 == nil || o2 == nil) {
		switch GetAluType(o1, o2) {
		case AluINTEGER, AluFLOAT:
			if o2.(float64) == 0 {
				panic("分母不能为0")
			}
			return TrimNumber(o1.(float64)/o2.(float64), o1, o2)
		default:
			panic("unsupported")
		}
	}
	return nil
}

func Mod(o1 interface{}, o2 interface{}) interface{} {
	if !(o1 == nil || o2 == nil) {
		switch GetAluType(o1, o2) {
		case AluINTEGER:
			return o1.(int64) % o2.(int64)
		default:
			panic("unsupported")
		}
	}
	return nil
}

func Negative(o1 interface{}) interface{} {
	if !(o1 == nil) {
		switch GetAluTypeSingle(o1) {
		case AluINTEGER:
			return -o1.(int64)
		case AluFLOAT:
			return -o1.(float64)
		default:
			panic("unsupported")
		}
	}
	return nil
}

func TrimNumber(d float64, o1 interface{}, o2 interface{}) interface{} {
	switch o1.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return int64(d)
	case float32, float64:
		return d
	default:
		switch o2.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			return int64(d)
		case float32, float64:
			return d
		}
	}
	return d
}

type NumberType int

const (
	NumberTypeInt NumberType = iota
	NumberTypeUInt
	NumberTypeFloat
	NumberTypeBInt
	NumberTypeBFloat
	NumberTypeUnknown
)

func IsNumberType(o interface{}) NumberType {
	switch o.(type) {
	case int, int8, int16, int32, int64:
		return NumberTypeInt
	case uint, uint8, uint16, uint32, uint64:
		return NumberTypeUInt
	case float32, float64:
		return NumberTypeFloat
	case *big.Int:
		return NumberTypeBInt
	case *big.Float:
		return NumberTypeBFloat
	}
	return NumberTypeUnknown
}

func GetPrivateField(v reflect.Value, field string) interface{} {
	v = UnpackValue(v)
	v = v.FieldByName(field)
	v = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
	return v.Interface()
}
