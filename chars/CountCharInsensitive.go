package chars

import "strings"

// CountCharInsensitive
//
// Counts and returns the count number based on chars present in the str (case: Insensitive)
// Returns 0 if str is nil or empty string.
func CountCharInsensitive(str *string, char1 uint8, startsAt int) int {
	length := len(*str)

	if length == 0 {
		return 0
	}

	strLower := strings.ToLower(*str)
	charLower := ToLower(char1)
	found := 0

	for ; startsAt < length; startsAt++ {
		char := (strLower)[startsAt]
		if char == charLower {
			found++
		}
	}

	return found
}
