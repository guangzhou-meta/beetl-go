package core

import (
	"fmt"
	"io"
)

import (
	"github.com/guangzhou-meta/beetl-go/core/cache"
	io2 "github.com/guangzhou-meta/beetl-go/core/io"
)

type BodyContent interface {
	GetBody() string
	Fill(out ByteWriterI)
}

type ByteBodyContent struct {
	bs    []byte
	count int
}

func NewByteBodyContent(bs []byte, count int) *ByteBodyContent {
	return &ByteBodyContent{
		bs:    bs,
		count: count,
	}
}

func (bc *ByteBodyContent) GetBody() string {
	return bc.String()
}

func (bc *ByteBodyContent) Fill(out ByteWriterI) {
	_, _ = out.Write(bc.bs[:bc.count])
}

func (bc *ByteBodyContent) String() string {
	return string(bc.bs[:bc.count])
}

type ByteWriterI interface {
	io.Writer

	GetTempWriter(parent ByteWriterI) ByteWriterI
	GetTempContent() BodyContent
	Fill(out ByteWriterI)
	Flush()

	WriteString(str string)
	WriteNumberChars(chars []rune, length int)
}

type ByteWriter struct {
	localBuffer *cache.ContextBuffer
	ctx         *Context
	parent      ByteWriterI
}

func NewByteWriter(ctx *Context) *ByteWriter {
	return &ByteWriter{
		ctx:         ctx,
		localBuffer: ctx.LocalBuffer,
	}
}

func (bw *ByteWriter) GetTempWriter(parent ByteWriterI) ByteWriterI {
	panic("not implement")
}

func (bw *ByteWriter) GetTempContent() BodyContent {
	panic("not implement")
}

func (bw *ByteWriter) Fill(out ByteWriterI) {
	panic("not implement")
}

func (bw *ByteWriter) Flush() {
	panic("not implement")
}

func (bw *ByteWriter) WriteString(str string) {
	panic("not implement")
}

func (bw *ByteWriter) WriteNumberChars(chars []rune, length int) {
	panic("not implement")
}

func (bw *ByteWriter) Write(p []byte) (int, error) {
	panic("not implement")
}

func (bw *ByteWriter) WriteContent(bodyContent BodyContent) {
	bodyContent.Fill(bw)
}

func (bw *ByteWriter) WriteAny(d interface{}) {
	var buf []rune
	switch v := d.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		buf = []rune(fmt.Sprintf("%d", v))
	case float32, float64:
		buf = []rune(fmt.Sprintf("%f", v))
	default:
		buf = []rune(fmt.Sprintf("%v", v))
	}
	bw.WriteNumberChars(buf, len(buf))
}

type ByteWriterByte struct {
	*ByteWriter

	os *io2.NoLockByteArrayWriter
	cs string
}

func NewByteWriterByte(os *io2.NoLockByteArrayWriter, cs string, ctx *Context, parent ByteWriterI) *ByteWriterByte {
	inst := &ByteWriterByte{
		ByteWriter: NewByteWriter(ctx),
		os:         os,
		cs:         cs,
	}
	if parent != nil {
		inst.parent = parent
	}
	return inst
}

func (bw *ByteWriterByte) GetTempWriter(parent ByteWriterI) ByteWriterI {
	return NewByteWriterByte(io2.NewNoLockByteArrayWriter(), bw.cs, bw.ctx, parent)
}

func (bw *ByteWriterByte) GetTempContent() BodyContent {
	return NewByteBodyContent(bw.os.GetBuf(), bw.os.GetPos())
}

func (bw *ByteWriterByte) Fill(out ByteWriterI) {
	if bwb, ok := out.(*ByteWriterByte); ok {
		_, _ = bw.Write(bwb.os.GetBuf())
	}
}

func (bw *ByteWriterByte) Flush() {
	//if bw.parent != nil {
	//	bw.parent.Flush()
	//}
	//_ = bw.os.Flush()
}

func (bw *ByteWriterByte) WriteString(str string) {
	_, _ = bw.os.WriteString(str)
}

func (bw *ByteWriterByte) WriteNumberChars(chars []rune, length int) {
	bw.WriteString(string(chars[:length]))
}

func (bw *ByteWriterByte) Write(p []byte) (int, error) {
	bw.WriteString(string(p))
	return len(p), nil
}
