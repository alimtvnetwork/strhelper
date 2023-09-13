package chars

// CountRunes
//
// Counts and returns the count number based on findingRunes present in the str
// Returns 0 if str is nil or empty string.
func CountRunes(str string, findingRunes []rune, startsAt int, isCaseSensitive bool) int {
	if len(str) == 0 {
		return 0
	}

	if isCaseSensitive {
		return CountRunesSensitive(str, findingRunes, startsAt)
	}

	return CountRunesInsensitive(str, findingRunes, startsAt)
}
