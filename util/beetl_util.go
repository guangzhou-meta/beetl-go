package util

import (
	"fmt"
)

const (
	AluNULL        = 0
	AluOBJECT      = 1
	AluSTRING      = 2
	AluDOUBLE      = 3
	AluFLOAT       = 4
	AluLONG        = 5
	AluINTEGER     = 6
	AluSHORT       = 7
	AluCHAR        = 8
	AluBIG_DECIMAL = 9
)

var (
	zhEnTokenMap map[string]string

	chars = []byte{
		// $,%,&,',(,),*,+,,,-, .,/, 0, 1, 2, 3, 4, 5, 6,
		36, 0, 0, 0, 0, 0, 0, 0, 0, 0, 46, 0, 48, 49, 50, 51, 52, 53, 54,
		// 7, 8, 9,:,;,<,=,>,?,@, A, B, C, D, E, F, G, H,
		55, 56, 57, 0, 0, 0, 0, 0, 0, 0, 65, 66, 67, 68, 69, 70, 71, 72,
		// I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W,
		73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
		// X, Y, Z,[,\,],^, _,`, a, b, c,  d,  e,  f,  g,
		88, 89, 90, 0, 0, 0, 0, 95, 0, 97, 98, 99, 100, 101, 102, 103,
		//  h,  i,  j,  k,  l,  m,  n,  o,  p,  q,  r,
		104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114,
		//  s,  t,  u,  v,  w,  x,  y,  z
		115, 116, 117, 118, 119, 120, 121, 122,
	}

	CheckResult = make([]int, 2)
)

type ComparableI interface {
	CompareTo(o ComparableI) int
}

func init() {
	zhEnTokenMap = make(map[string]string)
	initZhEnTokenMap()
}

func initZhEnTokenMap() {
	zhEnTokenMap["，"] = ","
	zhEnTokenMap["；"] = ";"
	zhEnTokenMap["）"] = ")"
	zhEnTokenMap["（"] = "("
	zhEnTokenMap["‘"] = "'"
	zhEnTokenMap["’"] = "'"
	zhEnTokenMap["“"] = "\""
	zhEnTokenMap["”"] = "\""
	zhEnTokenMap["。"] = "."
	zhEnTokenMap["｝"] = "}"
	zhEnTokenMap["｛"] = "{"
	zhEnTokenMap["＝"] = "="
	zhEnTokenMap["！"] = "!"
	zhEnTokenMap["％"] = "%"
	zhEnTokenMap["／"] = "/"
	zhEnTokenMap["＼"] = "\\"
	zhEnTokenMap["．"] = "."
}

func ReportChineseTokenError(msg string) string {
	maybeEnChar, ok := zhEnTokenMap[msg]
	if ok {
		return fmt.Sprintf("%s 貌似输入了中文符号,应该是 %s", msg, maybeEnChar)
	}
	return msg
}

func CheckNaming(str string) bool {
	strArray := []rune(str)
	length := len(strArray)
	if str == "" || length == 0 {
		return false
	}
	index := 0
	word := strArray[index]
	index++
	if word >= 46 && word <= 57 {
		CheckResult[0] = 1
		CheckResult[1] = int(word)
	} else if strArray[length-1] == 46 {
		CheckResult[0] = length
		CheckResult[1] = 46
	} else {
		for {
			if word < 36 || word > 122 || chars[word-36] == 0 {
				CheckResult[0] = index + 1
				CheckResult[1] = int(word)
				return false
			}
			if index == length {
				return true
			}
			word = strArray[index]
			index++
		}
	}
	return false
}
