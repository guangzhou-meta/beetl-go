package core

type ExpressionI interface {
	ASTNodeI
	Evaluate(ctx *Context) interface{}
}

type Expression struct {
	*ASTNode
}

func NewExpression(token *GrammarToken) *Expression {
	return &Expression{
		ASTNode: NewASTNode(token),
	}
}

func (e *Expression) Evaluate(ctx *Context) interface{} {
	panic("not implement")
}
