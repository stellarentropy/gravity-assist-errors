package utils

import (
	"strings"

	guuid "github.com/google/uuid"
	"github.com/segmentio/ksuid"
	"github.com/stellarentropy/uuid"
)

// init prepares the random number generator used by the [guuid] package for
// UUID generation, ensuring a pool of random numbers is available for creating
// unique identifiers throughout the program's runtime.
func init() {
	guuid.EnableRandPool()
}

// uuidCustomGen provides a common generator for producing new UUID strings that
// are unique in both space and time, suitable for use in distributed systems
// where distinct identifiers are required.
var uuidCustomGen = uuid.NewGen()

// NewCustomUUID generates a new version 4 UUID as a string using a common
// custom generator to ensure the UUID is unique across different dimensions,
// such as spatial and temporal, which is essential for distributed systems that
// rely on identifiers being distinct and unrepeatable. It returns a string
// representation of the UUID.
func NewCustomUUID() string {
	return uuidCustomGen.NewV4().String()
}

// NewUUID generates a universally unique identifier and returns it as a string,
// ensuring high randomness and uniqueness suitable for various systems and
// applications.
func NewUUID() string {
	return NewGoogleUUID()
}

// NewPrefixedUUID combines a given prefix with a new UUID, separated by a
// hyphen, to produce a unique and identifiable string. The resulting string is
// useful where identifiers need to be both human-readable and globally unique.
func NewPrefixedUUID(prefix string) string {
	return strings.Join([]string{prefix, NewUUID()}, "-")
}

// NewGoogleUUID generates a new universally unique identifier (UUID) and
// returns it as a string. The UUID is designed to be sufficiently random to
// ensure uniqueness across different systems and over time, making it
// appropriate for use in various applications that require distinct string
// identifiers.
func NewGoogleUUID() string {
	return guuid.NewString()
}

// SetSafeKSUID configures the KSUID package to utilize a source of randomness
// that is cryptographically secure for generating unique identifiers. This
// configuration is recommended for use cases where security considerations
// outweigh the need for fast identifier generation.
func SetSafeKSUID() {
	ksuid.SetRand(nil)
}

// SetFastKSUID configures the KSUID generator to use a faster,
// non-cryptographically secure source of randomness. It is suitable for use
// cases where identifier generation speed is prioritized over the cryptographic
// strength of the identifiers.
func SetFastKSUID() {
	ksuid.SetRand(ksuid.FastRander)
}

// NewKSUID generates a globally unique identifier according to the KSUID
// specification and returns it as a string. The identifier incorporates a
// timestamp, making it suitable for applications that require time-ordered,
// unique IDs.
func NewKSUID() string {
	return ksuid.New().String()
}
