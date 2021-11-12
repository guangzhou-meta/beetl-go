package core

import (
	"strings"
)

import (
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

var (
	EntTag  rune = '>'
	EntTags      = []rune{'/', '>'}
)

type HTMLTagContentParser struct {
	attributeNameConvert AttributeNameConvertI
	cs                   []rune
	index                int
	isStart              bool
	ts                   int
	te                   int
	status               int
	tagName              string
	varBidingStr         string
	lastKey              string
	_var                 string
	_export              string
	_rootExport          string
	isEmptyTag           bool
	hasVarBinding        bool
	hasExportBinding     bool
	hasRootExportBinding bool
	expMap               map[string]string
	quatMap              map[string]rune
	htmlAttributeNameMap map[string]string
	crKey                []string
}

func NewHTMLTagContentParser(attributeNameConvert AttributeNameConvertI, cs []rune, index int, bindingAttr string, isStart bool) *HTMLTagContentParser {
	inst := &HTMLTagContentParser{
		attributeNameConvert: attributeNameConvert,
		cs:                   cs,
		expMap:               make(map[string]string),
		quatMap:              make(map[string]rune),
		htmlAttributeNameMap: make(map[string]string),
		crKey:                make([]string, 0),
	}
	inst.parseBindingAttr(bindingAttr)
	inst.index = index
	inst.isStart = isStart
	inst.ts = index
	inst.te = index
	return inst
}

func (p *HTMLTagContentParser) parseBindingAttr(bindingAttr string) {
	bindConf := strings.Split(bindingAttr, ",")
	if len(bindConf) == 1 {
		p._var = bindConf[0]
		p._export = "export"
		p._rootExport = "$export$"
	} else {
		p._var = bindConf[0]
		p._export = bindConf[1]
		p._rootExport = "$" + p._export
	}
}

func (p *HTMLTagContentParser) Parser() {
	if p.isStart {
		p.parserStart()
	} else {
		p.parserEnd()
	}
}

func (p *HTMLTagContentParser) parserStart() {
	p.findTagName()

	p.findAttrs()
	p.findBindingFlag()
	if p.status != -1 {
		p.findVars()
	}

	p.endTag()
}

func (p *HTMLTagContentParser) parserEnd() {
	p.findTagName()
	if p.match('>') {
		p.move(1)
	} else {
		panic(p.tagName + "结束标签格式错")
	}
}

func (p *HTMLTagContentParser) findTagName() {
	p.idToken()
	if p.status == -1 {
		panic("非法标签名")
	}
	tagSb := estrings.NewStringBuilder()
	tagSb.Append(p.attributeNameConvert.Convert(p.subString()))
	p.t_consume()
	for p.match(':') {
		p.move(1)
		p.idToken()
		if p.status == -1 {
			panic("非法标签名")
		}
		tagSb.Append(":").Append(p.subString())
		p.t_consume()
	}
	p.tagName = tagSb.String()
}

func (p *HTMLTagContentParser) match(c rune) bool {
	return p.cs[p.index] == c
}

func (p *HTMLTagContentParser) matchCR() bool {
	return p.index < len(p.cs) && (p.cs[p.index] == '\r' || p.cs[p.index] == '\n')
}

func (p *HTMLTagContentParser) matchArr(expected []rune) bool {
	i := 0
	for p.index+i < len(p.cs) && i < len(expected) {
		if p.cs[p.index+i] != expected[i] {
			return false
		}
		i++
	}
	return i == len(expected)
}

func (p *HTMLTagContentParser) findAttrs() {
	p.findAttr()
	for p.status != -1 {
		if p.match(' ') || p.matchCR() {
			p.findAttr()
		} else {
			return
		}
	}
	p.t_recover()
}

func (p *HTMLTagContentParser) findAttr() {
	findCR := p.stripSpaceAndCR()
	p.idToken()
	if p.status == -1 {
		return
	}
	colName := p.subString()

	lastKey := p.attributeNameConvert.Convert(colName)
	p.htmlAttributeNameMap[lastKey] = colName

	p.t_consume()
	p.stripSpace()
	if p.match('=') {
		p.move(1)
		isSingleQuat := p.strToken()
		value := p.subString()
		p.t_consume()
		p.move(1)
		if lastKey == p._var {
			p.hasVarBinding = true
			p.varBidingStr = value
			return
		} else if lastKey == p._export {
			p.hasExportBinding = true
			p.varBidingStr = value
			return
		} else if lastKey == p._rootExport {
			p.hasRootExportBinding = true
			p.varBidingStr = value
			return
		}
		if isSingleQuat {
			p.quatMap[lastKey] = '\''
		} else {
			p.quatMap[lastKey] = '"'

		}
		p.expMap[lastKey] = value
		if findCR {
			p.crKey = append(p.crKey, lastKey)
		}
	} else {
		panic("没有找到属性")
	}
}

func (p *HTMLTagContentParser) findBindingFlag() {
	p.stripSpace()
	if !p.match(';') {
		p.status = -1
		return
	}
	p.move(1)
	p.hasVarBinding = true
}

func (p *HTMLTagContentParser) findVars() {
	p.stripSpace()
	p.idToken()
	sb := estrings.NewStringBuilder()
	for p.status != -1 {
		sb.Append(p.subString())
		p.t_consume()
		p.stripSpace()
		if p.match(',') {
			p.move(1)
			p.stripSpace()
			p.idToken()
			sb.Append(",")
		} else {
			break
		}
	}
	p.t_consume()
	if sb.Length() != 0 {
		sb.SetLength(sb.Length())
		p.varBidingStr = sb.String()
	}
}

func (p *HTMLTagContentParser) endTag() {
	p.stripSpace()
	if p.match(EntTag) {
		p.move(1)
		p.isEmptyTag = false
	} else if p.matchArr(EntTags) {
		p.isEmptyTag = true
		p.move(2)
	} else {
		illegal := p.cs[p.index]
		if illegal == '\r' || illegal == '\n' {
			panic("标签未正确结束:" + ",碰到换行符号")
		} else {
			panic("标签未正确结束:" + ",碰到非法符号'" + string(p.cs[p.index]) + "'")
		}
	}
}

func (p *HTMLTagContentParser) strToken() bool {
	p.stripSpace()
	if p.match('\'') {
		p.move(1)
		p.findOneChar('\'')
		if p.status == -1 {
			panic("错误的属性，缺少'")
		}
		return true
	} else if p.match('"') {
		p.move(1)
		p.findOneChar('"')
		if p.status == -1 {
			panic("错误的属性,缺少'")
		}
		return false
	} else {
		panic("属性必须使用双引号或者单引号")
	}
}

func (p *HTMLTagContentParser) idToken() {
	if p.ts > len(p.cs) {
		panic("解析错误")
	}
	c := p.cs[p.ts]
	if p.isID(c) {
		i := 1
		for p.ts < len(p.cs) {
			c = p.cs[p.ts+i]
			if p.isID(c) || p.isDigit(c) {
				i++
				continue
			} else {
				break
			}
		}
		p.t_forword(i)
	} else {
		p.status = -1
	}
}

func (p *HTMLTagContentParser) stripSpaceAndCR() bool {
	p.ts = p.index
	i := 0
	findCR := false
	for p.ts < len(p.cs) {
		c := p.cs[p.ts+i]
		if c == ' ' || c == '\t' {
			i++
			continue
		} else if c == '\n' || c == '\r' {
			i++
			findCR = true
		} else {
			break
		}
	}
	p.ts = p.ts + i
	p.te = p.ts
	p.index = p.te
	return findCR
}

func (p *HTMLTagContentParser) stripSpace() {
	p.ts = p.index
	i := 0
	for p.ts < len(p.cs) {
		c := p.cs[p.ts+i]
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			i++
			continue
		} else {
			break
		}
	}
	p.ts = p.ts + i
	p.te = p.ts
	p.index = p.te
}

func (p *HTMLTagContentParser) t_forword(forward int) {
	p.te = p.ts + forward
}

func (p *HTMLTagContentParser) t_consume() {
	p.index = p.te
	p.ts = p.te
	p.status = 0
}

func (p *HTMLTagContentParser) t_recover() {
	p.ts = p.index
	p.te = p.ts
}

func (p *HTMLTagContentParser) move(i int) {
	p.index = p.index + i
	p.te = p.index
	p.ts = p.te
	p.status = 0
}

func (p *HTMLTagContentParser) subString() string {
	return string(p.cs[p.ts:p.te])
}

func (p *HTMLTagContentParser) findOneChar(c rune) {
	i := 0
	for (p.ts + i) < len(p.cs) {
		ch := p.cs[p.ts+i]
		if ch != c {
			i++
			if ch == '\n' || ch == '\r' {
				p.status = -1
				p.t_recover()
				return
			}
		} else {
			p.t_forword(i)
			return
		}
	}
	p.status = -1
	p.t_recover()
}

func (p *HTMLTagContentParser) isID(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_' || c == '-' || c == '$'
}

func (p *HTMLTagContentParser) isDigit(c rune) bool {
	return c > '0' && c < '9'
}

func (p *HTMLTagContentParser) GetHtmlColMapAsString() string {
	if len(p.htmlAttributeNameMap) == 0 {
		return "$cols:{}"
	}
	sb := estrings.NewStringBuilderStr("$cols:{")
	for varName, colName := range p.htmlAttributeNameMap {
		sb.Append("'").Append(varName).Append("':'").Append(colName).Append("',")
	}

	sb.SetCharAt(sb.Length()-1, '}')
	return sb.String()
}

func (p *HTMLTagContentParser) IsEmptyTag() bool {
	return p.isEmptyTag
}

func (p *HTMLTagContentParser) GetIndex() int {
	return p.index
}

func (p *HTMLTagContentParser) GetTagName() string {
	return p.tagName
}

func (p *HTMLTagContentParser) GetExpMap() map[string]string {
	return p.expMap
}

func (p *HTMLTagContentParser) GetQuatMap() map[string]rune {
	return p.quatMap
}
