package core

type BlockStatement struct {
	*Statement

	nodes   []StatementI
	hasGoto bool
}

func NewBlockStatement(nodes []StatementI, token *GrammarToken) *BlockStatement {
	inst := &BlockStatement{
		Statement: NewStatement(token),
		nodes:     nodes,
	}
	return inst
}

func (s *BlockStatement) Execute(ctx *Context) {
	if s.hasGoto {
		for _, node := range s.nodes {
			node.Execute(ctx)
			if ctx.GotoFlag != 0 {
				return
			}
		}
	} else {
		for _, node := range s.nodes {
			node.Execute(ctx)
		}
	}
}

func (s *BlockStatement) HasGoto() bool {
	return s.hasGoto
}

func (s *BlockStatement) SetGoto(hasGoto bool) {
	s.hasGoto = hasGoto
}
