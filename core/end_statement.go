package core

type EndStatement struct {
	*Statement
}

func NewEndStatement() *EndStatement {
	return &EndStatement{}
}

func (s *EndStatement) Execute(ctx *Context) {

}
