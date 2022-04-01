package chars

import (
	"gitlab.com/evatix-go/core/constants"
)

// CountChar
//
// Counts and returns the count number based on char present in the str
// Returns 0 if str is nil or empty string.
func CountChar(str *string, char uint8, startsAt int, isCaseSensitive bool) int {
	if str == nil || *str == constants.EmptyString || len(*str) == 0 {
		return 0
	}

	if isCaseSensitive {
		return CountCharSensitive(str, char, startsAt)
	}

	return CountCharInsensitive(str, char, startsAt)
}
