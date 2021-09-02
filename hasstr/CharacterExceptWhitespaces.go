package hasstr

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/whitespace"
)

// CharacterExceptWhitespaces Has at least one character other than space or whitespace
func CharacterExceptWhitespaces(s string) bool {
	return !(s == constants.EmptyString || whitespace.IsNullOrWhitespacePtr(&s))
}
