package broker

import (
	"bytes"
	"io"
)

func newBufferSize(size int) io.ReadWriter {
	buff := new(bytes.Buffer)
	buff.Grow(size)
	return buff
}
