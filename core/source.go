package core

import (
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

type Source struct {
	cs               []rune
	p                int
	size             int
	pd               *PlaceHolderDelimeter
	sd               *ScriptDelimeter
	htmlTagConfig    *HtmlTagConfig
	parser           *TextParser
	CurLine          int
	isSupportHtmlTag bool
	lastCrSize       int
	lastTextFragment *TextFragment
	beetlFragment    *BeetlFragment
}

func NewSource(cs []rune) *Source {
	inst := &Source{
		cs:   cs,
		p:    0,
		size: len(cs),
	}
	return inst
}

func (s *Source) Init(parser *TextParser, pd *PlaceHolderDelimeter, sd *ScriptDelimeter) {
	s.pd = pd
	s.sd = sd
	pd.SetSource(s)
	sd.SetSource(s)
	s.parser = parser
}

func (s *Source) FindCr() *string {
	s.p = 0
	for !s.IsEof() {
		crCount := s.IsMatchCR()
		if crCount != 0 {
			res := string(s.ConsumeAndGetCR(crCount))
			return &res
		}
		s.Consume()
	}
	return nil
}

func (s *Source) IsPlaceholderStart() bool {
	return s.pd.MatchStart()
}

func (s *Source) IsPlaceHolderEnd(script *estrings.StringBuilder) bool {
	return s.pd.MatchEnd(script)
}

func (s *Source) IsScriptStart() bool {
	return s.sd.MatchStart()
}

func (s *Source) IsScriptEnd(script *estrings.StringBuilder) bool {
	return s.sd.MatchEnd(script)
}

func (s *Source) IsHtmlTagStart() bool {
	return s.isSupportHtmlTag && s.htmlTagConfig.MatchTagStart()
}

func (s *Source) IsHtmlTagEnd() bool {
	return s.isSupportHtmlTag && s.htmlTagConfig.MatchTagEnd()
}

func (s *Source) MatchAndSkip(text []rune) bool {
	isMatch := s.IsMatch(text) && !s.HasEscape()
	if isMatch {
		s.ConsumeX(len(text))
	}
	return isMatch
}

func (s *Source) MatchAndSkipEnd(text []rune) bool {
	isMatch := s.IsMatch(text) && !s.HasScriptEscape()
	if isMatch {
		s.ConsumeX(len(text))
	}
	return isMatch
}

func (s *Source) AddLine() {
	s.CurLine++
}

func (s *Source) HasEscape() bool {
	if s.p <= 0 {
		return false
	}
	if s.cs[s.p-1] != '\\' {
		return false
	}
	if s.p <= 1 {
		s.RemoveTextEscape()
		return true
	}
	if s.cs[s.p-2] == '\\' {
		s.RemoveTextEscape()
		return false
	}
	s.RemoveTextEscape()
	return true
}

func (s *Source) HasScriptEscape() bool {
	if s.p <= 0 {
		return false
	}
	if s.cs[s.p-1] != '\\' {
		return false
	}
	if s.p <= 1 {
		s.RemoveScriptEscape()
		return true
	}
	if s.cs[s.p-2] == '\\' {
		s.RemoveScriptEscape()
		return false
	}
	s.RemoveScriptEscape()
	return true
}

func (s *Source) RemoveScriptEscape() {
	s.beetlFragment.RemoveEscape()
}

func (s *Source) RemoveTextEscape() {
	s.lastTextFragment.RemoveTextEscape()
}

func (s *Source) IsMatch(str []rune) bool {
	cur := s.p
	for _, c := range str {
		if cur < s.size && s.cs[cur] == c {
			cur++
		} else {
			return false
		}
	}
	return true
}

func (s *Source) IsCrStart() bool {
	crLength := s.IsMatchCR()
	if crLength == 0 {
		return false
	}
	s.lastCrSize = crLength
	return true
}

func (s *Source) IsMatchCR() int {
	c := s.cs[s.p]
	if c == '\n' {
		return 1
	}
	if c == '\r' {
		if s.size > s.p+1 &&
			s.cs[s.p+1] == '\n' {
			return 2
		}
		return 1
	}
	return 0
}

func (s *Source) ConsumeAndGet() rune {
	c := s.Get()
	s.Consume()
	return c
}

func (s *Source) ConsumeAndGetCR(size int) []rune {
	cs := make([]rune, size)
	for i := 0; i < size; i++ {
		cs[i] = s.ConsumeAndGet()
	}
	s.AddLine()
	return cs
}

func (s *Source) Consume() {
	s.p++
}

func (s *Source) ConsumeX(x int) {
	s.p = s.p + x
}

func (s *Source) Get() rune {
	return s.cs[s.p]
}

func (s *Source) IsEof() bool {
	return s.p == s.size
}

func (s *Source) Move(p int) {
	s.p = p
}

func (s *Source) GetParser() *TextParser {
	return s.parser
}

func (s *Source) SetParser(parser *TextParser) {
	s.parser = parser
}
