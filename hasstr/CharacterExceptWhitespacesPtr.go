package hasstr

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/whitespace"
)

// CharacterExceptWhitespacesPtr Has at least one character other than space or whitespace
func CharacterExceptWhitespacesPtr(s *string) bool {
	return !(s == nil || *s == constants.EmptyString || whitespace.IsNullOrWhitespacePtr(s))
}
