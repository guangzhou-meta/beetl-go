package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

import (
	"github.com/guangzhou-meta/beetl-go/common"
	"github.com/guangzhou-meta/beetl-go/engine"
)

func TestDemo(t *testing.T) {
	root, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	templateEngine := engine.NewEngine().
		SetTemplateDir(root + "/resources/sql").
		SetPropertiesFilePath(nil).
		Build()

	instance := &DemoTest{}
	templateEngine.RegisterTemplate(instance)

	var sqlList []common.SqlResultI
	for i := 0; i < 10; i++ {
		fmt.Println("=== " + strconv.FormatInt(int64(i), 10) + " ===")
		p := strconv.FormatInt(int64(i), 10)
		params := &DemoTestDTO{
			Name: &p,
		}
		sqlList, err = instance.DemoTest(engine.NewRenderPageParams().SetObject(params))
		if err != nil {
			t.Errorf("Render Fail: %v", err)
			return
		}
		for _, sql := range sqlList {
			fmt.Println(sql.GetSqlType(), ": ", sql)
		}
	}
}
