package core

type StatementI interface {
	Execute(ctx *Context)
}

type Statement struct {
	*ASTNode
}

func NewStatement(token *GrammarToken) *Statement {
	inst := &Statement{
		ASTNode: NewASTNode(token),
	}
	return inst
}

func (*Statement) Execute(ctx *Context) {
	panic("not implement")
}
