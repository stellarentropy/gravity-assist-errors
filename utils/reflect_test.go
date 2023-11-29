package utils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// testStruct serves as a container for metadata and values in the context of
// unit testing, particularly for validating the functionality of tag-related
// operations. It provides a means to ensure accurate reflection-based
// processing of struct field tags, facilitating tests that require tag
// examination and verification.
type testStruct struct {
	Foo string `foo:"bar"`
}

// TestGetStructTag verifies that a struct tag for a specified field is
// retrieved correctly and matches an expected value. It uses reflection to
// ensure that the metadata associated with the field is accurate. The function
// is designed for unit testing scenarios where struct tag functionality needs
// to be confirmed. If the field or its tag does not meet expectations, the
// function reports an error through the [*testing.T] instance.
func TestGetStructTag(t *testing.T) {
	s := &testStruct{}
	field, ok := reflect.TypeOf(*s).FieldByName("Foo")
	if !ok {
		t.Error("field not found")
	}

	assert.Equal(t, GetStructTag(field, "foo"), "bar")
}
