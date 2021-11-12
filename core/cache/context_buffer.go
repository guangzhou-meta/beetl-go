package cache

import (
	"fmt"
)

var (
	EmptyCharArray         = make([]rune, 0)
	EmptyByteArray         = make([]byte, 0)
	ResizeFactor   float32 = 1.2
	DefaultSize            = 256
)

type ContextBuffer struct {
	maxSize    int
	inner      bool
	charBuffer []rune
	byteBuffer []byte
}

func NewContextBuffer(maxSize int, inner *bool) *ContextBuffer {
	if maxSize < DefaultSize {
		panic(fmt.Sprintf("buffer期望设置需要大于 %d", DefaultSize))
	}
	inst := &ContextBuffer{
		inner:   true,
		maxSize: maxSize,
	}
	if inner != nil {
		inst.inner = *inner
	}
	return inst
}

func (b *ContextBuffer) GetCharBuffer() []rune {
	return b.charBuffer
}

func (b *ContextBuffer) GetByteBuffer() []byte {
	return b.byteBuffer
}

func (b *ContextBuffer) GetCharBufferExpected(expected int) []rune {
	if len(b.charBuffer) >= expected {
		return b.charBuffer
	} else if expected < b.maxSize {
		b.charBuffer = make([]rune, int(float32(expected)*ResizeFactor))
		return b.charBuffer
	}
	return EmptyCharArray
}

func (b *ContextBuffer) GetByteBufferExpected(expected int) []byte {
	if len(b.byteBuffer) >= expected {
		return b.byteBuffer
	} else if expected < b.maxSize {
		b.byteBuffer = make([]byte, int(float32(expected)*ResizeFactor))
		return b.byteBuffer
	}
	return EmptyByteArray
}
