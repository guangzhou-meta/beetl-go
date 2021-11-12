package util

import (
	"bufio"
	"io"
	"os"
	"reflect"
	"runtime"
	"strings"
)

func Contains(src interface{}, target interface{}) bool {
	srcValue := reflect.ValueOf(src)
	switch reflect.TypeOf(src).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < srcValue.Len(); i++ {
			if srcValue.Index(i).Interface() == target {
				return true
			}
		}
	case reflect.Map:
		if srcValue.MapIndex(reflect.ValueOf(target)).IsValid() {
			return true
		}
	case reflect.String:
		return strings.Contains(src.(string), target.(string))
	}
	return false
}

func IndexOf(src interface{}, target interface{}) int {
	return IndexOfFrom(src, target, 0)
}

func IndexOfFrom(src interface{}, target interface{}, from int) int {
	srcValue := reflect.ValueOf(src)
	switch reflect.TypeOf(src).Kind() {
	case reflect.Slice, reflect.Array:
		srcLength := srcValue.Len()
		if from >= srcLength {
			panic("IndexOf的fromIndex超出数组长度")
		}
		if from < 0 {
			from = 0
		}
		for i := from; i < srcLength; i++ {
			if srcValue.Index(i).Interface() == target {
				return i
			}
		}
		break
	case reflect.String:
		return strings.Index(src.(string), target.(string))
	default:
		panic("IndexOf只支持Slice,String类型")
	}
	return -1
}

func LastIndexOf(src interface{}, target interface{}) int {
	return LastIndexOfFrom(src, target, -1)
}

func LastIndexOfFrom(src interface{}, target interface{}, from int) int {
	srcValue := reflect.ValueOf(src)
	switch reflect.TypeOf(src).Kind() {
	case reflect.Slice, reflect.Array:
		srcLength := srcValue.Len()
		if from >= srcLength {
			panic("LastIndexOf的fromIndex超出数组长度")
		}
		if from < 0 {
			from = srcLength - 1
		}
		for i := from; i >= 0; i++ {
			if srcValue.Index(i).Interface() == target {
				return i
			}
		}
		break
	case reflect.String:
		return strings.LastIndex(src.(string), target.(string))
	default:
		panic("LastIndexOf只支持Slice,String类型")
	}
	return -1
}

func GetLineSeparator() string {
	sys := runtime.GOOS
	switch sys {
	case "linux", "unix":
		return "\n"
	case "windows":
		return "\r\n"
	}
	return "\r"
}

func UnpackType(t reflect.Type) reflect.Type {
	for reflect.Ptr == t.Kind() {
		t = t.Elem()
	}
	return t
}

func UnpackValue(t reflect.Value) reflect.Value {
	for reflect.Ptr == t.Kind() {
		t = t.Elem()
	}
	return t
}

func LoadProperty(path string) (config map[string]string, err error) {
	var f *os.File
	f, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	config = make(map[string]string)
	reader := bufio.NewReader(f)
	var buf []byte
	for {
		buf, _, err = reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
			config = nil
			return
		}
		str := strings.TrimSpace(string(buf))
		splitIndex := strings.Index(str, "=")
		if splitIndex <= 0 {
			continue
		}
		key := strings.TrimSpace(str[:splitIndex])
		if len(key) == 0 ||
			strings.HasPrefix(key, "#") ||
			strings.HasPrefix(key, "//") ||
			strings.HasPrefix(key, "/*") {
			continue
		}
		value := strings.TrimSpace(str[splitIndex+1:])
		config[key] = value
	}
	return
}
