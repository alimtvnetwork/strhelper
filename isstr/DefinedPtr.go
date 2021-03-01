package isstr

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/whitespace"
)

// Has at least one character other than space or whitespace
func DefinedPtr(s *string) bool {
	return !(s == nil || *s == constants.EmptyString || len(*s) == 0 || whitespace.IsNullOrWhitespacePtr(s))
}
