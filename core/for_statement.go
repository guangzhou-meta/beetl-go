package core

type ForStatement struct {
	*Statement
	*IGoto

	idNode      ExpressionI
	exp         ExpressionI
	forPart     StatementI
	elseforPart StatementI
	hasGoto     bool
	hasSafe     bool
}

func NewForStatement(idNode *VarDefineNode, exp ExpressionI, hasSafe bool, forPart StatementI, elseforPart StatementI, token *GrammarToken) *ForStatement {
	return &ForStatement{
		Statement:   NewStatement(token),
		idNode:      idNode,
		exp:         exp,
		hasSafe:     hasSafe,
		elseforPart: elseforPart,
		forPart:     forPart,
	}
}

func (sf *ForStatement) Execute(ctx *Context) {
	varIndex := sf.idNode.GetVarIndex()
	collection := sf.exp.Evaluate(ctx)
	var it ILoopStatus
	if collection == nil {
		if !(sf.hasSafe || ctx.safeOutput) {
			panic(NewBeetlException(NULL, nil).SetToken(sf.exp.GetASTNode().Token))
		} else {
			it = NewGeneralLoopStatusByList([]interface{}{})
		}
	} else {
		it = GetGeneralLoopIteratorStatus(collection)
	}

	ctx.Vars[varIndex+1] = it

	if sf.hasGoto {
		for it.HasNext() {
			ctx.Vars[varIndex] = it.Next()
			sf.forPart.Execute(ctx)
			switch ctx.GotoFlag {
			case IgotoNormal:
				break
			case IgotoContinue:
				ctx.GotoFlag = IgotoNormal
				continue
			case IgotoReturn:
				return
			case IgotoBreak:
				ctx.GotoFlag = IgotoNormal
				return
			}
		}

		if !it.HasData() && sf.elseforPart != nil {
			sf.elseforPart.Execute(ctx)
		}
	} else {
		for it.HasNext() {
			ctx.Vars[varIndex] = it.Next()
			sf.forPart.Execute(ctx)
		}
		if !it.HasData() && sf.elseforPart != nil {
			sf.elseforPart.Execute(ctx)
		}
	}
}

func (sf *ForStatement) HasGoto() bool {
	return sf.hasGoto
}

func (sf *ForStatement) SetGoto(hasGoto bool) {
	sf.hasGoto = hasGoto
}
