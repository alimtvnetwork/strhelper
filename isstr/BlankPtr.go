package isstr

import (
	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/whitespace"
)

// BlankPtr returns true if IsNullOrWhitespace(s)
func BlankPtr(s *string) bool {
	return s == nil || *s == constants.EmptyString || whitespace.IsWhitespaces(*s)
}
