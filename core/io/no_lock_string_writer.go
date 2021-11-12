package io

import (
	"math"
)

import (
	"github.com/guangzhou-meta/go-lib/arrays"
)

const (
	DefaultBufferSize = 64
)

type NoLockStringWriter struct {
	buf []byte
	pos int
}

func NewNoLockStringWriter() *NoLockStringWriter {
	inst := &NoLockStringWriter{
		buf: make([]byte, DefaultBufferSize),
	}
	return inst
}

func (w *NoLockStringWriter) Write(p []byte) (n int, err error) {
	return w.WriteOfRange(p, 0, len(p))
}

func (w *NoLockStringWriter) WriteOfRange(src []byte, off int, length int) (int, error) {
	newPos := w.pos + length
	if newPos > len(w.buf) {
		buf := make([]byte, int(math.Max(float64(len(w.buf)<<1), float64(newPos))))
		arrays.CopyOf(w.buf, len(buf), buf)
		w.buf = buf
	}
	arrays.CopyFrom(src, off, w.buf, w.pos, length)
	w.pos = newPos
	return length, nil
}

func (w *NoLockStringWriter) WriteString(str string) (int, error) {
	return w.Write([]byte(str))
}

func (w *NoLockStringWriter) String() string {
	return string(w.buf[0:w.pos])
}
