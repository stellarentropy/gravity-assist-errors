package errors

import (
	"fmt"
)

var (
	// ErrJWTParsingFailed represents an error that occurs when the parsing of a
	// JSON Web Token fails, generally due to malformed structure or incorrect data
	// within the token itself. This error indicates that the token could not be
	// processed as expected.
	ErrJWTParsingFailed = fmt.Errorf("JWT parsing failed")

	// ErrJWTClaimNotFound represents an error that occurs when a specific claim is
	// expected within a JSON Web Token but is not present. This indicates that the
	// token lacks certain required information.
	ErrJWTClaimNotFound = fmt.Errorf("claim not found")

	// ErrJWTEmptyToken indicates the absence of a JSON Web Token when one was
	// expected, often implying a missing authentication credential in a request.
	ErrJWTEmptyToken = fmt.Errorf("token is empty")

	// ErrSignatureMismatch indicates a discrepancy between the expected signature
	// and the actual signature of a JSON Web Token (JWT), suggesting that the token
	// may have been altered or that there is an error in how the signature was
	// generated. It is commonly encountered during the validation process of a JWT.
	ErrSignatureMismatch = fmt.Errorf("signature mismatch")

	// ErrInvalidSignature represents an error encountered when the signature of a
	// JSON Web Token (JWT) fails to validate during the verification process,
	// indicating potential tampering or corruption of the token.
	ErrInvalidSignature = fmt.Errorf("invalid signature")
)
