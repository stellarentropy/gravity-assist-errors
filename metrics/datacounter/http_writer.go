package datacounter

import (
	"bufio"
	"context"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/stellarentropy/gravity-assist-common/metrics/tracer"
)

// HTTPWriterCounter wraps an [http.ResponseWriter] to measure and record the
// volume of data written in HTTP responses, the response status codes, and the
// duration of response writing. It also enables hijacking the connection for
// protocol upgrades and modifying response headers. This type is crucial for
// monitoring the performance and efficiency of HTTP response handling in web
// applications.
type HTTPWriterCounter struct {
	http.ResponseWriter
	ctx        context.Context
	count      uint64
	started    time.Time
	statusCode int
	component  string
}

// NewHTTPWriterCounter initializes a new [HTTPWriterCounter] that wraps an
// existing [http.ResponseWriter]. It starts tracking the amount of data sent to
// the client, the response write latency, and sets the initial time for these
// measurements. It returns a pointer to the newly created [HTTPWriterCounter].
func NewHTTPWriterCounter(ctx context.Context, component string, rw http.ResponseWriter) *HTTPWriterCounter {
	return &HTTPWriterCounter{
		ResponseWriter: rw,
		ctx:            ctx,
		started:        time.Now(),
		component:      component,
	}
}

// Write forwards the provided data to the wrapped [http.ResponseWriter] while
// tracking the number of bytes sent and updating related metrics. It returns
// the number of bytes written and any error that occurred during the write
// operation.
func (counter *HTTPWriterCounter) Write(buf []byte) (int, error) {
	n, err := counter.ResponseWriter.Write(buf)

	atomic.AddUint64(&counter.count, uint64(n))

	if err != nil {
		return n, err
	}

	err = tracer.AddInt64(counter.ctx, counter.component, "http.server.bytes.written", int64(n))

	return n, err
}

// Header returns the headers of the HTTP response that will be sent. This
// allows for modification of the header map before writing the response body.
func (counter *HTTPWriterCounter) Header() http.Header {
	return counter.ResponseWriter.Header()
}

// WriteHeader sets the HTTP status code for the current response, records it
// internally for future reference, and delegates the actual header writing to
// the wrapped [http.ResponseWriter].
func (counter *HTTPWriterCounter) WriteHeader(statusCode int) {
	counter.statusCode = statusCode
	counter.ResponseWriter.WriteHeader(statusCode)
}

// Hijack allows a client to take over the underlying TCP connection from the
// HTTP server. This is useful for switching protocols or performing operations
// at a lower level than HTTP. It returns the net.Conn, which represents the raw
// network connection, along with bufio.ReadWriter objects that facilitate
// buffered I/O on that connection. If hijacking isn't supported by the
// underlying ResponseWriter or another error occurs, an error will be returned.
func (counter *HTTPWriterCounter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return counter.ResponseWriter.(http.Hijacker).Hijack()
}

// Count retrieves the total number of bytes that have been written to the HTTP
// response, ensuring access is safe across multiple goroutines.
func (counter *HTTPWriterCounter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}

// Started returns the timestamp marking when the tracking of HTTP response
// write times and data counts commenced for the [HTTPWriterCounter].
func (counter *HTTPWriterCounter) Started() time.Time {
	return counter.started
}

// StatusCode retrieves the HTTP status code that has been recorded for the
// response. If no status code has been explicitly set via WriteHeader, this
// method may return the zero value for an integer.
func (counter *HTTPWriterCounter) StatusCode() int {
	return counter.statusCode
}

// Unwrap provides access to the encapsulated [http.ResponseWriter] used by the
// [HTTPWriterCounter]. It allows for interactions with the original response
// writer without any of the additional functionality provided by the counter,
// such as byte counting and latency tracking.
func (counter *HTTPWriterCounter) Unwrap() http.ResponseWriter {
	return counter.ResponseWriter
}
