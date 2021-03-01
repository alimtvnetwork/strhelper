package hasstr

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/whitespace"
)

// Has at least one character other than space or whitespace
func CharacterExceptWhitespaces(s string) bool {
	return !(s == constants.EmptyString || len(s) == 0 || whitespace.IsNullOrWhitespacePtr(&s))
}
