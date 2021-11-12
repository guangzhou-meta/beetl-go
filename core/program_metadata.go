package core

type ProgramMetaData struct {
	LineSeparator             string
	statements                []StatementI
	StaticTextArray           []interface{}
	varIndexSize              int
	tempVarStartIndex         int
	HasGoto                   bool
	globalIndexMap            map[string]int
	globalVarAttr             map[string]string
	ajaxs                     map[string]*AjaxStatement
	templateRootScopeIndexMap map[string]int
}

func NewProgramMetaData() *ProgramMetaData {
	inst := &ProgramMetaData{
		globalIndexMap:            make(map[string]int),
		globalVarAttr:             make(map[string]string),
		templateRootScopeIndexMap: make(map[string]int),
	}
	return inst
}

func (p *ProgramMetaData) InitContext(ctx *Context) {
	ctx.StaticTextArray = p.StaticTextArray
	ctx.TempVarStartIndex = p.tempVarStartIndex
	ctx.Vars = make([]interface{}, p.varIndexSize+1)
	ctx.Vars[p.varIndexSize] = NotExistObject
	p.putGlobalToArray(ctx)
}

func (p *ProgramMetaData) putGlobalToArray(ctx *Context) {
	globalVar := ctx.GlobalVar
	if globalVar == nil {
		for i := 0; i < p.tempVarStartIndex; i++ {
			ctx.Vars[i] = NotExistObject
		}
		return
	}
	for key, index := range p.globalIndexMap {
		if v, ok := globalVar[key]; ok {
			ctx.Vars[index] = v
		} else {
			ctx.Vars[index] = NotExistObject
		}
	}
}

func (p *ProgramMetaData) GetTemplateRootScopeIndexMap() map[string]int {
	return p.templateRootScopeIndexMap
}

func (p *ProgramMetaData) SetTemplateRootScopeIndexMap(templateRootScopeIndexMap map[string]int) {
	p.templateRootScopeIndexMap = templateRootScopeIndexMap
}

func (p *ProgramMetaData) GetAjax(anchor string) *AjaxStatement {
	if p.ajaxs == nil {
		msg := "该模板文件没有发现任何ajax锚点"
		panic(NewBeetlException(AJAX_NOT_FOUND, &msg).SetToken(NewGrammarToken(anchor, 0, 0)))
	}
	return p.ajaxs[anchor]
}
