package engine

import (
	"reflect"
	"strings"
)

import (
	"github.com/guangzhou-meta/beetl-go/common"
	"github.com/guangzhou-meta/beetl-go/core"
	"github.com/guangzhou-meta/beetl-go/util"
)

type RenderParamsI interface {
	SetParams(params map[string]interface{}, force bool) RenderParamsI
	SetParam(key string, value interface{}) RenderParamsI
	GetParam(key string, defaultValue interface{}) interface{}
	GetParams() map[string]interface{}
}

type RenderParams struct {
	params map[string]interface{}
}

func NewRenderParams() *RenderParams {
	return &RenderParams{
		params: make(map[string]interface{}),
	}
}

func (p *RenderParams) SetObject(m interface{}) RenderParamsI {
	mv := util.UnpackValue(reflect.ValueOf(m))
	mt := mv.Type()
	for i := 0; i < mv.NumField(); i++ {
		f := mv.Field(i)
		if f.IsNil() || f.IsZero() || !f.IsValid() {
			continue
		}
		fk := mt.Field(i).Tag.Get("json")
		if len(strings.TrimSpace(fk)) == 0 {
			fk = f.String()
		}
		fv := util.UnpackValue(f).Interface()
		p.SetParam(fk, fv)
	}
	return p
}

func (p *RenderParams) SetParams(params map[string]interface{}, force bool) RenderParamsI {
	if force { // 覆盖
		p.params = params
	} else {
		for k, v := range params {
			p.SetParam(k, v)
		}
	}
	return p
}

func (p *RenderParams) SetParam(key string, value interface{}) RenderParamsI {
	p.params[key] = value
	return p
}

func (p *RenderParams) GetParam(key string, defaultValue interface{}) interface{} {
	if v, ok := p.params[key]; ok {
		if v != nil {
			return v
		}
	}
	return defaultValue
}

func (p *RenderParams) GetParams() map[string]interface{} {
	return p.params
}

type RenderPageParams struct {
	*RenderParams

	Size int
	Page int
}

func NewRenderPageParams() *RenderPageParams {
	return &RenderPageParams{
		RenderParams: NewRenderParams(),
		Size:         10,
		Page:         1,
	}
}

func (p *RenderPageParams) SetSize(size int) *RenderPageParams {
	if size < 1 {
		panic("Size必须大于0")
	}
	p.Size = size
	return p
}

func (p *RenderPageParams) GetSize() int {
	return p.Size
}

func (p *RenderPageParams) SetPage(page int) *RenderPageParams {
	if page < 1 {
		panic("Page必须大于0")
	}
	p.Page = page
	return p
}

func (p *RenderPageParams) GetPage() int {
	return p.Page
}

func (p *RenderPageParams) SetParams(params map[string]interface{}, force bool) RenderParamsI {
	p.RenderParams.SetParams(params, force)
	return p
}

func (p *RenderPageParams) SetParam(key string, value interface{}) RenderParamsI {
	p.RenderParams.SetParam(key, value)
	return p
}

func (p *RenderPageParams) GetParam(key string, defaultValue interface{}) interface{} {
	p.RenderParams.GetParam(key, defaultValue)
	return defaultValue
}

func (p *RenderPageParams) GetParams() map[string]interface{} {
	return p.RenderParams.GetParams()
}

func (p *RenderPageParams) GetStart(offsetStartZero bool) int {
	o := 1
	if offsetStartZero {
		o = 0
	}
	o += (p.GetPage() - 1) * p.GetSize()
	return o
}

func (p *RenderPageParams) SetObject(m interface{}) *RenderPageParams {
	p.RenderParams.SetObject(m)
	return p
}

func (g *Engine) Build() *Engine {
	g.conf = core.NewConfiguration(g.propertiesFilePath)
	g.gt = core.NewGroupTemplate(g.conf)
	g.resL = core.NewStringTemplateResourceLoader()
	g.initialized = true
	return g
}

func (g *Engine) Render(sql string, params RenderParamsI) []common.SqlResultI {
	if !g.initialized {
		g.Build()
	}
	resource := g.resL.GetResource(sql)
	program := g.gt.LoadTemplate(resource)
	template := core.NewTemplate(g.gt, program, resource, g.conf, nil)
	template.Binding(params.GetParams())
	if pp, ok := params.(*RenderPageParams); ok {
		template.SetPageMode(pp.Page, pp.Size, pp.GetStart(false))
	}
	return template.Render()
}
