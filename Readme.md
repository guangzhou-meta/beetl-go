# Beetl-go

数据库模板引擎，适用于Golang。

## 介绍

核心思想来源于[https://gitee.com/xiandafu/beetl](https://gitee.com/xiandafu/beetl)

## 安装
```shell
go get -u github.com/guangzhou-meta/beetl-go
```
## 示例

### 1. 添加脚本文件

/resources/sql/DemoTest.md

```markdown
DemoTest
===
```sql
select #{page('*')} from sys_user 
where 1=1
-- @if(!isEmpty(name)){
 and name=#{name}
-- @}
-- @pageIgnoreTag(){
 order by createTime
-- @}
```[实际可参考 test/resources/sql/DemoTest.md]

DemoTest2
===
```sql
select #{page('*')} from sys_user 
where 1=1
-- @if(!isEmpty(name)){
 and name=#{name}
-- @}
-- @pageIgnoreTag(){
 order by createTime
-- @}
```[实际可参考 test/resources/sql/DemoTest.md]
```
### 2. 添加默认配置文件

/resources/beetl-default.properties
> 参考 resources/beetl-default.properties

### 3. 添加模板抽象Struct

```go
package dao

import (
	"github.com/guangzhou-meta/beetl-go/common"
	"github.com/guangzhou-meta/beetl-go/engine"
	"github.com/guangzhou-meta/beetl-go/template"
)

type DemoTest struct {
	DemoTest func(params engine.RenderParamsI) ([]common.SqlResultI, error)
}

func (*DemoTest) TemplateName() string {
	return "DemoTest"
}
```

### 4. 初始化引擎

```go
package main

import (
	"fmt"
	"os"
)

import (
	"github.com/guangzhou-meta/beetl-go/engine"
	"github.com/guangzhou-meta/beetl-go/template"
)

func init() {
	root, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	templateEngine := engine.NewEngine().
		SetTemplateDir(root + "/resources/sql"). // 模板扫描地址
		SetPropertiesFilePath(nil). // 自定义配置文件路径
		Build()          
}
```

### 5. 注册模板

```go
package main

import (
	"dao"
)

var demoTestInstance *dao.DemoTest

func register() {
	demoTestInstance = &dao.DemoTest{}
	templateEngine.RegisterTemplate(demoTestInstance)
}
```

### 6. 调用模板
```go
package main

import (
	"dao"
	"fmt"
)

import (
	"github.com/guangzhou-meta/beetl-go/engine"
	"github.com/guangzhou-meta/beetl-go/template"
)

var demoTestInstance *dao.DemoTest

type DemoTestDTO struct {
	Name *string `json:"name"`
}

func main()  {
    register() // 注册模板
	
	name := "Unizone"
	params := &DemoTestDTO{
		Name: &name,
	}
	sqlList, err := demoTestInstance.DemoTest(engine.NewRenderPageParams().SetObject(params))
	if err != nil {
		fmt.Printf("Render Fail: %v\n", err)
		return
	}
	for _, sql := range sqlList {
		fmt.Println(sql.GetSqlType(), ": ", sql)
	}
}
```

## 特别鸣谢

* [闲.大赋](https://gitee.com/xiandafu)
* [ibeetl项目组](https://ibeetl.com)
