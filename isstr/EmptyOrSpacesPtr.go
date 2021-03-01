package isstr

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/whitespace"
)

// returns true if IsNullOrWhitespace(s)
func EmptyOrSpacesPtr(s *string) bool {
	return s == nil || *s == constants.EmptyString || whitespace.IsWhitespaces(s)
}
