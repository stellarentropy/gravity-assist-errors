package datacounter

import (
	"context"
	"io"
	"sync/atomic"

	"github.com/stellarentropy/gravity-assist-common/metrics/tracer"
)

// UpstreamPCMReaderCounter wraps an [io.Reader] to monitor and record the volume
// of data read and the duration of read operations from an upstream source. It
// provides safe concurrent access for multiple goroutines to track cumulative
// bytes read, supports closing of the underlying reader if it is also an
// [io.Closer], and ensures precise latency measurements are captured. This type
// is essential for observing and analyzing reading performance and overall data
// throughput in real-time data processing scenarios.
type UpstreamPCMReaderCounter struct {
	ctx       context.Context
	count     uint64
	component string
	Reader    io.ReadCloser
}

// NewUpstreamPCMReaderCounter creates and returns a new [UpstreamPCMReaderCounter]
// that wraps an [io.ReadCloser] for the purpose of monitoring and recording the
// amount of data read as well as the time taken to read from an upstream
// source. It facilitates both the accurate collection of metrics and the
// provision of concurrent access without race conditions.
func NewUpstreamPCMReaderCounter(ctx context.Context, component string, r io.ReadCloser) *UpstreamPCMReaderCounter {
	return &UpstreamPCMReaderCounter{
		Reader:    r,
		ctx:       ctx,
		component: component,
	}
}

// Read populates the provided buffer with data from the upstream source,
// concurrently updates the count of total bytes read, and records the read
// latency metrics. It returns the number of bytes read into the buffer and any
// error encountered during the reading process. Negative byte counts, if
// returned by the upstream source, are not counted towards the total.
func (counter *UpstreamPCMReaderCounter) Read(buf []byte) (int, error) {
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

		go func() {
			_ = tracer.AddInt64(counter.ctx, counter.component, "upstream.pcm.bytes.read", int64(n))
		}()
	}

	return n, err
}

// Count retrieves the total number of bytes that have been read from the
// upstream data source. It is safe for concurrent use and returns the byte
// count as a [uint64].
func (counter *UpstreamPCMReaderCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}

// Close finalizes the operations of an [UpstreamPCMReaderCounter]. It terminates
// any ongoing read tracking and closes the underlying data source if it
// implements [io.Closer]. If the data source does not implement [io.Closer],
// Close will have no effect on it. Close returns an error if one occurs during
// the closing of the data source.
func (counter *UpstreamPCMReaderCounter) Close() error {
	return counter.Reader.Close()
}
