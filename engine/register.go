package engine

import (
	"fmt"
	"github.com/guangzhou-meta/beetl-go/core"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
)

import (
	"github.com/guangzhou-meta/beetl-go/common"
)

import (
	perrors "github.com/pkg/errors"
)

type Engine struct {
	templateDir        string
	templateRegistry   *TemplateRegistry
	propertiesFilePath *string
	gt                 *core.GroupTemplate
	resL               *core.StringTemplateResourceLoader
	conf               *core.Configuration
	initialized        bool
}

func NewEngine() *Engine {
	return &Engine{
		templateDir: "/",
		templateRegistry: &TemplateRegistry{
			registry: make(map[string]*structInfo),
		},
		propertiesFilePath: nil,
	}
}

func (g *Engine) SetTemplateDir(templateDir string) *Engine {
	templateDir = filepath.ToSlash(strings.Trim(templateDir, " "))
	if !filepath.IsAbs(templateDir) {
		panic(perrors.Errorf("模板目录必须为绝对路径"))
	}
	if !strings.HasSuffix(templateDir, common.C1) {
		templateDir = templateDir + common.C1
	}
	fmt.Println("Template Dir: ", templateDir)
	g.templateDir = templateDir
	return g
}

func (g *Engine) SetPropertiesFilePath(propertiesFilePath *string) *Engine {
	if propertiesFilePath == nil {
		g.propertiesFilePath = propertiesFilePath
		return g
	}
	fp := *propertiesFilePath
	fp = filepath.ToSlash(strings.Trim(fp, " "))
	if !filepath.IsAbs(fp) {
		panic(perrors.Errorf("配置文件必须为绝对路径"))
	}
	if !strings.HasSuffix(fp, common.C1) {
		fp = fp + common.C1
	}
	fmt.Println("Properties File Path: ", fp)
	g.propertiesFilePath = &fp
	return g
}

func (g *Engine) RegisterTemplate(t DbTemplate) *Engine {
	err := g.register(t.TemplateName(), t)
	if err != nil {
		panic(err)
	}
	return g
}

type TemplateRegistry struct {
	sync.RWMutex

	registry map[string]*structInfo
}

type structInfo struct {
	goType   reflect.Type
	goName   string
	instance interface{}

	methods map[string]*methodInfo
}

type methodInfo struct {
	Name        string
	Method      reflect.StructField
	ArgsType    []reflect.Type
	RepliesType []reflect.Type
}

type DbTemplate interface {
	TemplateName() string
}

func (g *Engine) register(templateName string, obj interface{}) error {
	templateRegistry := g.templateRegistry
	templateRegistry.Lock()
	defer templateRegistry.Unlock()

	if _, ok := templateRegistry.registry[templateName]; ok {
		return perrors.Errorf("You have already registed [%s].", templateName)
	}

	goType := getGoType(obj)
	goName := getGoName(goType)
	methods, err := getGoMethods(goType)
	if err != nil {
		return err
	}
	sqlMap, err := ParseMarkdown(g.templateDir, templateName)
	if err != nil {
		return err
	}
	g.ImplementFunc(templateName, obj, methods, sqlMap)

	templateRegistry.registry[templateName] = &structInfo{
		instance: obj,
		goType:   goType,
		goName:   goName,
		methods:  methods,
	}
	return nil
}

func getGoType(obj interface{}) reflect.Type {
	t := reflect.TypeOf(obj)
	for reflect.Ptr == t.Kind() {
		t = t.Elem()
	}
	return t
}

func getGoName(t reflect.Type) string {
	pkgPath := t.PkgPath()
	if pkgPath == "" {
		return t.Name()
	}
	return pkgPath + "/" + t.Name()
}

func getGoMethods(t reflect.Type) (map[string]*methodInfo, error) {
	methods := make(map[string]*methodInfo)
	for index := 0; index < t.NumField(); index++ {
		method := t.Field(index)
		if reflect.Func != method.Type.Kind() {
			continue
		}
		info, err := getGoMethod(method)
		if err != nil {
			return nil, err
		}
		if info != nil {
			methods[info.Name] = info
		}
	}
	return methods, nil
}

func getGoMethod(method reflect.StructField) (*methodInfo, error) {
	mType := method.Type
	mName := method.Name
	inNum := mType.NumIn()
	outNum := mType.NumOut()

	if mName == common.DbTemplateFuncName {
		return nil, nil
	}

	var (
		argsType    = make([]reflect.Type, 1)
		repliesType = make([]reflect.Type, 2)
	)

	if inNum != 1 {
		panic(perrors.Errorf("Method入参应为(params engine.RenderParamsI)"))
	}

	argsValue := mType.In(0)

	if argsValue.String() != "engine.RenderParamsI" {
		panic(perrors.Errorf("Method入参应为(params engine.RenderParamsI)"))
	}

	argsType[0] = argsValue

	if outNum != 2 {
		panic(perrors.Errorf("Method出参应为([]common.SqlResultI, error)"))
	}

	replyValue := mType.Out(0)
	replyError := mType.Out(1)

	_, isErrorType := reflect.New(replyError).Interface().(*error)
	if reflect.Slice != replyValue.Kind() ||
		replyValue.String() != "[]common.SqlResultI" ||
		!isErrorType {
		panic(perrors.Errorf("Method出参应为([]common.SqlResultI, error)"))
	}

	repliesType[0] = replyValue
	repliesType[1] = replyError

	return &methodInfo{
		Method:      method,
		ArgsType:    argsType,
		RepliesType: repliesType,
		Name:        mName,
	}, nil
}
