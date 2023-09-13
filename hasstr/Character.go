package hasstr

import (
	"gitlab.com/auk-go/core/constants"
)

// Character Has at least one character any, returns true even if a whitespace
func Character(s string) bool {
	return s != constants.EmptyString
}
