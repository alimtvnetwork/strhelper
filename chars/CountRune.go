package chars

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Counts and returns the count number based on char1 present in the rune
// Returns 0 if str is nil or empty string.
func CountRune(str *string, rune rune, startAt int, isCaseSensitive bool) int {
	if str == nil || *str == strconst.EmptyString || len(*str) == 0 {
		return 0
	}

	if isCaseSensitive {
		return CountRuneSensitive(str, rune, startAt)
	}

	return CountRuneInsensitive(str, rune, startAt)
}
