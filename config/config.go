package config

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/stellarentropy/gravity-assist-common/utils"

	"github.com/stellarentropy/gravity-assist-common/errors"
)

// Env manages and validates environment variables for an application, providing
// methods to fetch and parse values such as strings, integers, booleans, IP
// addresses, URLs, and file paths. It supports setting default values, creating
// directories for file paths, and enforces constraints like integer ranges or
// string option lists. Env methods can induce panics for validation failures or
// missing mandatory variables, facilitating early configuration error
// detection. The design encourages method chaining for concise configuration
// expressions.
type Env struct {
	key   string
	value string
}

// NewEnv creates and returns a new instance of [Env], initializing it with a
// given environment variable key. It fetches the corresponding value from the
// system's environment variables, or an empty string if the variable is not
// present.
func NewEnv(key string) Env {
	return Env{
		key:   key,
		value: os.Getenv(key),
	}
}

// checkRequired ensures that the environment variable associated with the [Env]
// instance is present and not empty. It panics if these conditions are not met,
// indicating a missing required environment variable.
func (e Env) checkRequired() {
	if e.value == "" {
		panic(errors.Wrap(
			errors.ErrMissingEnv,
			fmt.Errorf("environment variable %s is not set", e.key),
		))
	}
}

// GetString retrieves the string value of the environment variable associated
// with the [Env] instance. If the environment variable is not set or its value
// is empty, an empty string is returned.
func (e Env) GetString() string {
	return e.value
}

// GetPort retrieves the port number from the environment variable associated
// with the [Env] instance. It validates that the port is within the acceptable
// range for TCP/UDP ports (1-65535) and if it is not, or if the value is not a
// valid integer, the method will panic.
func (e Env) GetPort() int {
	return e.WithIntInRange(0, 65535).GetInt()
}

// GetAddress retrieves the string representation of an IP address from the
// environment variable associated with the [Env] instance. If the variable's
// value is "localhost", this exact value is returned. For other non-empty
// values, it parses them as IP addresses and returns the string form. If
// parsing fails or the value is empty, a panic is induced with appropriate
// error details.
func (e Env) GetAddress() string {
	if e.value == "localhost" {
		return e.value
	}

	ip := net.ParseIP(e.value)

	if ip == nil {
		panic(errors.Wrap(
			errors.ErrInvalidEnv,
			fmt.Errorf("env: %s", e.key),
			fmt.Errorf("value: %s", e.value),
		))
	}

	return ip.String()
}

// GetBool interprets the associated environment variable's value as a boolean
// and returns the result. It returns false if the value is an empty string. If
// the value cannot be interpreted as a boolean, it panics.
func (e Env) GetBool() bool {
	if e.value == "" {
		return false
	}

	b, err := strconv.ParseBool(e.value)
	if err != nil {
		panic(err)
	}

	return b
}

// GetInt retrieves an integer from the environment variable associated with the
// [Env] instance. If the variable is unset or empty, it defaults to 0.
// Conversion errors result in a panic.
func (e Env) GetInt() int {
	if e.value == "" {
		return 0
	}

	i, err := strconv.Atoi(e.value)
	if err != nil {
		panic(err)
	}

	return i
}

// GetPathOrCreate retrieves the file path specified by the environment variable
// for this [Env] instance, creating the directory at that path if it does not
// exist. It returns an empty string if the environment variable is unset or
// empty. If any error occurs during directory creation, it triggers a panic.
func (e Env) GetPathOrCreate() string {
	if e.value == "" {
		return e.value
	}

	if !utils.IsFile(e.value) {
		if err := os.MkdirAll(e.value, 0755); err != nil {
			panic(errors.Wrap(
				errors.ErrInvalidEnv,
				errors.ErrInvalidPath,
				fmt.Errorf("env: %s", e.key),
				fmt.Errorf("value: %s", e.value),
				err,
			))
		}
	}

	return e.value
}

// GetPath retrieves the file path specified by the environment variable of the
// current [Env] instance. If the environment variable is not set or its value
// is an empty string, an empty string is returned. If the path does not exist
// or is not a valid file path, a panic is induced with relevant error
// information.
func (e Env) GetPath() string {
	if e.value == "" {
		return e.value
	}

	if !utils.IsFile(e.value) {
		panic(errors.Wrap(
			errors.ErrInvalidEnv,
			errors.ErrInvalidPath,
			fmt.Errorf("env: %s", e.key),
			fmt.Errorf("value: %s", e.value),
		))
	}

	return e.value
}

func (e Env) GetDuration() time.Duration {
	if e.value == "" {
		return 0
	}

	d, err := time.ParseDuration(e.value)
	if err != nil {
		panic(errors.Wrap(
			errors.ErrInvalidEnv,
			fmt.Errorf("env: %s", e.key),
			fmt.Errorf("value: %s", e.value),
			err,
		))
	}

	return d
}

// GetURL retrieves the URL from the associated environment variable. It ensures
// that the value is a valid URL format and returns it as a string. If the
// environment variable is not set or its value is empty, an empty string is
// returned. A panic occurs if the value cannot be parsed into a valid URL.
func (e Env) GetURL() string {
	if e.value == "" {
		return e.value
	}

	if _, err := url.Parse(e.value); err != nil {
		panic(errors.Wrap(
			errors.ErrInvalidEnv,
			fmt.Errorf("env: %s", e.key),
			fmt.Errorf("value: %s", e.value),
			err,
		))
	}

	return e.value
}

// GetURLPath extracts the path component from a URL present in the environment
// variable. It ensures the environment variable's value is a valid URL with an
// included path. If the environment variable is unset, empty, or contains an
// invalid URL, it triggers a panic.
func (e Env) GetURLPath() string {
	if e.value == "" {
		return e.value
	}

	if _, err := url.Parse(e.value); err != nil {
		panic(errors.Wrap(
			errors.ErrInvalidEnv,
			fmt.Errorf("env: %s", e.key),
			fmt.Errorf("value: %s", e.value),
			err,
		))
	}

	return e.value
}

// WithIntInRange ensures that the integer value of the environment variable is
// within a specified inclusive range. If the value is outside of this range, it
// panics to signal an invalid configuration. It returns the same [Env] instance
// to enable method chaining.
func (e Env) WithIntInRange(min, max int) Env {
	v := e.GetInt()

	if v < min || v > max {
		panic(errors.Wrap(
			errors.ErrInvalidEnv,
			fmt.Errorf("env: %s", e.key),
			fmt.Errorf("value: %d", v),
			fmt.Errorf("min: %d", min),
			fmt.Errorf("max %d", max),
		))
	}

	return e
}

// WithDefault sets a default value for the environment variable if it is not
// already set. It allows for specifying a fallback value to be used when an
// environment variable is absent or empty. The method returns the [Env]
// instance with the updated value, supporting further method chaining.
func (e Env) WithDefault(value string) Env {
	if e.value == "" {
		e.value = value
	}

	return e
}

// WithRequiredIf enforces a conditional requirement on the environment
// variable's value associated with the [Env] instance, based on the value of
// another specified environment variable. If the specified environment
// variable's value matches any string within a provided list, then the [Env]
// instance must have a non-empty value; otherwise, a panic is triggered to
// signal that a required environment variable is missing. This method supports
// method chaining by returning the same [Env] instance for further
// configuration.
func (e Env) WithRequiredIf(key string, values []string) Env {
	ne := NewEnv(key)

	if utils.StringInSlice(ne.GetString(), values) {
		e.checkRequired()
	}

	return e
}

// WithRequired ensures that the environment variable associated with the [Env]
// instance has a non-empty value, panicking if this is not the case. It returns
// the same [Env] instance to facilitate method chaining.
func (e Env) WithRequired() Env {
	e.checkRequired()

	return e
}

// WithOptions validates the environment variable's value against a set of
// predefined acceptable strings. If the value is not within these options and
// is not empty, it raises an error to signal an invalid value. This method
// facilitates fluent configuration by returning the same [Env] instance for
// potential additional configurations.
func (e Env) WithOptions(opts []string) Env {
	if e.value == "" {
		return e
	}

	if !utils.StringInSlice(e.value, opts) {
		panic(errors.Wrap(
			errors.ErrInvalidEnv,
			fmt.Errorf("env: %s", e.key),
			fmt.Errorf("value: %s", e.value),
			fmt.Errorf("options: %v", opts),
		))
	}

	return e
}
