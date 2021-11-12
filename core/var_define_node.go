package core

type VarDefineNode struct {
	*Expression

	varIndex int
}

func NewVarDefineNode(token *GrammarToken) *VarDefineNode {
	return &VarDefineNode{
		Expression: NewExpression(token),
		varIndex:   0,
	}
}

func (vdn *VarDefineNode) Evaluate(ctx *Context) interface{} {
	return nil
}

func (vdn *VarDefineNode) SetVarIndex(index int) {
	vdn.varIndex = index
}

func (vdn *VarDefineNode) GetVarIndex() int {
	return vdn.varIndex
}
