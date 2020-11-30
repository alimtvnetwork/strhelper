package chars

import "strings"

func CountRunesInsensitive(
	str *string,
	findingRunes *[]rune,
	at int,
) int {
	length := len(*str)
	found := 0

	if length == 0 {
		return found
	}

	runesOfString := []rune(strings.ToLower(*str))
	lowerFindingRunes := ToLowerRunes(findingRunes)
	// it needs to be updated because previous length was for byte/asc it changes.
	length = len(runesOfString)

	for ; at < length; at++ {
		currentRune := runesOfString[at]
		if IsRunesContains(lowerFindingRunes, currentRune) {
			found++
		}
	}

	return found
}
