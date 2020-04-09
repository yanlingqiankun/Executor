package io

import "io"

type EscapeError struct{}

func (EscapeError) Error() string {
	return "read escape sequence"
}

type escapeProxy struct {
	escapeKeys   []byte
	escapeKeyPos int
	r            io.Reader
}

func NewEscapeProxy(r io.Reader, escapeKeys []byte) io.Reader {
	return &escapeProxy{
		escapeKeys: escapeKeys,
		r:          r,
	}
}

func (r *escapeProxy) Read(buf []byte) (int, error) {
	nr, err := r.r.Read(buf)

	if len(r.escapeKeys) == 0 {
		return nr, err
	}

	preserve := func() {
		nr += r.escapeKeyPos
		preserve := make([]byte, 0, r.escapeKeyPos+len(buf))
		preserve = append(preserve, r.escapeKeys[:r.escapeKeyPos]...)
		preserve = append(preserve, buf...)
		r.escapeKeyPos = 0
		copy(buf[0:nr], preserve)
	}

	if nr != 1 || err != nil {
		if r.escapeKeyPos > 0 {
			preserve()
		}
		return nr, err
	}

	if buf[0] != r.escapeKeys[r.escapeKeyPos] {
		if r.escapeKeyPos > 0 {
			preserve()
		}
		return nr, nil
	}

	if r.escapeKeyPos == len(r.escapeKeys)-1 {
		return 0, EscapeError{}
	}

	r.escapeKeyPos++
	return nr - r.escapeKeyPos, nil
}
