package splits

func LastByRunes(
	s string,
	limits int,
	runes ...rune,
) []string {
	if s == "" {
		return defaultResult()
	}

	if limits == 0 {
		return []string{}
	}

	length := len(runes)
	if length == 0 {
		return []string{s}
	}

	runesMap := make(
		map[rune]bool, length)

	for _, r := range runes {
		runesMap[r] = true
	}

	return LastByRunesMap(
		s,
		runesMap,
		limits)
}
