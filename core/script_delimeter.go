package core

import (
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

type ScriptDelimeter struct {
	*Delimeter

	endIsCr bool
}

func NewScriptDelimeter(start []rune, end []rune, start1 *[]rune, end1 *[]rune) *ScriptDelimeter {
	inst := &ScriptDelimeter{
		Delimeter: NewDelimeter(start, end, start1, end1),
	}
	return inst
}

func (d *ScriptDelimeter) MatchStart() bool {
	match := d.Delimeter.MatchStart()
	if match {
		d.endIsCr = false
	}
	return match
}

func (d *ScriptDelimeter) MatchEnd(script *estrings.StringBuilder) bool {
	if d.isMatchFirstGroup {
		return d.MatchEnd1(d.end, script)
	}
	return d.hasTwoGroup && d.MatchEnd1(d.end1, script)
}

func (d *ScriptDelimeter) MatchEnd1(end []rune, script *estrings.StringBuilder) bool {
	if end != nil && len(end) > 0 {
		return d.source.MatchAndSkipEnd(end)
	}
	matchCount := d.source.IsMatchCR()
	if matchCount != 0 {
		d.endIsCr = true
		d.source.ConsumeX(matchCount)
		d.source.AddLine()
		script.Append(cr1)
		return true
	}
	return false
}
