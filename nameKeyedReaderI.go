package overlay

// xlOverlay_go/nameKeyedReaderI.go

// This call is synchronous: it blocks
type NameKeyedReaderI interface {
	Get(key string) error
}
