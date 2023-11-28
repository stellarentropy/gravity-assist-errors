package errors

import (
	"fmt"
)

// ErrInvalidEnv represents an error that occurs when an environment variable
// fails to meet the required format or constraints for the application's
// correct operation, indicating non-compliance with the necessary criteria.
var ErrInvalidEnv = fmt.Errorf("invalid environment variable")

// ErrInvalidPath represents an error that occurs when a provided path is not
// valid or does not adhere to the expected format, structure, or syntax
// necessary for proper processing or functioning within an application.
var ErrInvalidPath = fmt.Errorf("invalid path")

// ErrMissingEnv represents the error condition where a required environment
// variable is not present in the environment. This typically indicates that the
// application configuration or execution cannot proceed without this variable
// being set.
var ErrMissingEnv = fmt.Errorf("missing environment variable")
