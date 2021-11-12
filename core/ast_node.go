package core

type ASTNodeI interface {
	IVarIndexI

	GetASTNode() *ASTNode
}

type ASTNode struct {
	Token *GrammarToken
}

func NewASTNode(token *GrammarToken) *ASTNode {
	return &ASTNode{
		Token: token,
	}
}

func (n *ASTNode) GetASTNode() *ASTNode {
	return n
}

func (n *ASTNode) SetVarIndex(index int) {
	panic("not implement")
}

func (n *ASTNode) GetVarIndex() int {
	panic("not implement")
}
