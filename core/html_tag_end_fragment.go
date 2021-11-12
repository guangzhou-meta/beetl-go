package core

import (
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

type HtmlTagEndFragment struct {
	*ScriptFragment

	html     *HTMLTagContentParser
	appendCr bool
}

func NewHtmlTagEndFragment(source *Source) *HtmlTagEndFragment {
	inst := &HtmlTagEndFragment{
		ScriptFragment: NewScriptFragment(source),
	}
	return inst
}

func (f *HtmlTagEndFragment) GetScript() *estrings.StringBuilder {
	tagName := ""
	htmlTagStack := f.Source.htmlTagConfig.htmlTagStack
	html := f.html
	script := f.script
	defer func() {
		if err := recover(); err != nil {
			// TODO
		}
	}()
	tagName = html.GetTagName()
	if htmlTagStack.Size() == 0 {
		panic("解析html tag出错")
	}
	lastTag := (htmlTagStack.Peek()).(string)
	if tagName == lastTag {
		htmlTagStack.Pop()
		script.Append("}")
	} else {
		panic("解析html tag出错,期望匹配标签" + lastTag)
	}
	if f.appendCr {
		script.Append(cr1)
	}
	return script
}

func (f *HtmlTagEndFragment) ConsumeAndReturnNext() FragmentI {
	htmlTagBindingAttribute := f.Source.htmlTagConfig.htmlTagBindingAttribute
	f.html = NewHTMLTagContentParser(f.Source.GetParser().attributeNameConvert, f.Source.cs, f.Source.p, htmlTagBindingAttribute, false)
	f.html.Parser()
	f.Source.Move(f.html.index)
	f.endLine = f.startLine
	return f.ScriptFragment.FindNext()
}

func (f *HtmlTagEndFragment) AppendCr() {
	f.appendCr = true
}
