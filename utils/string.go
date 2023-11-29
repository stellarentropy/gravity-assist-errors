package utils

import (
	"strings"
)

// StringInSlice reports whether the provided string is present in the given
// slice of strings. It returns true if the string is found, otherwise false.
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// FastCaseInsensitiveStringCompare compares two strings for equality, ignoring
// case sensitivity. It determines whether the strings are equivalent when case
// is disregarded, provided they have identical length. It returns true if they
// are considered equal under these conditions or false otherwise. This function
// is optimized for quick comparisons.
func FastCaseInsensitiveStringCompare(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	return strings.EqualFold(a, b)
}
