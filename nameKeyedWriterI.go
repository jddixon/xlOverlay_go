package overlay

// xlOverlay_go/nameKeyedWriterI.go

// This call is synchronous: it blocks.
type NameKeyedWriterI interface {
	Delete(key string) error
	Put(key string, buffer []byte) error
}
