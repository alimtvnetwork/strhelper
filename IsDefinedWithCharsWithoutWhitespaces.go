package strhelper

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

// Has at least one character other than space or whitespace
func IsDefinedWithCharsWithoutWhitespaces(s string) bool {
	return !(s == constants.EmptyString || len(s) == 0 || len(strings.TrimSpace(s)) == 0)
}
