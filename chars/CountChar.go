package chars

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Counts and returns the count number based on char present in the str
// Returns 0 if str is nil or empty string.
func CountChar(str *string, char uint8, startAt int, isCaseSensitive bool) int {
	if str == nil || *str == strconst.EmptyString || len(*str) == 0 {
		return 0
	}

	if isCaseSensitive {
		return CountCharSensitive(str, char, startAt)
	}

	return CountCharInsensitive(str, char, startAt)
}
