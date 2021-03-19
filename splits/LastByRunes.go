package splits

import (
	"gitlab.com/evatix-go/core"
)

func LastByRunes(
	s *string,
	limits int,
	runes ...rune,
) *[]string {
	if s == nil || *s == "" {
		return defaultResult()
	}

	if limits == 0 {
		return core.EmptyStringsPtr()
	}

	if runes == nil {
		return defaultResultWithStr(s)
	}

	length := len(runes)
	if length == 0 {
		return defaultResultWithStr(s)
	}

	runesMap := make(
		map[rune]bool, length)

	for _, r := range runes {
		runesMap[r] = true
	}

	return LastByRunesMap(
		s,
		&runesMap,
		limits)
}
