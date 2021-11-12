package core

type ScriptFragmentI interface {
	AppendCr()
}

type ScriptFragment struct {
	*BeetlFragment
}

func NewScriptFragment(source *Source) *ScriptFragment {
	inst := &ScriptFragment{
		BeetlFragment: NewBeetlFragment(source),
	}
	return inst
}

func (f *ScriptFragment) AppendCr() {
	f.script.Append(cr1)
}
