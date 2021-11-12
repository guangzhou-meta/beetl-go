package core

type BeetlParserException struct {
	BeetlException
}

func NewBeetlParserException(detailCode string, msg *string) *BeetlParserException {
	i := &BeetlParserException{}
	i.BeetlException = *NewBeetlException(detailCode, msg)
	return i
}
