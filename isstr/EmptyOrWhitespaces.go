package isstr

import (
	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/whitespace"
)

// EmptyOrWhitespaces returns true if IsNullOrWhitespace(s)
func EmptyOrWhitespaces(s string) bool {
	return s == constants.EmptyString || whitespace.IsWhitespaces(s)
}
