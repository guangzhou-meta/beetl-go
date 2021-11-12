package core

type NativeNodeI interface {
	GetName() string
}

type NativeNode struct {
}

func NewNativeNode() *NativeNode {
	return &NativeNode{}
}

func (nn *NativeNode) GetName() string {
	panic("not implement")
}

type InstanceNode struct {
	*NativeNode
	ref *VarRef
}

func NewInstanceNode(ref *VarRef) *InstanceNode {
	return &InstanceNode{
		NativeNode: NewNativeNode(),
		ref:        ref,
	}
}

func (nn *InstanceNode) GetName() string {
	return nn.ref.Token.Text
}

type ClassNode struct {
	*NativeNode
	clazz string
}

func NewClassNode(clazz string) *ClassNode {
	return &ClassNode{
		NativeNode: NewNativeNode(),
		clazz:      clazz,
	}
}

func (nn *ClassNode) GetName() string {
	return nn.clazz
}
