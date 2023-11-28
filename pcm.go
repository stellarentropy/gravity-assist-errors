package errors

import (
	"fmt"
)

// ErrPCMForwarding represents an error that occurs during the forwarding of
// reports to PCM. This error typically arises from issues with data
// transmission or communication problems with PCM, hindering the successful
// delivery of reports.
var ErrPCMForwarding = fmt.Errorf("error forwarding reports to pcm")

// ErrRequestParsing represents the error that occurs when the format or
// structure of an incoming request cannot be interpreted. This usually means
// the request does not match the expected schema or is malformed, obstructing
// proper processing.
var ErrRequestParsing = fmt.Errorf("error parsing request")

// ErrResponseParsing signifies a failure in interpreting a received response
// when there is a mismatch between the expected data format and the actual
// structure encountered.
var ErrResponseParsing = fmt.Errorf("error parsing response")
