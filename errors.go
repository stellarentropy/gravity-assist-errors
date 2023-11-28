package errors

import (
	errs "errors"
)

// New creates a new error with the specified message and returns an [error]. If
// the message is empty, the returned error will have an empty message string.
func New(msg string) error {
	return errs.New(msg)
}

// Wrap consolidates a sequence of [error]s into a single [error], streamlining
// error handling by combining multiple errors into one. It returns nil if there
// is no error to wrap, indicated by the last [error] in the sequence being nil.
func Wrap(errl ...error) error {
	if errl[len(errl)-1] == nil {
		return nil
	}

	return errs.Join(errl...)
}

// Is reports whether any [error] in err's chain matches target, considering
// wrapped errors and allowing for comparison of the underlying error values.
func Is(err, target error) bool {
	return errs.Is(err, target)
}

// As attempts to assign an [error] to a target of any interface type, checking
// if the error can be represented as the target's type. It returns true if the
// assignment is successful or false otherwise. The target must be a non-nil
// pointer to the type that may represent the [error]. If the target is not a
// non-nil pointer or does not conform to the appropriate type, a panic will
// occur.
func As(err error, target any) bool {
	return errs.As(err, target)
}
