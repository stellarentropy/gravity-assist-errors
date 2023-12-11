package datacounter

import (
	"io"
	"sync/atomic"
)

// ReaderCounter wraps an [io.Reader] to monitor the number of bytes read during
// its lifetime. It provides functionality to read data, query the accumulated
// byte count, and close the underlying reader when applicable. The count is
// maintained in a thread-safe manner, ensuring accurate reporting even in
// concurrent environments.
type ReaderCounter struct {
	count  uint64
	Reader io.ReadCloser
}

// NewReaderCounter returns a new [ReaderCounter] that wraps an [io.ReadCloser],
// allowing for the monitoring and recording of the total number of bytes read.
func NewReaderCounter(r io.ReadCloser) *ReaderCounter {
	return &ReaderCounter{
		Reader: r,
	}
}

// Read retrieves data from the wrapped [io.Reader], increments the byte count
// of the [ReaderCounter] by the number of bytes read, and returns the amount
// read along with any error encountered during the operation. If the read byte
// count is zero or positive, it is added to the total; negative counts are
// ignored.
func (counter *ReaderCounter) Read(buf []byte) (int, error) {
	n, err := counter.Reader.Read(buf)

	// Read() should always return a non-negative `n`.
	// But since `n` is a signed integer, some custom
	// implementation of an io.Reader may return negative
	// values.
	//
	// Excluding such invalid values from counting,
	// thus `if n >= 0`:
	if n >= 0 {
		atomic.AddUint64(&counter.count, uint64(n))
	}

	return n, err
}

// Count retrieves the current cumulative number of bytes that have been read
// using the [ReaderCounter].
func (counter *ReaderCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}

// Close terminates the reading process from the underlying [io.Reader] of the
// [ReaderCounter]. It closes the reader if it implements the [io.Closer]
// interface. If the reader does not implement [io.Closer], Close has no effect
// and returns nil. Any error encountered during the close operation is returned
// to the caller.
func (counter *ReaderCounter) Close() error {
	return counter.Reader.Close()
}
