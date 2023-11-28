package errors

import (
	"fmt"
)

// ErrRequestHandling represents the occurrence of an error during the
// processing of an HTTP request, indicating that the request could not be
// handled as expected due to either a malfunction in the request handling logic
// or an unforeseen issue arising while fulfilling the request.
var ErrRequestHandling = fmt.Errorf("error handling request")

// ErrResponseWriting indicates an error encountered during the process of
// sending an HTTP response back to the client. It is used to signal failure in
// the response transmission which can arise from various issues such as network
// failures, encoding errors, or interruptions in the data stream to the client.
var ErrResponseWriting = fmt.Errorf("error writing response")

// ErrRequestBodyReading represents an error encountered when trying to read the
// body of an HTTP request. It is used to indicate difficulties in obtaining
// request data, which may arise from network issues, improperly formatted
// input, or a premature end of the data stream.
var ErrRequestBodyReading = fmt.Errorf("error reading request body")
