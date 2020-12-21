package isstr

import (
	"gitlab.com/evatix-go/strhelper/strconst"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

// returns true if IsNullOrWhitespace(s)
func EmptyOrSpacesPtr(s *string) bool {
	return s == nil || *s == strconst.EmptyString || whitespace.IsWhitespaces(s)
}
