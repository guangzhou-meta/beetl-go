package core

import (
	"strings"
)

import (
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

type DefaultAttributeNameConvert struct {
}

func NewDefaultAttributeNameConvert() *DefaultAttributeNameConvert {
	return &DefaultAttributeNameConvert{}
}

func (*DefaultAttributeNameConvert) Convert(origin string) string {
	cs := []rune(origin)
	sb := estrings.NewStringBuilder()
	upper := false
	for _, c := range cs {
		if upper {
			if c == '-' {
				continue
			}
			sb.Append(strings.ToUpper(string(c)))
			upper = false
		} else {
			if c == '-' {
				upper = true
			} else {
				sb.Append(c)
			}
		}
	}
	return sb.String()
}
