package core

import (
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

type PlaceholderFragment struct {
	*BeetlFragment
}

func NewPlaceholderFragment(source *Source) *PlaceholderFragment {
	inst := &PlaceholderFragment{
		BeetlFragment: NewBeetlFragment(source),
	}
	inst.script.Append("<<")
	inst.SetEndLine()
	return inst
}

func (b *PlaceholderFragment) GetScript() *estrings.StringBuilder {
	return b.script
}

func (b *PlaceholderFragment) ConsumeAndReturnNext() FragmentI {
	script := estrings.NewStringBuilder()
	if b.Source.pd.isMatchFirstGroup {
		script.Append("<<")
	} else {
		script.Append("<#")
	}
	b.script = script

	for !b.Source.IsEof() {
		if b.Source.IsPlaceHolderEnd(script) {
			script.Append(">>")
			break
		} else {
			script.Append(b.Source.ConsumeAndGet())
		}
	}
	if b.Source.IsEof() {
		return nil
	}
	return b.FindNext()
}
