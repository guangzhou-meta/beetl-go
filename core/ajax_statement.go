package core

type AjaxStatement struct {
	*Statement
	block                *BlockStatement
	defaultRender        bool
	ajaxId               string
	localProgramMetaData *ProgramMetaData
}

func NewAjaxStatement(block *BlockStatement, token *GrammarToken, defaultRender bool) *AjaxStatement {
	inst := &AjaxStatement{
		Statement:     NewStatement(token),
		defaultRender: true,
	}
	inst.block = block
	inst.defaultRender = defaultRender
	inst.ajaxId = token.Text
	return inst
}

func (s *AjaxStatement) Execute(ctx *Context) {
	if ctx.Template.AjaxId == nil && (!s.defaultRender) {
		// 渲染整个模板情况下，设置成不渲染
		return
	}
	s.block.Execute(ctx)
}

func (s *AjaxStatement) GetLocalProgramMetaData() *ProgramMetaData {
	return s.localProgramMetaData
}

func (s *AjaxStatement) SetLocalProgramMetaData(localProgramMetaData *ProgramMetaData) {
	s.localProgramMetaData = localProgramMetaData
}
