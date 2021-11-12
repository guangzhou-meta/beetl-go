package core

type TemplateEngineI interface {
	CreateProgram(resource ResourceI, input string, textMap map[int]string, cr string, gt *GroupTemplate) *Program

	getAntlrBuilder(groupTemplate *GroupTemplate) *AntlrProgramBuilder

	getGrammarCreator(groupTemplate *GroupTemplate) GrammarCreatorI
}
