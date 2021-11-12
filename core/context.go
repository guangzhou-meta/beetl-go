package core

import (
	"github.com/guangzhou-meta/beetl-go/core/cache"
)

var (
	NotExistObject = new(interface{})
)

type Context struct {
	byteWriter        ByteWriterI
	Template          *Template
	gt                *GroupTemplate
	GlobalVar         map[string]interface{}
	byteOutputMode    bool
	Vars              []interface{}
	StaticTextArray   []interface{}
	TempVarStartIndex int
	GotoFlag          IGoto
	isInit            bool
	safeOutput        bool
	LocalBuffer       *cache.ContextBuffer
}

func NewContext(gt *GroupTemplate, buffer *cache.ContextBuffer) *Context {
	inst := &Context{
		gt:          gt,
		LocalBuffer: buffer,
		GlobalVar:   make(map[string]interface{}),
	}
	return inst
}

func (c *Context) exist(i int) bool {
	if i >= c.TempVarStartIndex {
		return true
	}
	obj := c.Vars[i]
	return obj != NotExistObject
}

func (c *Context) Set(key string, value interface{}) {
	c.GlobalVar[key] = value
}

func (c *Context) GetGlobal(key string) interface{} {
	return c.GlobalVar[key]
}

func (c *Context) GetResourceId() interface{} {
	return c.Template.Resource.GetId()
}

func (c *Context) GetResource() ResourceI {
	return c.Template.Resource
}

func (c *Context) SetCurrentTag(tag *Tag) {
	m := c.GetGlobal("$page")
	if m == nil {
		m = make(map[string]interface{})
		c.Set("$page", m)
	}
	m.(map[string]interface{})["$parentTag"] = tag
}

func (c *Context) GetCurrentTag() *Tag {
	m := c.GetGlobal("$page")
	if m == nil {
		return nil
	}
	if t, ok := m.(map[string]interface{})["$parentTag"]; ok {
		return t.(*Tag)
	}
	return nil
}

func (c *Context) Destroy() {
	//c.gt.buffers.PutContextLocalBuffer(c.LocalBuffer)
}

func (c *Context) GetByteWriter() ByteWriterI {
	return c.byteWriter
}

//func (c *Context) GetByteWriter() *estrings.StringBuilder {
//	return c.byteWriter
//}
