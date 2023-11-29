package utils

import (
	"fmt"
	"time"
)

// GetTimePath converts a given [time.Time] to its string representation in the
// path-like format "Year/Month/Day/Hour". It ensures each component of the time
// is zero-padded for consistent length, facilitating the creation of
// hierarchical identifiers suitable for organizing data chronologically in
// filesystems or URLs.
func GetTimePath(t time.Time) string {
	t = t.UTC()
	return fmt.Sprintf("%d/%02d/%02d/%02d", t.Year(), t.Month(), t.Day(), t.Hour())
}

// GetPreviousTimePath calculates the time one hour prior to a given [time.Time]
// and returns its path-like string representation formatted as
// "Year/Month/Day/Hour", where each component is zero-padded. This utility is
// often used for generating chronological identifiers for use in hierarchical
// structures like file paths or URLs.
func GetPreviousTimePath(t time.Time) string {
	t = t.UTC()
	t = t.Add(-1 * time.Hour)

	return fmt.Sprintf("%d/%02d/%02d/%02d", t.Year(), t.Month(), t.Day(), t.Hour())
}
