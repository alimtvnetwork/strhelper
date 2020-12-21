package isstr

import (
	"gitlab.com/evatix-go/strhelper/strconst"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

// returns true if IsNullOrWhitespace(s)
func EmptyOrSpaces(s string) bool {
	return s == strconst.EmptyString || whitespace.IsWhitespaces(&s)
}
