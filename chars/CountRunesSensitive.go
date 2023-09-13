package chars

func CountRunesSensitive(
	str string,
	findingRunes []rune,
	at int,
) int {
	length := len(str)
	found := 0

	if length == 0 {
		return found
	}

	runesOfString := []rune(str)
	// it needs to be updated because previous length was for byte/asc it changes.
	length = len(runesOfString)

	for ; at < length; at++ {
		if IsRunesContains(findingRunes, runesOfString[at]) {
			found++
		}
	}

	return found
}
