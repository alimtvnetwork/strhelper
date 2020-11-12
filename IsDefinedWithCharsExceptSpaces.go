package strhelper

import (
	"gitlab.com/evatix-go/strhelper/constants"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

// Has at least one character other than space or whitespace
func IsDefinedWithCharsExceptSpaces(s string) bool {
	return !(s == constants.EmptyString || len(s) == 0 || whitespace.IsWhitespaces(&s))
}
