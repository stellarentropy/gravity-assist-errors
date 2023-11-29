package utils

import (
	"strings"
	"testing"
)

// BenchmarkCompareOperators benchmarks the performance of string equality
// checks using the '==' operator over a series of iterations dictated by
// [testing.B]. This benchmark is useful for determining the efficiency of
// direct string comparisons in Go.
func BenchmarkCompareOperators(b *testing.B) {
	for n := 0; n < b.N; n++ {
		func() bool {
			return "Test123" == "Test123"
		}()
	}
}

// BenchmarkCompareOperatorsCaseInsensitive measures the performance of
// comparing two strings for equality in a case-insensitive manner by converting
// them to lower case before using the equality operator. It performs this
// comparison multiple times, with the number of iterations determined by the
// [testing.B] parameter. This benchmark is intended to provide insights into
// the efficiency of case-insensitive string comparisons where strings are
// preprocessed to a common case.
func BenchmarkCompareOperatorsCaseInsensitive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		func() bool {
			return strings.ToLower("Test123") == strings.ToLower("Test123")
		}()
	}
}

// BenchmarkCompareString measures the execution time of string comparison using
// the Compare function from the [strings] package. It conducts this benchmark
// by comparing two identical strings for a number of iterations defined by
// [b.N] from the [testing] package, without considering case sensitivity.
func BenchmarkCompareString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strings.Compare("Test123", "Test123")
	}
}

// BenchmarkCompareStringCaseInsensitive assesses the efficiency of performing
// case-insensitive string comparisons by utilizing the Compare function from
// the [strings] package. It iterates the comparison process as per the count
// specified by [testing.B], which is useful for gauging the performance
// implications of such comparisons in Go programs.
func BenchmarkCompareStringCaseInsensitive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strings.Compare(strings.ToLower("Test123"), strings.ToLower("Test123"))
	}
}

// BenchmarkEqualFold assesses the performance of case-insensitive string
// comparisons using the EqualFold function from the [strings] package for a
// number of iterations determined by [testing.B]. It helps in gauging the
// efficiency and speed of this comparison method when executed repeatedly.
func BenchmarkEqualFold(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strings.EqualFold("Test123", "Test123")
	}
}

// BenchmarkFastStringCompare measures the execution efficiency of a custom
// method for case-insensitive string comparison over a series of iterations
// determined by [testing.B]. This benchmark is useful for gauging the
// performance of this particular approach to comparing strings without
// considering their case.
func BenchmarkFastStringCompare(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FastCaseInsensitiveStringCompare("Test123", "Test123")
	}
}
