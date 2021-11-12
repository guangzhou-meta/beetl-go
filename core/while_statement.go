package core

type WhileStatement struct {
	*Statement
	*IGoto

	exp       ExpressionI
	whileBody StatementI
	hasGoto   bool
}

func NewWhileStatement(exp ExpressionI, whileBody StatementI, token *GrammarToken) *WhileStatement {
	inst := &WhileStatement{
		Statement: NewStatement(token),
		exp:       exp,
		whileBody: whileBody,
	}
	if b, ok := inst.whileBody.(*BlockStatement); ok {
		inst.hasGoto = b.hasGoto
	}
	return inst
}

func (sf *WhileStatement) Execute(ctx *Context) {
	if sf.hasGoto {
		for {
			result := sf.exp.Evaluate(ctx)
			if b, ok := result.(bool); ok {
				if b {
					sf.whileBody.Execute(ctx)
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
				} else {
					break
				}
			} else {
				panic(NewBeetlException(BOOLEAN_EXPECTED_ERROR, nil).SetToken(sf.exp.GetASTNode().Token))
			}
		}
	} else {
		for {
			result := sf.exp.Evaluate(ctx)
			if b, ok := result.(bool); ok {
				if b {
					sf.whileBody.Execute(ctx)
				} else {
					break
				}
			} else {
				panic(NewBeetlException(BOOLEAN_EXPECTED_ERROR, nil).SetToken(sf.exp.GetASTNode().Token))
			}
		}
	}
}

func (sf *WhileStatement) HasGoto() bool {
	return sf.hasGoto
}

func (sf *WhileStatement) SetGoto(hasGoto bool) {
	sf.hasGoto = hasGoto
}
