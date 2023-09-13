package isstr

import (
	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/whitespace"
)

// DefinedWithCharsExceptWhitespaces Has at least one character other than space or whitespace
func DefinedWithCharsExceptWhitespaces(s string) bool {
	return !(s == constants.EmptyString || whitespace.IsWhitespaces(s))
}
