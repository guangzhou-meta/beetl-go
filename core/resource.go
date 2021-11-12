package core

import (
	"bufio"
	"bytes"
	"io"
)

import (
	"github.com/guangzhou-meta/beetl-go/util"
)

import (
	estrings "github.com/guangzhou-meta/go-lib/strings"
)

type ResourceI interface {
	OpenReader() io.Reader
	IsModified() bool

	GetId() interface{}

	GetContent(start int, end int) string
}

type Resource struct {
	resourceLoader ResourceLoaderI
	id             interface{}
	children       ResourceI
}

func NewResource(id interface{}, loader ResourceLoaderI) *Resource {
	inst := &Resource{
		id:             id,
		resourceLoader: loader,
	}
	return inst
}

func (r *Resource) GetResourceLoader() ResourceLoaderI {
	return r.resourceLoader
}

func (r *Resource) SetResourceLoader(resourceLoader ResourceLoaderI) {
	r.resourceLoader = resourceLoader
}

func (r *Resource) GetId() interface{} {
	return r.id
}

func (r *Resource) GetContent(start int, end int) string {
	lineSeparator := util.GetLineSeparator()
	br := r.OpenReader()
	reader := bufio.NewReader(br)
	var line []byte
	var err error
	sb := estrings.NewStringBuilder()
	index := 0
	line, _, err = reader.ReadLine()
	for err == nil {
		index++
		if index >= start && index <= end {
			sb.Append(string(line)).Append(lineSeparator)
			if index == end {
				break
			}
		}
		line, _, err = reader.ReadLine()
	}
	return sb.String()
}

func (r *Resource) OpenReader() io.Reader {
	panic("not implement")
}

func (r *Resource) IsModified() bool {
	panic("not implement")
}

type StringTemplateResource struct {
	*Resource
}

func NewStringTemplateResource(template string, resourceLoader ResourceLoaderI) *StringTemplateResource {
	inst := &StringTemplateResource{
		Resource: NewResource(template, resourceLoader),
	}
	return inst
}

func (r *StringTemplateResource) OpenReader() io.Reader {
	return bufio.NewReader(bytes.NewBufferString(r.id.(string)))
}

func (r *StringTemplateResource) IsModified() bool {
	return false
}
