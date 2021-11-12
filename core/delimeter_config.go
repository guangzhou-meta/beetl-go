package core

type DelimeterConfigI interface {
	GetPlaceHolder(id interface{}) *DelimeterHolder
	GetStatement(id interface{}) *DelimeterHolder
}

type DefaultDelimeterConfig struct {
	configuration            *Configuration
	statementDelimeterHolder *DelimeterHolder
	placeDelimeterHolder     *DelimeterHolder
}

func NewDefaultDelimeterConfig(configuration *Configuration) *DefaultDelimeterConfig {
	inst := &DefaultDelimeterConfig{
		configuration:            configuration,
		statementDelimeterHolder: configuration.getScriptDelimeter(),
		placeDelimeterHolder:     configuration.getPlaceHolderDelimeter(),
	}
	return inst
}

func (dc *DefaultDelimeterConfig) GetPlaceHolder(id interface{}) *DelimeterHolder {
	return dc.placeDelimeterHolder
}

func (dc *DefaultDelimeterConfig) GetStatement(id interface{}) *DelimeterHolder {
	return dc.statementDelimeterHolder
}
