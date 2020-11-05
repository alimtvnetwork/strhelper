package charhelper

func IsMatchPtr(char1 *uint8, char2 *uint8, isCaseSensitive bool) bool {
	if isCaseSensitive {
		return *char1 == *char2
	}

	// Insensitive case
	return IsMatchCaseInsensitivePtr(char1, char2)
}
