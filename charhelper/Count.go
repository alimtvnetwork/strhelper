package charhelper

import (
	"gitlab.com/evatix-go/strhelper/constants"
)

// Counts and returns the count number based on char present in the str
// Returns 0 if str is nil or empty string.
func Count(str *string, char uint8, startAt int, isCaseSensitive bool) int {
	if str == nil || *str == constants.EmptyString || len(*str) == 0 {
		return 0
	}

	if isCaseSensitive {
		return CountSensitive(str, char, startAt)
	}

	return CountInsensitive(str, char, startAt)
}
