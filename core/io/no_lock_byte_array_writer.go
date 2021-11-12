package io

import (
	"bufio"
	"bytes"
	"reflect"
)

import (
	"github.com/guangzhou-meta/beetl-go/util"
)

type NoLockByteArrayWriter struct {
	*bufio.Writer

	p int
}

func NewNoLockByteArrayWriter() *NoLockByteArrayWriter {
	return &NoLockByteArrayWriter{
		Writer: bufio.NewWriter(bytes.NewBuffer(nil)),
	}
}

func (w *NoLockByteArrayWriter) GetBuf() []byte {
	v := util.UnpackValue(reflect.ValueOf(w))
	b := util.GetPrivateField(v, "buf").([]byte)
	return b
}

func (w *NoLockByteArrayWriter) GetPos() int {
	v := util.UnpackValue(reflect.ValueOf(w))
	b := util.GetPrivateField(v, "n").(int)
	return b
}

func (w *NoLockByteArrayWriter) String() string {
	return string(w.GetBuf()[:w.GetPos()])
}

func (w *NoLockByteArrayWriter) WriteString(str string) (i int, err error) {
	i, err = w.Writer.WriteString(str)
	if err == nil {
		w.p += i
	}
	return
}
