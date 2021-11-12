package core

import (
	"fmt"
	"strings"
)

import (
	"github.com/guangzhou-meta/beetl-go/util"
)

import (
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

type HtmlTagStartFragment struct {
	*ScriptFragment

	html     *HTMLTagContentParser
	appendCr bool
}

func NewHtmlTagStartFragment(source *Source) *HtmlTagStartFragment {
	inst := &HtmlTagStartFragment{
		ScriptFragment: NewScriptFragment(source),
	}
	return inst
}

func (f *HtmlTagStartFragment) GetScript() *estrings.StringBuilder {
	tagName := ""
	htmlTagStack := f.Source.htmlTagConfig.htmlTagStack
	html := f.html
	script := f.script
	defer func() {
		if err := recover(); err != nil {
			// TODO
		}
	}()
	if html.hasVarBinding && html.hasExportBinding {
		panic("不能同时有var 和 export ，只能选一个")
	}
	if html.hasVarBinding {
		script.Append("htmltagvar")
	} else if html.hasExportBinding {
		script.Append("htmltagexport")
	} else if html.hasRootExportBinding {
		script.Append("htmltagRootExport")
	} else {
		script.Append("htmltag")
	}
	tagName = html.GetTagName()
	script.Append("('").Append(tagName).Append("',")

	expMap := html.GetExpMap()
	quatMap := html.GetQuatMap()
	if len(expMap) != 0 {
		script.Append("{")
	}
	for key, value := range expMap {
		if util.Contains(html.crKey, key) {
			script.Append(cr1)
		}
		script.Append("\"").Append(key).Append("\"")
		attrValue := f.ParseAttr(quatMap[key], value)
		script.Append(attrValue).Append(",")
	}

	if len(expMap) != 0 {
		script.Append(html.GetHtmlColMapAsString())
		script.Append("}")
	} else {
		script.SetLength(script.Length() - 1)
	}

	if html.hasVarBinding || html.hasExportBinding || html.hasRootExportBinding {
		if len(expMap) == 0 {
			script.Append(",{}")
		}
		if len(strings.Trim(html.varBidingStr, " ")) == 0 {
			var defaultVarName string
			index := strings.LastIndex(tagName, ":")
			if index == -1 {
				defaultVarName = tagName
			} else {
				defaultVarName = tagName[index+1:]
			}
			script.Append(fmt.Sprintf(",'%s'", defaultVarName))
		} else {
			script.Append(fmt.Sprintf(",'%s'", html.varBidingStr))
		}
	}

	script.Append("){")
	if html.IsEmptyTag() {
		script.Append("}")
	} else {
		htmlTagStack.Push(tagName)
	}

	if f.appendCr {
		script.Append(cr1)
	}
	return script
}

func (f *HtmlTagStartFragment) ConsumeAndReturnNext() FragmentI {
	htmlTagBindingAttribute := f.Source.htmlTagConfig.htmlTagBindingAttribute
	f.html = NewHTMLTagContentParser(f.Source.GetParser().attributeNameConvert, f.Source.cs, f.Source.p, htmlTagBindingAttribute, true)
	f.html.Parser()
	f.Source.Move(f.html.index)
	f.endLine = f.startLine + len(f.html.crKey)
	return f.ScriptFragment.FindNext()
}

func (f *HtmlTagStartFragment) AppendCr() {
	f.appendCr = true
}

func (f *HtmlTagStartFragment) ParseAttr(q rune, attr string) string {
	phStart := f.Source.htmlTagConfig.phStart
	phEnd := f.Source.htmlTagConfig.phEnd
	sb := estrings.NewStringBuilder()
	start := 0
	end := 0
	index := strings.Index(attr[start:], phStart)
	qs := string(q)
	for index != -1 {
		holdStart := index
		end = strings.Index(attr[holdStart:], phEnd)
		for end != -1 {
			r := []rune(attr)
			if r[end-1] == '\\' {
				holdStart = end + 1
			} else {
				break
			}
			end = strings.Index(attr[holdStart:], phEnd)
		}
		if end == -1 {
			panic(attr + " 标签属性错误，有站位符号，但找不到到结束符号")
		}
		if index != 0 {
			sb.Append(qs + attr[start:index] + qs)
		}
		value := attr[index+len(phStart) : end]
		value = strings.ReplaceAll(value, "\\}", "}")
		sb.Append(fmt.Sprintf("(%s)+", value))
		start = end + len(phEnd)
		index = strings.Index(attr[start:], phStart)
	}
	if start == 0 {
		sb.Append(qs + attr + qs)
		return sb.String()
	}
	if start != len(attr) {
		sb.Append(qs + attr[start:] + qs)
	} else {
		sb.SetLength(sb.Length() - 1)
	}
	return sb.String()
}
