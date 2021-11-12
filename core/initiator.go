package core

type EngineFunction struct {
	debugEnable    bool
	assertEnable   bool
	registerFunMap map[string]FunctionI
	registerFmtMap map[string]FormatI
}

func NewEngineFunction() *EngineFunction {
	inst := &EngineFunction{
		debugEnable:    true,
		assertEnable:   true,
		registerFunMap: make(map[string]FunctionI),
		registerFmtMap: make(map[string]FormatI),
	}
	inst.init()
	return inst
}

func (f *EngineFunction) init() {
	f.RegisterFunction(&DateFunction{})
	f.RegisterFunction(&NVLFunction{})
	f.RegisterFunction(&DebugFunction{})
	f.RegisterFunction(&CheckExistFunction{})
	f.RegisterFunction(&PrintfFunction{})
	f.RegisterFunction(&DecodeFunction{})
	f.RegisterFunction(&AssertFunction{})
	f.RegisterFunction(&PrintFunction{})
	f.RegisterFunction(&PrintlnFunction{})
	f.RegisterFunction(&PrintFileFunction{})
	f.RegisterFunction(&TruncFunction{})
	f.RegisterFunction(&TruncFunction2{})
	f.RegisterFunction(&QuestionMarkFunction{})
	f.RegisterFunction(&EmptyExpressionFunction{})
	f.RegisterFunction(&IsNotEmptyExpressionFunction{})
	f.RegisterFunction(&ParseIntFunction{})
	f.RegisterFunction(&ParseLongFunction{})
	f.RegisterFunction(&ParseDoubleFunction{})
	f.RegisterFunction(&RangeFunction{})
	f.RegisterFunction(&FlushFunction{})
	f.RegisterFunction(&JsonFunction{})
	f.RegisterFunction(&PageContextFunction{})
	f.RegisterFunction(&TypeNewFunction{})
	f.RegisterFunction(&TypeNameFunction{})
	f.RegisterFunction(&DynamicGlobalValueFunction{})
	f.RegisterFunction(&AllGlobaAsJsonlFunction{})
	f.RegisterFunction(&HasAttributeFunction{})
	f.RegisterFunction(&EnvFunction{})
	f.RegisterFunction(&ParentTagFunction{})
	f.RegisterFunction(&PageQueryFunction{})
	f.RegisterFunction(&IsBlank{})
}

func (f *EngineFunction) EnableDebug(b bool) bool {
	f.debugEnable = b
	return b
}

func (f *EngineFunction) EnableAssert(b bool) bool {
	f.assertEnable = b
	return b
}

func (f *EngineFunction) RegisterFunction(fun FunctionI) {
	clsName := fun.ClassName()
	f.registerFunMap[clsName] = fun
}

func (f *EngineFunction) UnregisterFunction(fun FunctionI) {
	clsName := fun.ClassName()
	delete(f.registerFunMap, clsName)
}

func (f *EngineFunction) RegisterFormat(fun FormatI) {
	clsName := fun.ClassName()
	f.registerFmtMap[clsName] = fun
}

func (f *EngineFunction) UnregisterFormat(fun FormatI) {
	clsName := fun.ClassName()
	delete(f.registerFmtMap, clsName)
}

type EngineTag struct {
	registerTagMap        map[string]TagI
	registerTagFactoryMap map[string]TagFactoryI
}

func NewEngineTag() *EngineTag {
	inst := &EngineTag{
		registerTagMap:        make(map[string]TagI),
		registerTagFactoryMap: make(map[string]TagFactoryI),
	}
	inst.init()
	return inst
}

func (t *EngineTag) init() {
	t.RegisterTag(NewPageQueryIgnoreTag())
	t.RegisterTag(NewPageQueryTag())
}

func (t *EngineTag) RegisterTag(fun TagI) {
	clsName := fun.ClassName()
	t.registerTagMap[clsName] = fun
}

func (t *EngineTag) UnregisterTag(fun TagI) {
	clsName := fun.ClassName()
	delete(t.registerTagMap, clsName)
}
