package streams

type BaseStream struct {
	fd int
	TtySize [2]int
}
