package isstr

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/whitespace"
)

// EmptyOrWhitespacesPtr returns true if IsNullOrWhitespace(s)
func EmptyOrWhitespacesPtr(s *string) bool {
	return s == nil || *s == constants.EmptyString || whitespace.IsWhitespaces(*s)
}
