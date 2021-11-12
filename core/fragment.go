package core

import (
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

type FragmentStatus int32

const (
	normal FragmentStatus = iota
	del
	merge
)

type FragmentI interface {
	GetScript() *estrings.StringBuilder

	ConsumeAndReturnNext() FragmentI

	GetEndLine() int

	GetStatus() FragmentStatus

	SetStatus(status FragmentStatus)
}

type Fragment struct {
	Source    *Source
	startLine int
	endLine   int
	status    FragmentStatus
}

func NewFragment(source *Source) *Fragment {
	inst := &Fragment{
		Source: source,
		status: normal,
	}
	inst.SetStartLine()
	return inst
}

func (f *Fragment) FindNext() FragmentI {
	source := f.Source
	if source.IsEof() {
		return nil
	}
	if source.IsScriptStart() {
		f.SetEndLine()
		return NewScriptBlockFragment(source)
	}
	if source.IsPlaceholderStart() {
		f.SetEndLine()
		return NewPlaceholderFragment(source)
	}
	if source.IsHtmlTagStart() {
		f.SetEndLine()
		return NewHtmlTagStartFragment(source)
	}
	if source.IsHtmlTagEnd() {
		f.SetEndLine()
		return NewHtmlTagEndFragment(source)
	}
	if source.IsCrStart() {
		return NewCRFragment(source)
	}
	f.SetEndLine()
	return NewTextFragment(source)
}

func (f *Fragment) SetStartLine() {
	f.startLine = f.Source.CurLine
}

func (f *Fragment) GetEndLine() int {
	return f.endLine
}

func (f *Fragment) SetEndLine() {
	f.endLine = f.Source.CurLine
}

func (f *Fragment) GetStatus() FragmentStatus {
	return f.status
}

func (f *Fragment) SetStatus(status FragmentStatus) {
	f.status = status
}

func (f *Fragment) ConsumeAndReturnNext() FragmentI {
	panic("not implement")
}
func (f *Fragment) GetScript() *estrings.StringBuilder {
	panic("not implement")
}
