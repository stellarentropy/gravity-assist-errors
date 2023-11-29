package utils

import (
	"fmt"
	"testing"
	"time"
)

// mockTime represents a predefined moment in time used as a reference for
// testing time-dependent functionalities.
var mockTime = time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)

// TestGetTimePath verifies that GetTimePath correctly returns a string
// representing a provided time formatted as "YYYY/MM/DD/HH". It checks if the
// function's output matches an expected value and reports discrepancies to
// ensure the time formatting is consistent with requirements.
func TestGetTimePath(t *testing.T) {
	// Setup
	expectedResult := fmt.Sprintf("%d/%02d/%02d/%02d", 2022, time.January, 1, 0)

	// Execution
	actualResult := GetTimePath(mockTime)

	// Assertion
	if actualResult != expectedResult {
		t.Errorf("Expected '%s' but got '%s'", expectedResult, actualResult)
	}
}

// TestGetPreviousTimePath verifies the accuracy of the GetPreviousTimePath
// function in calculating the path for the hour immediately before a given
// time. It checks that the output is correctly formatted as "YYYY/MM/DD/HH" and
// corresponds with the expected result. Any deviation from the expected output
// is considered a test failure.
func TestGetPreviousTimePath(t *testing.T) {
	// Setup
	expectedResult := fmt.Sprintf("%02d/%02d/%02d/%02d", 2021, time.December, 31, 23)

	// Execution
	actualResult := GetPreviousTimePath(mockTime)

	// Assertion
	if actualResult != expectedResult {
		t.Errorf("Expected '%s' but got '%s'", expectedResult, actualResult)
	}
}
