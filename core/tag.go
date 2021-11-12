package core

import (
	"reflect"
	"strings"
)

import (
	"github.com/guangzhou-meta/beetl-go/util"
)

type TagI interface {
	Render()
	Init(ctx *Context, args []interface{}, block StatementI)
	AfterRender()

	ClassName() string

	GetClassType() reflect.Type
}

type Tag struct {
	args   []interface{}
	gt     *GroupTemplate
	ctx    *Context
	bw     ByteWriterI
	bs     StatementI
	parent *Tag
}

func NewTag() *Tag {
	inst := &Tag{}
	return inst
}

func (t *Tag) Init(ctx *Context, args []interface{}, st StatementI) {
	t.InitBase(ctx, args, st)
	t.setTagParent()
}

func (t *Tag) InitBase(ctx *Context, args []interface{}, st StatementI) {
	t.ctx = ctx
	t.bw = ctx.byteWriter
	t.gt = ctx.gt
	t.args = args
	t.bs = st
}

func (t *Tag) DoBodyRender() {
	t.bs.Execute(t.ctx)
}

func (t *Tag) Render() {
	panic("not implement")
}

func (t *Tag) AfterRender() {
	t.ctx.SetCurrentTag(t.parent)
}

func (t *Tag) setTagParent() {
	t.parent = t.ctx.GetCurrentTag()
	t.ctx.SetCurrentTag(t)
}

func (t *Tag) GetParent() *Tag {
	return t.parent
}

func (t *Tag) GetArgs() []interface{} {
	return t.args
}

func (t *Tag) GetHtmlAttribute(attrName string) interface{} {
	obj := t.args[1]
	v := util.UnpackValue(reflect.ValueOf(obj))
	switch v.Kind() {
	case reflect.Map:
		return v.MapIndex(reflect.ValueOf(attrName)).Interface()
	}
	panic("非html标签")
}

func (t *Tag) ContainHtmlAttribute(attrName string) bool {
	obj := t.args[1]
	v := util.UnpackValue(reflect.ValueOf(obj))
	switch v.Kind() {
	case reflect.Map:
		return v.MapIndex(reflect.ValueOf(attrName)).IsValid()
	}
	panic("非html标签")
}

func (t *Tag) ClassName() string {
	panic("not implement")
}

func (t *Tag) GetClassType() reflect.Type {
	return t.GetType(t)
}

func (t *Tag) GetType(i interface{}) reflect.Type {
	return util.UnpackType(reflect.TypeOf(i))
}

type HTMLTagVarBindingWrapper struct {
	*Tag

	tag *GeneralVarTagBinding
}

func NewHTMLTagVarBindingWrapper() *HTMLTagVarBindingWrapper {
	return &HTMLTagVarBindingWrapper{
		Tag: NewTag(),
	}
}

func (t *HTMLTagVarBindingWrapper) Init(ctx *Context, args []interface{}, st StatementI) {
	t.Tag.Init(ctx, args, st)
	if len(args) == 0 || len(args) > 3 {
		panic("参数错误，期望child,Map .....")
	}
	child := args[0].(string)
	functionTagName := strings.Replace(child, ":", ".", 1)
	tagFactory := t.gt.GetTagFactory(functionTagName)
	if tagFactory == nil {
		panic("标签初始化错误，未找到指定的标签实现类" + functionTagName)
	}
	temp := tagFactory.CreateTag()
	if temp == nil {
		panic("找不到注册的Tag")
	}
	if tag, ok := temp.(*GeneralVarTagBinding); ok {
		t.tag = tag
		tag.Init(ctx, args, st)
	}
}

func (t *HTMLTagVarBindingWrapper) Render() {
	t.tag.Render()
}

func (t *HTMLTagVarBindingWrapper) MapName2Index(m map[string]int) {
	t.tag.MapName2Index(m)
}

func (t *HTMLTagVarBindingWrapper) GetClassType() reflect.Type {
	return t.Tag.GetType(t)
}

type GeneralVarTagBinding struct {
	*Tag

	name2Index map[string]int
}

func NewGeneralVarTagBinding() *GeneralVarTagBinding {
	return &GeneralVarTagBinding{
		Tag: NewTag(),
	}
}

func (b *GeneralVarTagBinding) MapName2Index(m map[string]int) {
	b.name2Index = m
}

func (b *GeneralVarTagBinding) GetAttributeValue(attrName string) interface{} {
	m := b.args[1].(map[string]interface{})
	return m[attrName]
}

func (b *GeneralVarTagBinding) GetHtmlTagName() string {
	return b.args[1].(string)
}

func (b *GeneralVarTagBinding) GetAttributes() map[string]interface{} {
	return b.args[1].(map[string]interface{})
}

func (b *GeneralVarTagBinding) GetClassType() reflect.Type {
	return b.Tag.GetType(b)
}

type SimpleTag struct {
	*Tag
}

func NewSimpleTag() *SimpleTag {
	return &SimpleTag{
		Tag: NewTag(),
	}
}

func (f *SimpleTag) Render() {
}

func (f *SimpleTag) ClassName() string {
	return "org.beetl.ext.tag.SimpleTag"
}

func (f *SimpleTag) GetClassType() reflect.Type {
	return f.Tag.GetType(f)
}
