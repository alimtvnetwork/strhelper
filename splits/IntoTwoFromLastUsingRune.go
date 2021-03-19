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

	if length == 2 {
		return (*splits)[coreindexes.First], (*splits)[coreindexes.Second]
	}

	return (*splits)[coreindexes.First], constants.EmptyString
}
