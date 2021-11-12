package core

import (
	"fmt"
	"io"
)

import (
	"github.com/guangzhou-meta/beetl-go/util"
)

import (
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

const (
	cr1 rune = '\n'
)

type TextParser struct {
	textVars             map[int]string
	textNameSuffix       int
	pd                   *PlaceHolderDelimeter
	sd                   *ScriptDelimeter
	source               *Source
	script               *estrings.StringBuilder
	sourceFragment       *SourceFragment
	textCr               string
	attributeNameConvert AttributeNameConvertI
}

func NewTextParser(groupTemplate *GroupTemplate, pdConfig *DelimeterHolder, sdConfig *DelimeterHolder) *TextParser {
	inst := &TextParser{
		pd:             pdConfig.CreatePhd(),
		sd:             sdConfig.CreateSd(),
		sourceFragment: NewSourceFragment(),
		script:         estrings.NewStringBuilder(),
		textVars:       make(map[int]string),
	}
	if groupTemplate == nil {
		inst.attributeNameConvert = NewDefaultAttributeNameConvert()
	} else {
		inst.attributeNameConvert = groupTemplate.GetHtmlTagAttrNameConvert()
	}
	return inst
}

func (p *TextParser) DoParse(orginal io.Reader) {
	p.Scan1(orginal)

	err := p.sourceFragment.Merge()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range p.sourceFragment.List {
		if f.GetStatus() == del {
			continue
		}
		p.script.Append(f.GetScript().String())
	}

	cr := p.source.FindCr()
	if cr != nil {
		p.textCr = *cr
	} else {
		p.textCr = util.GetLineSeparator()
	}
}

func (p *TextParser) Scan1(orginal io.Reader) {
	cs, err := io.ReadAll(orginal)
	if err != nil {
		panic(err)
	}
	csa := []rune(string(cs))
	p.source = NewSource(csa)
	p.source.Init(p, p.pd, p.sd)
	var test FragmentI = NewTextFragment(p.source)
	var next FragmentI
	for {
		next = test.ConsumeAndReturnNext()
		p.sourceFragment.Add(test)
		if next == nil {
			break
		}
		test = next
	}
}

func (p *TextParser) GetScript() *estrings.StringBuilder {
	return p.script
}

func (p *TextParser) GetTextVars() map[int]string {
	return p.textVars
}

func (p *TextParser) GetRandomeTextVarName() int {
	r := p.textNameSuffix
	p.textNameSuffix++
	return r
}

func (p *TextParser) GetTextCr() string {
	return p.textCr
}
