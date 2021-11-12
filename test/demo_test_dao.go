package main

import (
	"github.com/guangzhou-meta/beetl-go/common"
	"github.com/guangzhou-meta/beetl-go/engine"
)

type DemoTest struct {
	DemoTest func(params engine.RenderParamsI) ([]common.SqlResultI, error)
}

func (*DemoTest) TemplateName() string {
	return "DemoTest"
}

type DemoTestDTO struct {
	Name *string `json:"name"`
}
