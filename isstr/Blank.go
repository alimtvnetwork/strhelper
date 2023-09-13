package isstr

import (
	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/whitespace"
)

// Blank returns true if IsNullOrWhitespace(s)
func Blank(s string) bool {
	return s == constants.EmptyString || whitespace.IsWhitespaces(s)
}
