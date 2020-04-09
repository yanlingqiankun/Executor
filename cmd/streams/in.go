package streams

import (
	"io"
)

type In struct {
	BaseStream
	Input io.Reader
}

func (i *In) GetFd() int {
	return i.fd
}

func NewIn(reader io.Reader) In {
	return In{
		Input:      reader,
	}
}
