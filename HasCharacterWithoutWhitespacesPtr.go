package strhelper

import (
	"gitlab.com/evatix-go/strhelper/constants"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

// Has at least one character other than space or whitespace
func HasCharacterWithoutWhitespacesPtr(s *string) bool {
	return !(s == nil || *s == constants.EmptyString || len(*s) == 0 || whitespace.IsNullOrWhitespacePtr(s))
}
