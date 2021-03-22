package splits

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
)

func IntoTwoFromLastUsingRune(s *string, splitRune rune) (left, right string) {
	splits := LastByRune(
		s,
		splitRune,
		constants.Two)

	length := len(*splits)
	first := (*splits)[coreindexes.First]

	if length == constants.Two {
		return (*splits)[coreindexes.Second], first
	}

	return constants.EmptyString, first
}
