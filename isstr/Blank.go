package isstr

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/whitespace"
)

// Blank returns true if IsNullOrWhitespace(s)
func Blank(s string) bool {
	return s == constants.EmptyString || whitespace.IsWhitespaces(s)
}
