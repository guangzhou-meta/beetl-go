package core

type IGotoI interface {
	HasGoto() bool
	SetGoto(occur bool)
}

type IGoto int

const (
	IgotoNormal IGoto = iota
	IgotoContinue
	IgotoBreak
	IgotoReturn
)

func (IGoto) HasGoto() bool {
	panic("not implement")
}
func (IGoto) SetGoto(occur bool) {
	panic("not implement")
}

type Program struct {
	MetaData   *ProgramMetaData
	ResourceId interface{}
}

func NewProgram() *Program {
	return &Program{}
}

func (p *Program) Execute(ctx *Context) {
	if p.MetaData.HasGoto {
		p.runWitchGoCheck(p.MetaData.statements, ctx)
	} else {
		p.run(p.MetaData.statements, ctx)
	}
}

func (p *Program) run(statements []StatementI, ctx *Context) {
	for _, statement := range statements {
		if statement != nil {
			statement.Execute(ctx)
		}
	}
}

func (p *Program) runWitchGoCheck(statements []StatementI, ctx *Context) {
	for _, statement := range statements {
		statement.Execute(ctx)
		if ctx.GotoFlag == IgotoReturn {
			return
		}
	}
}
