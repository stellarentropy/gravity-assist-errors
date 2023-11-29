package utils

import "reflect"

// GetStructTag retrieves the value associated with a given tag name from a
// struct field's tag. It takes a [reflect.StructField] and a string
// representing the tag name to search for. If the specified tag is not present,
// an empty string is returned.
func GetStructTag(f reflect.StructField, tagName string) string {
	return f.Tag.Get(tagName)
}
