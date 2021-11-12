package core

type TagFactoryI interface {
	CreateTag() TagI
}

type DefaultTagFactory struct {
	tagCls TagI
}

func NewDefaultTagFactory(tagCls TagI) *DefaultTagFactory {
	return &DefaultTagFactory{
		tagCls: tagCls,
	}
}

func (tf *DefaultTagFactory) CreateTag() TagI {
	return tf.tagCls
}
