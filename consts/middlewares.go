package consts

// CtxKey serves as a unique identifier for storing and retrieving
// context-specific data throughout middleware operations. It promotes
// consistent and type-safe access to values such as HTTP headers and
// request-related information that are pertinent to the handling of an HTTP
// request's context.
type CtxKey string

const (
	// SignatureHeaderKey is the key for an HTTP header that contains a digital
	// signature. This signature is used to verify the authenticity and integrity of
	// a message within middleware operations.
	SignatureHeaderKey CtxKey = "X-Gravity-Assist-Signature"

	// TokenHeaderKey identifies the HTTP header used to convey an authentication
	// token. Middleware utilizes this key to locate and retrieve the token from
	// incoming requests, which is essential for authenticating the requester's
	// identity.
	TokenHeaderKey CtxKey = "X-Gravity-Assist-Token"

	// PrincipalIdHeaderKey is the key for an HTTP header that identifies the
	// principal (user, application, or service) making the request. It is utilized
	// in middleware to retrieve and utilize the principal ID for processes such as
	// authentication, authorization, and activity tracking.
	PrincipalIdHeaderKey CtxKey = "X-Gravity-Assist-PrincipalId"

	// RequestIdHeaderKey identifies the HTTP header name intended to transport a
	// unique identifier for each request, facilitating the tracing and management
	// of requests within the system.
	RequestIdHeaderKey CtxKey = "X-Gravity-Assist-RequestId"

	// FailureSourceHeaderKey identifies the origin of a failure within the context
	// of HTTP request processing. It serves to convey information about where and
	// why a problem has occurred, enabling error tracking and diagnostics in
	// middleware layers.
	FailureSourceHeaderKey CtxKey = "X-Gravity-Assist-FailureSource"

	// ForwardedHeaderKey represents the name of the HTTP header that indicates
	// whether a request or response has been forwarded. It is often used within
	// middleware to inspect or modify this header.
	ForwardedHeaderKey CtxKey = "X-Gravity-Assist-Forwarded"

	// PersistedHeaderKey indicates if a request has been persisted. It is used in
	// middleware to determine or change the persistence status of a request,
	// providing awareness of whether the request has been stored or recorded.
	PersistedHeaderKey CtxKey = "X-Gravity-Assist-Persisted"

	// ForwardRequestHeaderKey represents the key for the HTTP header that carries
	// information about the origin of a forwarded request. This key is commonly
	// used in middleware operations to identify and possibly modify this header.
	ForwardRequestHeaderKey CtxKey = "X-Gravity-Assist-ForwardRequest"

	// ValidateReportsHeaderKey indicates whether a report has been validated. It is
	// commonly employed in middleware operations to examine or modify the
	// corresponding header.
	ValidateReportsHeaderKey CtxKey = "X-Gravity-Assist-ValidateReports"
)
