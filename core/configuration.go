package core

import (
	"fmt"
	"strconv"
	"strings"
)

import (
	"github.com/guangzhou-meta/beetl-go/util"
)

const (
	DelimiterPlaceholderStart  = "DELIMITER_PLACEHOLDER_START"
	DelimiterPlaceholderEnd    = "DELIMITER_PLACEHOLDER_END"
	DelimiterStatementStart    = "DELIMITER_STATEMENT_START"
	DelimiterStatementEnd      = "DELIMITER_STATEMENT_END"
	DelimiterPlaceholderStart2 = "DELIMITER_PLACEHOLDER_START2"
	DelimiterPlaceholderEnd2   = "DELIMITER_PLACEHOLDER_END2"
	DelimiterStatementStart2   = "DELIMITER_STATEMENT_START2"
	DelimiterStatementEnd2     = "DELIMITER_STATEMENT_END2"
	NativeCall                 = "NATIVE_CALL"
	IgnoreClientIoError        = "IGNORE_CLIENT_IO_ERROR"
	DirectByteOutput           = "DIRECT_BYTE_OUTPUT"
	TemplateRoot               = "TEMPLATE_ROOT"
	TemplateCharset            = "TEMPLATE_CHARSET"
	ErrorHandler               = "ERROR_HANDLER"
	MvcStrict                  = "MVC_STRICT"
	WebappExt                  = "WEBAPP_EXT"
	HtmlTagSupport             = "HTML_TAG_SUPPORT"
	HtmlTagFlag                = "HTML_TAG_FLAG"
	HtmlTagAttrConvert         = "HTML_TAG_ATTR_CONVERT"
	ImportPackage              = "IMPORT_PACKAGE"
	ENGINE                     = "ENGINE"
	NativeSecuartyManager      = "NATIVE_SECUARTY_MANAGER"
	ResourceLoader             = "RESOURCE_LOADER"
	HtmlTagBindingAttribute    = "HTML_TAG_BINDING_ATTRIBUTE"
	BufferSize                 = "buffer.maxSize"
	BufferNum                  = "buffer.num"
	SafeOutput                 = "SAFE_OUTPUT"
	DelimeterConfig            = "DELIMETER_CONFIG"
	CACHE                      = "CACHE"
)

type Configuration struct {
	charset                 string
	placeholderStart        string
	placeholderEnd          string
	placeholderStart2       *string
	placeholderEnd2         *string
	statementStart          string
	statementEnd            string
	statementStart2         *string
	statementEnd2           *string
	SafeOutput              bool
	nativeCall              bool
	directByteOutput        bool
	isStrict                bool
	isIgnoreClientIOError   bool
	htmlTagBindingAttribute string
	htmlTagAttributeConvert string
	pkgList                 map[string]interface{}
	webAppExt               string
	hasFunctionLimiter      bool
	functionLimiterStart    string
	functionLimiterEnd      string
	delimeterClass          *string

	ps                         map[string]string
	resourceMap                map[string]string
	fnMap                      map[string]string
	fnPkgMap                   map[string]string
	formatMap                  map[string]string
	virtualClass               map[string]string
	tagFactoryMap              map[string]string
	tagMap                     map[string]string
	defaultFormatMap           map[string]string
	generalVirtualAttributeSet map[string]interface{}

	pd *DelimeterHolder
	sd *DelimeterHolder
}

func NewConfiguration(customPropertiesFilePath *string) *Configuration {
	placeholderStart := "${"
	placeholderEnd := "}"
	statementStart := "<%"
	statementEnd := "%>"
	inst := &Configuration{
		placeholderStart:           placeholderStart,
		placeholderEnd:             placeholderEnd,
		placeholderStart2:          nil,
		placeholderEnd2:            nil,
		statementStart:             statementStart,
		statementEnd:               statementEnd,
		statementStart2:            nil,
		statementEnd2:              nil,
		isIgnoreClientIOError:      true,
		ps:                         make(map[string]string),
		resourceMap:                make(map[string]string),
		pkgList:                    make(map[string]interface{}),
		fnMap:                      make(map[string]string),
		fnPkgMap:                   make(map[string]string),
		formatMap:                  make(map[string]string),
		virtualClass:               make(map[string]string),
		tagFactoryMap:              make(map[string]string),
		tagMap:                     make(map[string]string),
		defaultFormatMap:           make(map[string]string),
		generalVirtualAttributeSet: make(map[string]interface{}),
	}
	inst.initDefault()
	if customPropertiesFilePath != nil {
		inst.initCustomProperties(*customPropertiesFilePath)
	}
	return inst
}

func (c *Configuration) Build() {
	if len(strings.TrimSpace(c.placeholderStart)) == 0 ||
		len(strings.TrimSpace(c.placeholderEnd)) == 0 {
		panic("占位符不能为空")
	}

	if c.placeholderStart2 != nil {
		if len(strings.TrimSpace(*c.placeholderStart2)) == 0 ||
			c.placeholderEnd2 == nil ||
			len(strings.TrimSpace(*c.placeholderEnd2)) == 0 {
			panic("定义了2对占位符配置，但占位符2不能为空")
		}
		start2 := []rune(*c.placeholderStart2)
		end2 := []rune(*c.placeholderEnd2)
		c.pd = NewDelimeterHolder([]rune(c.placeholderStart), []rune(c.placeholderEnd), &start2, &end2)
	} else {
		c.pd = NewDelimeterHolder([]rune(c.placeholderStart), []rune(c.placeholderEnd), nil, nil)
	}

	if len(strings.TrimSpace(c.statementStart)) == 0 {
		panic("定界符起始符号不能为空")
	}
	if c.statementStart2 != nil {
		if len(strings.TrimSpace(*c.statementStart2)) == 0 ||
			c.statementEnd2 == nil {
			panic("定义了2对定界符配置，但定界符2不能为空")
		}
		start2 := []rune(*c.statementStart2)
		end2 := []rune(*c.statementEnd2)
		c.sd = NewDelimeterHolder([]rune(c.statementStart), []rune(c.statementEnd), &start2, &end2)
	} else {
		c.sd = NewDelimeterHolder([]rune(c.statementStart), []rune(c.statementEnd), nil, nil)
	}
}

func (c *Configuration) getScriptDelimeter() *DelimeterHolder {
	return c.sd
}

func (c *Configuration) getPlaceHolderDelimeter() *DelimeterHolder {
	return c.pd
}

func (c *Configuration) getResourceMap() map[string]string {
	return c.resourceMap
}

func (c *Configuration) initDefault() {
	ps, err := util.LoadProperty("resources/beetl-default.properties")
	if err != nil {
		fmt.Println(err)
		panic("请在项目根目录添加 resources/beetl-default.properties")
	}
	c.ps = ps
	c.parseProperties(ps)
}

func (c *Configuration) initCustomProperties(filePath string) {
	ps, err := util.LoadProperty(filePath)
	if err != nil {
		fmt.Println(err)
		panic("请添加文件 " + filePath)
	}
	c.parseProperties(ps)
}

func (c *Configuration) parseProperties(ps map[string]string) {
	for key, value := range ps {
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		v := value

		if equalsIgnoreCase(key, DelimeterConfig) {
			c.delimeterClass = &v
		} else if equalsIgnoreCase(key, DelimiterPlaceholderStart) {
			c.placeholderStart = value
		} else if equalsIgnoreCase(key, DelimiterPlaceholderEnd) {
			c.placeholderEnd = value
		} else if equalsIgnoreCase(key, DelimiterStatementStart) {
			c.statementStart = value
		} else if equalsIgnoreCase(key, DelimiterStatementEnd) {
			if value == "null" || value == "nil" {
				value = ""
			}
			c.statementEnd = value
		} else if equalsIgnoreCase(key, DelimiterPlaceholderStart2) {
			c.placeholderStart2 = &v
		} else if equalsIgnoreCase(key, DelimiterPlaceholderEnd2) {
			c.placeholderEnd2 = &v
		} else if equalsIgnoreCase(key, DelimiterStatementStart2) {
			c.statementStart2 = &v
		} else if equalsIgnoreCase(key, DelimiterStatementEnd2) {
			if value == "null" || value == "nil" {
				value = ""
			}
			c.statementEnd2 = &v
		} else if equalsIgnoreCase(key, NativeCall) {
			b, err := strconv.ParseBool(value)
			if err == nil {
				c.nativeCall = b
			}
		} else if equalsIgnoreCase(key, SafeOutput) {
			b, err := strconv.ParseBool(value)
			if err == nil {
				c.SafeOutput = b
			}
		} else if equalsIgnoreCase(key, IgnoreClientIoError) {
			b, err := strconv.ParseBool(value)
			if err == nil {
				c.isIgnoreClientIOError = b
			}
		} else if equalsIgnoreCase(key, DirectByteOutput) {
			b, err := strconv.ParseBool(value)
			if err == nil {
				c.directByteOutput = b
			}
		} else if equalsIgnoreCase(key, MvcStrict) {
			b, err := strconv.ParseBool(value)
			if err == nil {
				c.isStrict = b
			}
		} else {
			if hasPrefixIgnoreCase(key, "fn.") {
				fn := checkValue(value)
				c.fnMap[getExtName(key)] = fn
			} else if hasPrefixIgnoreCase(key, "fnp.") {
				fn := checkValue(value)
				c.fnPkgMap[getExtName(key)] = fn
			} else if hasPrefixIgnoreCase(key, "ft.") {
				fn := checkValue(value)
				c.formatMap[getExtName(key)] = fn
			} else if hasPrefixIgnoreCase(key, "ftc.") {
				fn := checkValue(value)
				c.defaultFormatMap[getExtName(key)] = fn
			} else if hasPrefixIgnoreCase(key, "virtual.") {
				fn := checkValue(value)
				c.virtualClass[getExtName(key)] = fn
			} else if hasPrefixIgnoreCase(key, "tag.") {
				fn := checkValue(value)
				c.tagMap[getExtName(key)] = fn
			} else if hasPrefixIgnoreCase(key, "tagf.") {
				c.tagFactoryMap[getExtName(key)] = value
			} else if hasPrefixIgnoreCase(key, "resource.") {
				c.resourceMap[getExtName(key)] = value
			}
		}
	}
}

func equalsIgnoreCase(key0 string, key1 string) bool {
	return strings.ToUpper(key0) == strings.ToUpper(key1)
}

func hasPrefixIgnoreCase(src string, prefix string) bool {
	return strings.HasPrefix(strings.ToUpper(src), strings.ToUpper(prefix))
}

func checkValue(value string) string {
	vals := strings.Split(value, ",")
	if len(vals) == 1 {
		return value
	}
	//cls := vals[1]
	// TODO 判断是否已注册
	return vals[0]
}

func getExtName(key string) string {
	return key[strings.Index(key, ".")+1:]
}

type DelimeterHolder struct {
	start  []rune
	end    []rune
	start1 *[]rune
	end1   *[]rune
}

func NewDelimeterHolder(start []rune, end []rune, start1 *[]rune, end1 *[]rune) *DelimeterHolder {
	inst := &DelimeterHolder{
		start:  start,
		end:    end,
		start1: start1,
		end1:   end1,
	}
	return inst
}

func (h *DelimeterHolder) CreatePhd() *PlaceHolderDelimeter {
	return NewPlaceHolderDelimeter(h.start, h.end, h.start1, h.end1)
}

func (h *DelimeterHolder) CreateSd() *ScriptDelimeter {
	return NewScriptDelimeter(h.start, h.end, h.start1, h.end1)
}
