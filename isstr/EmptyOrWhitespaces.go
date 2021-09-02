package isstr

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/whitespace"
)

// EmptyOrWhitespaces returns true if IsNullOrWhitespace(s)
func EmptyOrWhitespaces(s string) bool {
	return s == constants.EmptyString || whitespace.IsWhitespaces(s)
}
