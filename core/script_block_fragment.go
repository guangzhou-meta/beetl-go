package core

import (
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

type ScriptBlockFragment struct {
	*ScriptFragment
}

func NewScriptBlockFragment(source *Source) *ScriptBlockFragment {
	inst := &ScriptBlockFragment{
		ScriptFragment: NewScriptFragment(source),
	}
	return inst
}

func (f *ScriptBlockFragment) GetScript() *estrings.StringBuilder {
	return f.script
}

func (f *ScriptBlockFragment) ConsumeAndReturnNext() FragmentI {
	for !f.Source.IsEof() && !f.Source.IsScriptEnd(f.script) {
		f.script.Append(f.Source.ConsumeAndGet())
	}
	f.SetEndLine()
	if f.Source.IsEof() {
		return nil
	}
	return f.FindNext()
}

func (f *ScriptBlockFragment) AppendCr() {
	f.script.Append(cr1)
}
