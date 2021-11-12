package core

const (

	/** 不支持 */
	DO_NOT_SUPPORT = "DO_NOT_SUPPORT"
	/** 变量已经定义过了 */
	VAR_ALREADY_DEFINED = "VAR_ALREADY_DEFINED"
	/** AJAX 已经定义过了 */
	AJAX_ALREADY_DEFINED = "AJAX_ALREADY_DEFINED"
	/** AJAX 未找到 */
	AJAX_NOT_FOUND = "AJAX_NOT_FOUND"
	/** AJAX 属性定义错误 */
	AJAX_PROPERTY_ERROR = "AJAX_PROPERTY_ERROR"
	/** 变量只读: 通过template.set()的变量，即用于渲染模板的变量都是只读的，不能在模板中被改写 */
	VAR_READONLY = "VAR_READONLY"
	/** 模板引用未定义的变量 */
	VAR_NOT_DEFINED = "VAR_NOT_DEFINED"
	/** 函数未定义 */
	FUNCTION_NOT_FOUND = "FUNCTION_NOT_FOUND"
	/** 函数定义无效 */
	FUNCTION_INVALID = "FUNCTION_INVALID"
	/** 标签未找到 */
	TAG_NOT_FOUND = "TAG_NOT_FOUND"
	/** 虚拟属性未找到 */
	VIRTUAL_NOT_FOUND = "VIRTUAL_NOT_FOUND"
	/** 标签实例化错误 */
	TAG_INSTANCE_ERROR = "TAG_INSTANCE_ERROR"
	/** 本地(Java)调用错误，如 classnotfound, methodnotfound 等 */
	NATIVE_CALL_INVALID = "NATIVE_CALL_INVALID"
	/** 被调用的 class 抛出了异常 */
	NATIVE_CALL_EXCEPTION = "NATIVE_CALL_EXCEPTION"
	/** 不允许本地调用 */
	NATIVE_SECUARITY_EXCEPTION = "NATIVE_SECUARITY_EXCEPTION"
	/** 格式化函数未找到 */
	FORMAT_NOT_FOUND = "FORMAT_NOT_FOUND"
	/** 默认的格式化函数未找到 */
	DEFAULT_FORMAT_NOT_FOUND = "DEFAULT_FORMAT_NOT_FOUND"
	/** 引用属性失效 */
	ATTRIBUTE_INVALID = "ATTRIBUTE_INVALID"
	/** 属性未找到 */
	ATTRIBUTE_NOT_FOUND = "ATTRIBUTE_NOT_FOUND"
	/** 属性不存在，但可能是私有 */
	ATTRIBUTE_NOT_FOUND_PRIVATE = "ATTRIBUTE_NOT_FOUND_PRIVATE"
	/** 模板里使用了[]指示应该是 Map 或者 List，但实际上不是此类型 */
	CAST_LIST_OR_MAP_ERROR = "CAST_LIST_OR_MAP_ERROR"
	/** NULL 标识 */
	NULL = "NULL"
	/** 除 0 错误 */
	DIV_ZERO_ERROR = "DIV_ZERO_ERROR"
	/** 数组 index 错 */
	ARRAY_INDEX_ERROR = "ARRAY_INDEX_ERROR"
	/** 必须是数组 */
	ARRAY_TYPE_ERROR = "ARRAY_TYPE_ERROR"
	/** 期望表达式返回 bool 值，但却是其他类型 */
	BOOLEAN_EXPECTED_ERROR = "BOOLEAN_EXPECTED_ERROR"
	/** 期望是个集合，数组，Iterator，Iterable类型，但却是其他类型 */
	COLLECTION_EXPECTED_ERROR = "COLLECTION_EXPECTED_ERROR"
	/** 期望是数字类型，但却是其他类型 */
	NUMBER_EXPECTED_ERROR = "NUMBER_EXPECTED_ERROR"
	/** 严格MVC错误，模板使用严格MVC，但模板内容含有不允许的语法 */
	STRICK_MVC = "STRICK_MVC"
	/** 不允许调用本地方法 */
	NATIVE_NOT_ALLOWED = "NATIVE_NOT_ALLOWED"
	/** 对象的属性get方法出错 */
	GET_CALL_ERROR = "GET_CALL_ERROR"
	/** has函数调用错误 */
	HAS_CALL_ILLEGAL = "HAS_CALL_ILLEGAL"
	/** 语法错误 */
	ERROR = "ERROR"
	/** Token 错误 */
	TOKEN_ERROR = "TOKEN_ERROR"
	/** 解析时产生的未知错误 */
	PARSER_UNKNOW_ERROR = "PARSER_UNKNOW_ERROR"
	/** 解析时产生语法错误 */
	PARSER_VIABLE_ERROR = "PARSER_VIABLE_ERROR"
	/** 解析时缺少符号 */
	PARSER_MISS_ERROR = "PARSER_MISS_ERROR"
	/** 解析时产生的谓语错误 */
	PARSER_PREDICATE_ERROR = "PARSER_PREDICATE_ERROR"
	/** HTML TAG 解析出错 */
	PARSER_HTML_TAG_ERROR = "PARSER_HTML_TAG_ERROR"
	/** 类型识别错误 */
	TYPE_SEARCH_ERROR = "TYPE_SEARCH_ERROR"
	/** 本地调用格式错误,如a[0](1,2) */
	PARSER_NATIVE_ERROR = "PARSER_NATIVE_ERROR"
	/** 模板加载失败 */
	TEMPLATE_LOAD_ERROR = "TEMPLATE_LOAD_ERROR"
	/** Client IO */
	CLIENT_IO_ERROR_ERROR = "CLIENT_IO_ERROR_ERROR"
	/** 表达式类型不一致，无法运算,如俩个字符串相乘 */
	EXPRESSION_NOT_COMPATIBLE = "EXPRESSION_NOT_COMPATIBLE"
	/** 不允许的语法 */
	GRAMMAR_NOT_ALLOWED = "GRAMMAR_NOT_ALLOWED"

	DEFAULT_INITIAL_CAPACITY = 3
)

type BeetlException struct {
	detailCode string
	msg        *string
	token      *GrammarToken
}

func (e *BeetlException) SetToken(newToken *GrammarToken) *BeetlException {
	e.token = newToken
	return e
}

func (e *BeetlException) Error() string {
	if e.msg != nil {
		return *e.msg
	}
	return e.detailCode
}

func NewBeetlException(detailCode string, msg *string) *BeetlException {
	return &BeetlException{
		detailCode: detailCode,
		msg:        msg,
	}
}
