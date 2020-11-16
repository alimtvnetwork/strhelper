package isstr

import (
	"gitlab.com/evatix-go/strhelper/strconst"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

// returns true if IsNullOrWhitespace(s)
func BlankPtr(s *string) bool {
	return s == nil || *s == strconst.EmptyString || len(*s) == 0 || whitespace.IsWhitespaces(s)
}
