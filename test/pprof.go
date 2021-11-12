package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
)

import (
	"github.com/guangzhou-meta/beetl-go/common"
	"github.com/guangzhou-meta/beetl-go/engine"
)

type PPROFDemoTest struct {
	DemoTest func(params engine.RenderParamsI) ([]common.SqlResultI, error)
}

func (*PPROFDemoTest) TemplateName() string {
	return "DemoTest"
}

type PPROFDemoTestDTO struct {
	Name *string `json:"name"`
}

func main() {
	go func() {
		root, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}
		templateEngine := engine.NewEngine().
			SetTemplateDir(root + "/test/resources/sql").
			SetPropertiesFilePath(nil).
			Build()

		instance := &PPROFDemoTest{}
		templateEngine.RegisterTemplate(instance)

		var sqlList []common.SqlResultI
		for i := 0; i < 100000; i++ {
			fmt.Println("=== " + strconv.FormatInt(int64(i), 10) + " ===")
			p := strconv.FormatInt(int64(i), 10)
			params := &PPROFDemoTestDTO{
				Name: &p,
			}
			sqlList, err = instance.DemoTest(engine.NewRenderPageParams().SetObject(params))
			if err != nil {
				fmt.Printf("Render Fail: %v\n", err)
				return
			}
			for _, sql := range sqlList {
				fmt.Println(sql.GetSqlType(), ": ", sql)
			}
		}
	}()

	http.ListenAndServe("0.0.0.0:25052", nil)
}
