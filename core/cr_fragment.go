package core

import (
	"fmt"
)

import (
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

type CRFragment struct {
	*Fragment

	cr []rune
}

func NewCRFragment(source *Source) *CRFragment {
	inst := &CRFragment{
		Fragment: NewFragment(source),
		cr:       source.ConsumeAndGetCR(source.lastCrSize),
	}
	return inst
}

func (f *CRFragment) GetScript() *estrings.StringBuilder {
	script := estrings.NewStringBuilder()
	varName := f.Source.GetParser().GetRandomeTextVarName()
	script.Append(fmt.Sprintf("<$%d>>", varName))
	script.Append(f.cr)
	f.Source.GetParser().GetTextVars()[varName] = string(f.cr)
	return script
}

func (f *CRFragment) ConsumeAndReturnNext() FragmentI {
	f.endLine = f.Source.CurLine - 1
	source := f.Source
	if !f.Source.IsEof() {
		if source.IsPlaceholderStart() {
			return NewPlaceholderFragment(source)
		} else if source.IsScriptStart() {
			return NewScriptBlockFragment(source)
		} else if source.IsHtmlTagStart() {
			return NewHtmlTagStartFragment(source)
		} else if source.IsHtmlTagEnd() {
			return NewHtmlTagEndFragment(source)
		} else if source.IsCrStart() {
			return NewCRFragment(source)
		} else {
			return NewTextFragment(source)
		}
	}
	return nil
}
