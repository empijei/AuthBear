package tokens

import "time"

type Token interface {
	// Valid returns whether the token is valid, if it is not valid
	// the error will contain information on why
	Valid() (bool, error)

	// Expiry date
	Expiry() time.Time

	// Emission date
	Emission() time.Time

	// Get the token as a modifiable, decoded string
	Dump() string

	// Parses a modified, unencoded string into a token
	Parse(string) error

	// Returns the encoded token
	Token() string
}
