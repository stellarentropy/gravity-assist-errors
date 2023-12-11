package datacounter

import (
	"io"
	"sync/atomic"
)

// WriterCounter tracks the cumulative number of bytes successfully written to
// an underlying [io.Writer]. It provides a thread-safe way to monitor write
// operations and retrieve the total byte count at any moment. WriterCounter
// also supports closing the underlying writer if it implements the [io.Closer]
// interface, allowing for proper resource management.
type WriterCounter struct {
	count  uint64
	Writer io.WriteCloser
}

// NewWriterCounter returns a new instance of [WriterCounter] that wraps an
// [io.WriteCloser] for tracking the cumulative number of bytes written. The
// count is accessible at any time and is thread-safe. If the underlying writer
// supports closing, the [WriterCounter] will also be closable.
func NewWriterCounter(w io.WriteCloser) *WriterCounter {
	return &WriterCounter{
		Writer: w,
	}
}

// Write sends a slice of bytes to the wrapped [io.Writer], increments the count
// of total bytes written, and returns the number of bytes written along with
// any error that occurred during the write. It ensures thread-safe updating of
// the byte count and excludes negative values from the count.
func (counter *WriterCounter) Write(buf []byte) (int, error) {
	n, err := counter.Writer.Write(buf)

	// Write() should always return a non-negative `n`.
	// But since `n` is a signed integer, some custom
	// implementation of an io.Writer may return negative
	// values.
	//
	// Excluding such invalid values from counting,
	// thus `if n >= 0`:
	if n >= 0 {
		atomic.AddUint64(&counter.count, uint64(n))
	}

	return n, err
}

// Count retrieves the total number of bytes that have been successfully written
// to the wrapped [io.Writer]. It ensures thread safety and returns the byte
// count as a uint64 value.
func (counter *WriterCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}

// Close finalizes the [WriterCounter] by closing the underlying [io.Writer]. If
// the [io.Writer] does not support the [io.Closer] interface, Close returns an
// error (if any) from closing the underlying writer.
func (counter *WriterCounter) Close() error {
	return counter.Writer.Close()
}
