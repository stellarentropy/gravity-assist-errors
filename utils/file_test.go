package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIsFile verifies the IsFile function by ensuring that a temporary file is
// correctly identified as a file. It reports errors if the file cannot be
// created or does not exist. This test is part of the package's test suite and
// is intended to validate the behavior of IsFile within different file
// scenarios.
func TestIsFile(t *testing.T) {
	tempFile, err := os.CreateTemp(os.TempDir(), "test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Remove(tempFile.Name()) }()

	assert.True(t, IsFile(tempFile.Name()))
}
