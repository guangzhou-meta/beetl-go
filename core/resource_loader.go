package core

import (
	"io"
)

type ResourceLoaderI interface {
	io.Closer

	GetResource(key interface{}) ResourceI

	IsModified(key ResourceI) bool

	Exist(key interface{}) bool

	Init(gt *GroupTemplate)

	GetResourceId(resource ResourceI, key interface{}) interface{}

	GetInfo() string
}

type StringTemplateResourceLoader struct {
}

func NewStringTemplateResourceLoader() *StringTemplateResourceLoader {
	return &StringTemplateResourceLoader{}
}

func (rl *StringTemplateResourceLoader) GetResource(key interface{}) ResourceI {
	return NewStringTemplateResource(key.(string), rl)
}

func (rl *StringTemplateResourceLoader) IsModified(key ResourceI) bool {
	return false
}

func (rl *StringTemplateResourceLoader) Exist(key interface{}) bool {
	return true
}

func (rl *StringTemplateResourceLoader) Init(gt *GroupTemplate) {}

func (rl *StringTemplateResourceLoader) GetResourceId(resource ResourceI, key interface{}) interface{} {
	return key.(string)
}

func (rl *StringTemplateResourceLoader) GetInfo() string {
	return "StringTemplateResourceLoader "
}

func (rl *StringTemplateResourceLoader) Close() error {
	return nil
}
