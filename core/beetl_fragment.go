package core

import (
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

type BeetlFragmentI interface {
	GetBeetlFragment() BeetlFragmentI
}

type BeetlFragment struct {
	*Fragment

	script *estrings.StringBuilder
}

func NewBeetlFragment(source *Source) *BeetlFragment {
	inst := &BeetlFragment{
		Fragment: NewFragment(source),
		script:   estrings.NewStringBuilder(),
	}
	source.beetlFragment = inst
	return inst
}

func (b *BeetlFragment) RemoveEscape() {
	b.script.SetLength(b.script.Length() - 1)
}

func (b *BeetlFragment) GetScript() *estrings.StringBuilder {
	return b.script
}

func (b *BeetlFragment) ConsumeAndReturnNext() FragmentI {
	return nil
}

func (b *BeetlFragment) GetBeetlFragment() BeetlFragmentI {
	return b
}
