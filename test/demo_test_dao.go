package main

import (
	"github.com/guangzhou-meta/beetl-go/common"
	"github.com/guangzhou-meta/beetl-go/engine"
)

type DemoTest struct {
	DemoTest  func(params engine.RenderParamsI) ([]common.SqlResultI, error)
	DemoTest2 func(params engine.RenderParamsI) ([]common.SqlResultI, error)
}

func (*DemoTest) TemplateName() string {
	return "DemoTest"
}

type DemoTestDTO struct {
	Name *string `json:"name"`
}

type DemoTest2DTO struct {
	Name  *string `json:"name"`
	Phone *string `json:"phone"`
	Email *string `json:"email"`
}
