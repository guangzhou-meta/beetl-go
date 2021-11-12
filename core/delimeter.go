package core

import (
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

type Delimeter struct {
	source            *Source
	start             []rune
	end               []rune
	start1            []rune
	end1              []rune
	isMatchFirstGroup bool
	hasTwoGroup       bool
}

func NewDelimeter(start []rune, end []rune, start1 *[]rune, end1 *[]rune) *Delimeter {
	inst := &Delimeter{
		start: start,
		end:   end,
	}
	if start1 != nil && len(*start1) > 0 {
		inst.start1 = *start1
		inst.hasTwoGroup = true
	}
	if end1 != nil {
		inst.end1 = *end1
	}
	return inst
}

func (d *Delimeter) MatchStart() bool {
	match := d.source.MatchAndSkip(d.start)
	if !match && d.hasTwoGroup {
		match = d.source.MatchAndSkip(d.start1)
		if match {
			d.isMatchFirstGroup = false
		}
	} else {
		d.isMatchFirstGroup = true
	}
	//fmt.Println(string(d.source.cs[0:d.source.p]), fmt.Sprintf("[ddd = %v]",match))
	return match
}

func (d *Delimeter) MatchEnd(script *estrings.StringBuilder) bool {
	if d.isMatchFirstGroup {
		return d.source.MatchAndSkipEnd(d.end)
	}
	if d.hasTwoGroup {
		return d.source.MatchAndSkipEnd(d.end1)
	}
	return false
}

func (d *Delimeter) GetSource() *Source {
	return d.source
}

func (d *Delimeter) SetSource(source *Source) {
	d.source = source
}
