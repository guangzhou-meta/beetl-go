package core

import (
	collections "github.com/guangzhou-meta/go-lib/collections"
)

type HtmlTagConfig struct {
	htmlTagStart            string
	htmlTagEnd              string
	htmlTagBindingAttribute string

	start []rune
	end   []rune

	phStart string
	phEnd   string

	htmlTagStack *collections.Stack
	source       *Source
}

func NewHtmlTagConfig(htmlTagStart string, htmlTagEnd string, htmlTagBindingAttribute string) *HtmlTagConfig {
	inst := &HtmlTagConfig{
		htmlTagStart:            htmlTagStart,
		htmlTagEnd:              htmlTagEnd,
		htmlTagBindingAttribute: htmlTagBindingAttribute,
		start:                   []rune(htmlTagStart),
		end:                     []rune(htmlTagEnd),

		htmlTagStack: collections.NewStack(),

		phStart: "${",
		phEnd:   "}",
	}
	return inst
}

func (h *HtmlTagConfig) Init(source *Source) {
	h.source = source
	h.phStart = string(source.pd.start)
	h.phEnd = string(source.pd.end)
}

func (h *HtmlTagConfig) MatchTagStart() bool {
	return h.source.MatchAndSkip(h.start)
}

func (h *HtmlTagConfig) MatchTagEnd() bool {
	return h.source.MatchAndSkip(h.end)
}
