package core

import (
	"fmt"
)

import (
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

type TextFragment struct {
	*Fragment

	text    *estrings.StringBuilder
	crCount int
}

func NewTextFragment(source *Source) *TextFragment {
	inst := &TextFragment{
		Fragment: NewFragment(source),
		text:     estrings.NewStringBuilder(),
	}
	inst.endLine = inst.startLine
	source.lastTextFragment = inst
	return inst
}

func (t *TextFragment) AppendTextFragment(fr FragmentI) error {
	switch v := fr.(type) {
	case *TextFragment:
		t.text.Append(v.text.String())
	case *CRFragment:
		t.text.Append(v.cr)
		t.crCount++
	default:
		return fmt.Errorf("无效参数 %v", fr)
	}
	return nil
}

func (t *TextFragment) GetScript() *estrings.StringBuilder {
	script := estrings.NewStringBuilder()
	if t.text.Length() == 0 {
		return script
	}
	varName := t.Source.GetParser().GetRandomeTextVarName()
	script.Append("<$").Append(varName).Append(">>")
	for i := 0; i < t.crCount; i++ {
		script.Append(cr1)
	}
	t.Source.GetParser().GetTextVars()[varName] = t.text.String()
	return script
}

func (t *TextFragment) ConsumeAndReturnNext() FragmentI {
	source := t.Source
	for !source.IsEof() {
		if source.IsPlaceholderStart() {
			t.SetEndLine()
			return NewPlaceholderFragment(source)
		}
		if source.IsScriptStart() {
			t.SetEndLine()
			return NewScriptBlockFragment(source)
		}
		if source.IsHtmlTagStart() {
			t.SetEndLine()
			return NewHtmlTagStartFragment(source)
		}
		if source.IsHtmlTagEnd() {
			t.SetEndLine()
			return NewHtmlTagEndFragment(source)
		}
		if source.IsCrStart() {
			return NewCRFragment(source)
		}
		c := source.ConsumeAndGet()
		t.text.Append(c)
	}
	return nil
}

func (t *TextFragment) OnlySpace() bool {
	for i := 0; i < t.text.Length(); i++ {
		c := t.text.CharAt(i)
		if c != ' ' && c != '\t' {
			return false
		}
	}
	return true
}

func (t *TextFragment) RemoveTextEscape() {
	t.text.SetLength(t.text.Length() - 1)
}
