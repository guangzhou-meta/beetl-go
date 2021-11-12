package core

import (
	"github.com/guangzhou-meta/beetl-go/parser"
)

import (
	"github.com/guangzhou-meta/antlr4/runtime/Go/antlr"
)

var (
	StrictDisableGrammars = []string{
		"Arth",
		"ClassNativeCall",
		"Compare",
		"Function",
		"IncDec",
		"InstanceNativeCall",
		"VarAssign",
		"VarRefAssign",
		"VarRefAssignExp",
	}
)

type DefaultTemplateEngine struct {
	antlrErrorStrategy *BeetlAntlrErrorStrategy
	syntaxError        *SyntaxErrorListener
	child              TemplateEngineI
}

func NewDefaultTemplateEngine() *DefaultTemplateEngine {
	inst := &DefaultTemplateEngine{
		antlrErrorStrategy: NewBeetlAntlrErrorStrategy(),
		syntaxError:        NewSyntaxErrorListener(),
	}
	return inst
}

func (e *DefaultTemplateEngine) CreateProgram(resource ResourceI, reader string, textMap map[int]string, cr string, gt *GroupTemplate) *Program {
	//lexer := parser.NewBeetlLexer(antlr.NewInputStream(reader))
	//s:="<$0>><<page('*')>><$1>>\nif(!isEmpty(name)){\n<$2>><<name>><$3>>\n}"
	lexer := parser.NewBeetlLexer(antlr.NewInputStream(reader))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(e.syntaxError)

	tokens := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewBeetlParser(tokens)
	parser.SetErrorHandler(e.antlrErrorStrategy)
	tree := parser.Prog()
	program := NewProgram()
	program.ResourceId = resource.GetId()
	ec := e.child
	if ec == nil {
		ec = e
	}
	program.MetaData = ec.getAntlrBuilder(gt).Build(tree, resource)
	program.MetaData.LineSeparator = cr
	program.MetaData.StaticTextArray = make([]interface{}, len(textMap))

	conf := gt.conf
	directByteOutput := conf.directByteOutput
	for index, v := range textMap {
		var vv interface{}
		if directByteOutput {
			vv = []byte(v)
		} else {
			vv = []rune(v)
		}
		program.MetaData.StaticTextArray[index] = vv
	}
	if program.MetaData.ajaxs != nil {
		for _, ajax := range program.MetaData.ajaxs {
			metaData := ajax.GetLocalProgramMetaData()
			metaData.StaticTextArray = program.MetaData.StaticTextArray
			metaData.LineSeparator = cr
		}
	}
	return program
}

func (e *DefaultTemplateEngine) getAntlrBuilder(groupTemplate *GroupTemplate) *AntlrProgramBuilder {
	return NewAntlrProgramBuilder(groupTemplate, e.getGrammarCreator(groupTemplate))
}

func (e *DefaultTemplateEngine) getGrammarCreator(groupTemplate *GroupTemplate) GrammarCreatorI {
	result := NewGrammarCreator()
	e.setStrictDisableGrammars(result, groupTemplate)
	return result
}

func (e *DefaultTemplateEngine) setStrictDisableGrammars(grammarCreator *GrammarCreator, groupTemplate *GroupTemplate) {
	if groupTemplate.conf.isStrict {
		for _, disableGrammar := range StrictDisableGrammars {
			grammarCreator.AddDisableGrammar(disableGrammar)
		}
	}
}

func (e *DefaultTemplateEngine) inject(engine TemplateEngineI) {
	e.child = engine
}
