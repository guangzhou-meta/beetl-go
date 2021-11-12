package core

import (
	"math"
)

type ProgramBuilderContext struct {
	root           *BlockEnvContext
	current        *BlockEnvContext
	listNodeEval   []interface{}
	globalVar      map[string]*VarDescrption
	varIndexSize   int
	globalIndexMap map[string]int
	isSafeOutput   bool
	rootIndexMap   map[string]int
}

func NewProgramBuilderContext() *ProgramBuilderContext {
	root := NewBlockEnvContext()
	return &ProgramBuilderContext{
		root:           root,
		current:        root,
		listNodeEval:   make([]interface{}, 0),
		globalVar:      make(map[string]*VarDescrption),
		globalIndexMap: make(map[string]int),
		rootIndexMap:   make(map[string]int),
	}
}

func (c *ProgramBuilderContext) EnterBlock() {
	blockVar := NewBlockEnvContext()
	blockVar.SetParent(c.current)
	c.current = blockVar
}

func (c *ProgramBuilderContext) ExitBlock() {
	c.current = c.current.parent
}

func (c *ProgramBuilderContext) AddVarAndPosition(first ASTNodeI) {
	c.AddVar(first.GetASTNode().Token.Text)
	c.SetVarPosition(first.GetASTNode().Token.Text, first)
}

func (c *ProgramBuilderContext) AddRootVarAdnPosition(first ASTNodeI) bool {
	varName := first.GetASTNode().Token.Text
	if c.searchVar(c.root, varName) != nil {
		return true
	}

	varDesc := NewVarDescrption()
	varDesc.SetVarName(varName)
	varDesc.where = append(varDesc.where, first)
	c.root.GetVars()[varName] = varDesc
	return false
}

func (c *ProgramBuilderContext) searchVar(ctx *BlockEnvContext, name string) ASTNodeI {
	if ctx.GetVarDescrption(name) != nil {
		return ctx.GetVarDescrption(name).where[0]
	}
	for _, child := range ctx.blockList {
		node := c.searchVar(child, name)
		if node != nil {
			return node
		}
	}
	return nil
}

func (c *ProgramBuilderContext) AddVar(varName string) {
	varDesc := NewVarDescrption()
	varDesc.SetVarName(varName)
	c.current.GetVars()[varName] = varDesc
}

func (c *ProgramBuilderContext) Contain(varName string) ASTNodeI {
	varDesc := c.current.GetVars()[varName]
	if varDesc == nil {
		return nil
	}
	return varDesc.where[0]
}

func (c *ProgramBuilderContext) SetVarAttr(varName string, attrName string) {
	varDesc := c.FindVar(varName)
	varDesc.attrList[attrName] = nil
}

func (c *ProgramBuilderContext) SetVarPosition(varName string, where ASTNodeI) {
	varDesc := c.FindVar(varName)
	varDesc.where = append(varDesc.where, where)
}

func (c *ProgramBuilderContext) HasDefined(varName string) *GrammarToken {
	scope := c.current
	for scope != nil {
		varDesc := scope.GetVarDescrption(varName)
		if varDesc != nil {
			if n, ok := varDesc.where[0].(ASTNodeI); ok {
				return n.GetASTNode().Token
			}
		}
		scope = scope.parent

	}
	return nil
}

func (c *ProgramBuilderContext) FindVar(varName string) *VarDescrption {
	scope := c.current
	for scope != nil {
		varDesc := scope.GetVarDescrption(varName)
		if varDesc != nil {
			return varDesc
		} else {
			scope = scope.parent
		}
	}

	var desc *VarDescrption
	var ok bool
	if desc, ok = c.globalVar[varName]; !ok {
		desc = NewVarDescrption()
		c.globalVar[varName] = desc
	}
	return desc
}

func (c *ProgramBuilderContext) SetNodeEvalObject(o interface{}) int {
	c.listNodeEval = append(c.listNodeEval, o)
	return len(c.listNodeEval) - 1
}

func (c *ProgramBuilderContext) AnzlyszeGlobal() {
	index := 0
	for key, value := range c.globalVar {
		c.globalIndexMap[key] = index

		for _, node := range value.where {
			node.SetVarIndex(index)
		}
		index++
	}
}

func (c *ProgramBuilderContext) AnzlyszeLocal() {
	c.anzlyze(c.root, len(c.globalVar), true)
}

func (c *ProgramBuilderContext) anzlyze(block *BlockEnvContext, nextIndex int, isRoot bool) {
	for _, vd := range block.vars {
		for _, node := range vd.where {
			node.SetVarIndex(nextIndex)
			if isRoot {
				c.rootIndexMap[vd.GetVarName()] = nextIndex
			}
		}
		nextIndex++
	}
	c.varIndexSize = int(math.Max(float64(c.varIndexSize), float64(nextIndex)))

	for _, subBlock := range block.blockList {
		c.anzlyze(subBlock, nextIndex, false)
		inc := len(subBlock.vars)
		c.varIndexSize = int(math.Max(float64(c.varIndexSize), float64(nextIndex+inc)))
	}
}

type BlockEnvContext struct {
	vars                     map[string]*VarDescrption
	blockList                []*BlockEnvContext
	parent                   *BlockEnvContext
	gotoValue                IGotoI
	canStopContinueBreakFlag bool
}

func NewBlockEnvContext() *BlockEnvContext {
	return &BlockEnvContext{
		vars:      make(map[string]*VarDescrption),
		blockList: make([]*BlockEnvContext, 0),
		gotoValue: IgotoNormal,
	}
}

func (c *BlockEnvContext) GetVars() map[string]*VarDescrption {
	return c.vars
}

func (c *BlockEnvContext) SetVars(vars map[string]*VarDescrption) {
	c.vars = vars
}

func (c *BlockEnvContext) GetBlockList() []*BlockEnvContext {
	return c.blockList
}

func (c *BlockEnvContext) SetBlockList(blockList []*BlockEnvContext) {
	c.blockList = blockList
}

func (c *BlockEnvContext) GetParent() *BlockEnvContext {
	return c.parent
}

func (c *BlockEnvContext) SetParent(parent *BlockEnvContext) {
	c.parent = parent
	parent.blockList = append(parent.blockList, c)
}

func (c *BlockEnvContext) GetVarDescrption(varName string) *VarDescrption {
	return c.vars[varName]
}

type VarDescrption struct {
	varName  string
	attrList map[string]interface{}
	where    []ASTNodeI
}

func NewVarDescrption() *VarDescrption {
	return &VarDescrption{
		attrList: make(map[string]interface{}),
		where:    make([]ASTNodeI, 0),
	}
}

func (this *VarDescrption) GetVarName() string {
	return this.varName
}

func (this *VarDescrption) SetVarName(varName string) {
	this.varName = varName
}

func (this *VarDescrption) GetAttrList() map[string]interface{} {
	return this.attrList
}

func (this *VarDescrption) SetAttrList(attrList map[string]interface{}) {
	this.attrList = attrList
}
