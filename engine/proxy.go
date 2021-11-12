package engine

import (
	"bufio"
	"bytes"
	"reflect"
	"strings"
)

import (
	perrors "github.com/pkg/errors"
)

func (g *Engine) ImplementFunc(templateName string, obj interface{}, info map[string]*methodInfo, sqlMap map[string]string) {
	objValue := reflect.ValueOf(obj)
	for reflect.Ptr == objValue.Kind() {
		objValue = objValue.Elem()
	}

	makeFunc := func(templateName string, methodName string) func(params []reflect.Value) (results []reflect.Value) {
		m := info[methodName]
		return func(params []reflect.Value) (results []reflect.Value) {
			reply := reflect.New(m.RepliesType[0]).Elem()
			sqlTemplate, ok := sqlMap[methodName]
			if !ok {
				e := perrors.Errorf("`%s`中找不到`%s`Sql模板", templateName, methodName)
				return []reflect.Value{reply, reflect.ValueOf(e)}
			}
			// 处理SQL模板
			res := g.Render(sqlTemplate, params[0].Interface().(RenderParamsI))
			for _, sql := range res {
				nSql := genSql(sql.GetSql())
				sql.SetSql(nSql)
				reply.Set(reflect.Append(reply, reflect.ValueOf(sql)))
			}
			return []reflect.Value{
				reply,
				reflect.New(m.RepliesType[1]).Elem(),
			}
		}
	}

	for k := range info {
		m := objValue.FieldByName(k)
		if m.IsValid() && m.CanSet() {
			m.Set(reflect.MakeFunc(m.Type(), makeFunc(templateName, k)))
		}
	}
}

func genSql(sql string) string {
	reader := bufio.NewReader(bytes.NewBufferString(sql))
	sql = ""
	for l, _, e := reader.ReadLine(); e == nil; {
		sql += strings.TrimSpace(string(l)) + " "
		l, _, e = reader.ReadLine()
	}
	return strings.TrimSpace(sql)
}
