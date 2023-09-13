package hasstr

import (
	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/whitespace"
)

// CharacterExceptWhitespaces Has at least one character other than space or whitespace
func CharacterExceptWhitespaces(s string) bool {
	return !(s == constants.EmptyString || whitespace.IsNullOrWhitespacePtr(&s))
}
