package core

type GeneralForStatement struct {
	*Statement
	*IGoto

	expInit      []ExpressionI
	condition    ExpressionI
	expUpdate    []ExpressionI
	forPart      StatementI
	elseforPart  StatementI
	varAssignSeq *VarAssignSeqStatement
	hasGoto      bool
	itType       int
}

func NewGeneralForStatement(varAssignSeq *VarAssignSeqStatement, expInit []ExpressionI, condition ExpressionI, expUpdate []ExpressionI, forPart StatementI, elseforPart StatementI, token *GrammarToken) *GeneralForStatement {
	return &GeneralForStatement{
		Statement:    NewStatement(token),
		varAssignSeq: varAssignSeq,
		expInit:      expInit,
		condition:    condition,
		expUpdate:    expUpdate,
		elseforPart:  elseforPart,
		forPart:      forPart,
	}
}

func (sf *GeneralForStatement) Execute(ctx *Context) {
	if sf.expInit != nil {
		for _, exp := range sf.expInit {
			exp.Evaluate(ctx)
		}
	}
	if sf.varAssignSeq != nil {
		sf.varAssignSeq.Execute(ctx)
	}

	for {
		val := sf.condition.Evaluate(ctx)
		b := false
		if c, ok := val.(bool); ok {
			b = c
		} else {
			panic(NewBeetlException(BOOLEAN_EXPECTED_ERROR, nil).SetToken(sf.condition.GetASTNode().Token))
		}
		if b {
			sf.forPart.Execute(ctx)
			switch ctx.GotoFlag {
			case IgotoNormal:
				break
			case IgotoContinue:
				ctx.GotoFlag = IgotoNormal
				break
			case IgotoReturn:
				return
			case IgotoBreak:
				ctx.GotoFlag = IgotoNormal
				return
			}
		} else {
			break
		}
		if sf.expUpdate != nil {
			for _, exp := range sf.expUpdate {
				exp.Evaluate(ctx)
			}
		}
	}
}

func (sf *GeneralForStatement) HasGoto() bool {
	return sf.hasGoto
}

func (sf *GeneralForStatement) SetGoto(hasGoto bool) {
	sf.hasGoto = hasGoto
}
