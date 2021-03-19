package splits

import (
	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
)


func IntoTwoFromLastCaseSensitive(s, separator *string) (left, right string) {
	splits := LastByLimitPtr(
		s,
		separator,
		true,
		constants.One)

	length := len(*splits)

	if length == 2 {
		return (*splits)[coreindexes.First], (*splits)[coreindexes.Second]
	}

	return (*splits)[coreindexes.First], constants.EmptyString
}


func IntoTwoFromLast(s, separator *string, isCaseSensitive bool) (left, right string) {
	splits := LastByLimitPtr(
		s,
		separator,
		isCaseSensitive,
		constants.One)

	length := len(*splits)

	if length == 2 {
		return (*splits)[coreindexes.First], (*splits)[coreindexes.Second]
	}

	return (*splits)[coreindexes.First], constants.EmptyString
}

func IntoTwoFromLastUsingRune(s *string, splitRune rune) (left, right string) {
	splits := LastByRune(
		s,
		splitRune,
		constants.One)

	length := len(*splits)

	if length == 2 {
		return (*splits)[coreindexes.First], (*splits)[coreindexes.Second]
	}

	return (*splits)[coreindexes.First], constants.EmptyString
}

func LastByForwardSlash(
	s *string,
	limits int,
) *[]string {
	return LastByRune(
		s,
		constants.ForwardRune,
		limits)
}

func LastByBackwardSlash(
	s *string,
	limits int,
) *[]string {
	return LastByRune(
		s,
		constants.BackwardRune,
		limits)
}

func LastByBothSlashes(
	s *string,
	limits int,
) *[]string {
	return LastByRunesMap(
		s,
		bothSlashesMap,
		limits)
}

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
