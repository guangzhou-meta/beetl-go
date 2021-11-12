package core

import (
	"fmt"
	"reflect"
	"strings"
)

import (
	"github.com/guangzhou-meta/beetl-go/util"
)

type GroupTemplate struct {
	conf                   *Configuration
	engine                 TemplateEngineI
	fnMap                  map[string]FunctionI
	formatMap              map[string]FormatI
	defaultFormatMap       map[reflect.Type]FormatI
	virtualAttributeList   []VirtualAttributeEval
	virtualClass           map[reflect.Type]VirtualClassAttribute
	tagFactoryMap          map[string]TagFactoryI
	sharedVars             map[string]interface{}
	htmlTagAttrNameConvert AttributeNameConvertI
	delimeterConfig        DelimeterConfigI
	Function               *EngineFunction
	Tag                    *EngineTag
}

func NewGroupTemplate(conf *Configuration) *GroupTemplate {
	inst := &GroupTemplate{
		conf:                 conf,
		engine:               NewBeetlSQLTemplateEngine(),
		virtualAttributeList: make([]VirtualAttributeEval, 0),
		virtualClass:         make(map[reflect.Type]VirtualClassAttribute),
		formatMap:            make(map[string]FormatI),
		defaultFormatMap:     make(map[reflect.Type]FormatI),
		fnMap:                make(map[string]FunctionI),
		tagFactoryMap:        make(map[string]TagFactoryI),
		sharedVars:           make(map[string]interface{}),
	}
	inst.init()
	return inst
}

func (t *GroupTemplate) LoadTemplate(res ResourceI) *Program {
	render := res.OpenReader()
	text := NewTextParser(t, t.delimeterConfig.GetPlaceHolder(res.GetId()), t.delimeterConfig.GetStatement(res.GetId()))
	text.DoParse(render)
	return t.engine.CreateProgram(res, text.GetScript().String(), text.GetTextVars(), text.GetTextCr(), t)
}

func (t *GroupTemplate) GetHtmlTagAttrNameConvert() AttributeNameConvertI {
	return t.htmlTagAttrNameConvert
}

func (t *GroupTemplate) init() {
	t.conf.Build()

	t.initFunction()
	t.initFormatter()
	t.initTag()
	t.initVirtual()
	if t.conf.delimeterClass == nil {
		t.delimeterConfig = NewDefaultDelimeterConfig(t.conf)
	}
}

func (t *GroupTemplate) RegisterVirtualAttributeEval(e VirtualAttributeEval) {
	t.virtualAttributeList = append(t.virtualAttributeList, e)
}

func (t *GroupTemplate) RegisterVirtualAttributeClass(c reflect.Type, virtual VirtualClassAttribute) {
	t.virtualClass[c] = virtual
}

func (t *GroupTemplate) GetVirtualAttributeEval(c reflect.Type, attributeName string) VirtualClassAttribute {
	attr := t.virtualClass[c]
	if attr != nil {
		return attr
	}
	for _, eval := range t.virtualAttributeList {
		if eval.IsSupport(c, attributeName) {
			return eval
		}
	}
	return nil
}

func (t *GroupTemplate) registerFormat(name string, format FormatI) {
	t.formatMap[name] = format
}

func (t *GroupTemplate) registerDefaultFormat(tp reflect.Type, format FormatI) {
	t.defaultFormatMap[tp] = format
}

func (t *GroupTemplate) GetFormat(name string) FormatI {
	return t.formatMap[name]
}

func (t *GroupTemplate) GetDefaultFormat(tp reflect.Type) FormatI {
	return t.defaultFormatMap[tp]
}

func (t *GroupTemplate) GetTagFactory(name string) TagFactoryI {
	return t.tagFactoryMap[name]
}

func (t *GroupTemplate) checkTagName(name string) {
	if !util.CheckNaming(name) {
		log := util.CheckResult
		c := string([]rune{rune(log[1])})
		if log[1] == 58 {
			panic(fmt.Sprintf("注册Tag名称不合法:%s,错误位置:%d,出现错误的字符:%s,请使用'.'", name, log[0], c))
		} else {
			panic(fmt.Sprintf("注册Tag名称不合法:%s,错误位置:%d,出现错误的字符:%s", name, log[0], c))
		}
	}
}

func (t *GroupTemplate) GetFunction(name string) FunctionI {
	return t.fnMap[name]
}

func (t *GroupTemplate) initFunction() {
	t.Function = NewEngineFunction()
	fnMap := t.conf.fnMap
	for name, clsName := range fnMap {
		t.registerFunction(name, clsName)
	}

	fnPkgMap := t.conf.fnPkgMap
	for name, clsName := range fnPkgMap {
		t.registerFunctionPackage(name, clsName)
	}
}

func (t *GroupTemplate) initFormatter() {
	formatMap := t.conf.formatMap
	for name, clsName := range formatMap {
		t.registerFormat(name, t.Function.registerFmtMap[clsName])
	}
}

func (t *GroupTemplate) initTag() {
	t.Tag = NewEngineTag()
	tagMap := t.conf.tagMap
	for name, clsName := range tagMap {
		t.registerTag(name, clsName)
	}

	tagFactoryMap := t.conf.tagFactoryMap
	for name, clsName := range tagFactoryMap {
		t.registerTagFactory(name, clsName)
	}
}

func (t *GroupTemplate) initVirtual() {
}

func (t *GroupTemplate) registerFunction(name string, className string) {
	if !util.CheckNaming(name) {
		result := util.CheckResult
		panic(fmt.Sprintf("注册方法名不合法:%s,错误位置:%d,出现错误的字符:%s", name, result[0], string(result[1])))
	}
	if fun, ok := t.Function.registerFunMap[className]; ok {
		t.fnMap[name] = fun
	}
}

func (t *GroupTemplate) registerFunctionPackage(packageName string, className string) {
	if "_root" == packageName {
		t.registerFunctionPackageAsRoot(className)
	} else {
		list := GetFunctionWrapper(packageName,
			t.Function.registerFunMap[className].GetClassType())
		for _, fw := range list {
			t.registerFunction(fw.GetMethodName(), className)
		}
	}
}

func (t *GroupTemplate) registerFunctionPackageAsRoot(className string) {
	packageName := "_root"
	list := GetFunctionWrapper(packageName, t.Function.registerFunMap[className].GetClassType())
	for _, fw := range list {
		functionName := strings.Replace(fw.GetMethodName(), "_root", "", 1)
		t.registerFunction(functionName, className)
	}
}

func (t *GroupTemplate) GetSharedVars() map[string]interface{} {
	return t.sharedVars
}

func (t *GroupTemplate) SetSharedVars(m map[string]interface{}) {
	for k, v := range m {
		t.sharedVars[k] = v
	}
}

func (t *GroupTemplate) registerTag(name string, className string) {
	if !util.CheckNaming(name) {
		result := util.CheckResult
		panic(fmt.Sprintf("注册方法名不合法:%s,错误位置:%d,出现错误的字符:%s", name, result[0], string(result[1])))
	}
	if fun, ok := t.Tag.registerTagMap[className]; ok {
		t.tagFactoryMap[name] = NewDefaultTagFactory(fun)
	}
}

func (t *GroupTemplate) registerTagFactory(name string, className string) {
	if !util.CheckNaming(name) {
		result := util.CheckResult
		panic(fmt.Sprintf("注册方法名不合法:%s,错误位置:%d,出现错误的字符:%s", name, result[0], string(result[1])))
	}
	if fun, ok := t.Tag.registerTagFactoryMap[className]; ok {
		t.tagFactoryMap[name] = fun
	}
}
