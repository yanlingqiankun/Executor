package streams

import "io"

type Out struct {
	BaseStream
	Output io.Writer
}

func NewOut(output io.Writer) Out {
	return Out{
		Output:     output,
	}
}