package consts

const (
	// PCM serves as an identifier for the "PCM" system component in error tracking
	// during private report request processing, aiding in pinpointing the error
	// location within operations like request parsing, forwarding, and response
	// management.
	PCM = "PCM"

	// SE acts as an identifier for the source of errors in private report request
	// processing, aiding in tracing errors across various stages including data
	// processing, marshaling, response reading, uploading to S3, and response
	// building. It is used in response headers for error source tracing.
	SE = "SE"

	// PARTICIPANT represents the entity responsible for initiating a private report
	// request, used to trace the origin of errors that occur during the processing
	// of such requests.
	PARTICIPANT = "PARTICIPANT"
)
