package utils

import (
	"regexp"
	"testing"

	guuid "github.com/google/uuid"
)

// BenchmarkNewUUID measures the performance of generating new UUIDs. It
// executes a specified number of iterations to create UUIDs using the NewUUID
// function as defined by the benchmarking framework. This benchmark is
// instrumental for evaluating the speed and efficiency of UUID creation within
// performance tests.
func BenchmarkNewUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewUUID()
	}
}

// BenchmarkNewGoogleUUID measures the performance of generating new Google
// UUIDs without the benefit of random number generation pooling. It disables
// the pool, resets the benchmark timer, and runs a loop that generates UUIDs
// for a predetermined number of iterations, as specified by the benchmarking
// framework. This helps in understanding the performance impact of not using a
// pooled random number generator during UUID creation.
func BenchmarkNewGoogleUUID(b *testing.B) {
	guuid.DisableRandPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewGoogleUUID()
	}
}

// BenchmarkNewGoogleUUIDWithRandPool measures the performance of generating new
// Google UUIDs when the random number generator pool is enabled. It benchmarks
// execution time and efficiency using a loop to create UUIDs for a
// predetermined number of iterations specified by the benchmarking framework,
// with the intention of evaluating the impact of the random pool on
// performance.
func BenchmarkNewGoogleUUIDWithRandPool(b *testing.B) {
	guuid.EnableRandPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewGoogleUUID()
	}
}

// BenchmarkNewKSUID measures the performance of generating new KSUIDs in fast
// mode. It enables fast mode for KSUID generation, resets the benchmark timer
// for accurate timing, and repeatedly invokes NewKSUID for a number of
// iterations determined by the benchmarking tools. This benchmark is used to
// evaluate the speed and efficiency of KSUID creation in this mode.
func BenchmarkNewKSUID(b *testing.B) {
	SetFastKSUID()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewKSUID()
	}
}

// BenchmarkNewSafeKSUID measures the performance of generating KSUIDs in safe
// mode across a number of iterations specified by the testing framework. It
// ensures that the generation process is evaluated in an environment where
// thread safety is prioritized, providing an assessment of generation speed and
// efficiency under these constraints.
func BenchmarkNewSafeKSUID(b *testing.B) {
	SetSafeKSUID()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewKSUID()
	}
}

// TestNewUUID verifies the UUID generated by NewUUID adheres to the standard
// UUID format. It ensures that the produced string matches a specific regular
// expression pattern, confirming that the format of the generated UUID is
// correct. If the generated UUID does not match this pattern, an error is
// reported to indicate a failure in maintaining the expected UUID format.
func TestNewUUID(t *testing.T) {
	expectedPattern := "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	match, _ := regexp.MatchString(expectedPattern, NewUUID())
	if !match {
		t.Errorf("UUID does not match expected pattern.")
	}
}

// TestNewKSUID verifies the correctness of KSUIDs produced by NewKSUID by
// ensuring they conform to an expected format. It compares the generated KSUID
// against a regular expression pattern and reports an error if the format is
// not as anticipated, thus affirming the consistency and validity of the output
// KSUIDs.
func TestNewKSUID(t *testing.T) {
	expectedPattern := "^[0-9A-Za-z]{27}$"
	match, _ := regexp.MatchString(expectedPattern, NewKSUID())
	if !match {
		t.Errorf("KSUID does not match expected pattern.")
	}
}
