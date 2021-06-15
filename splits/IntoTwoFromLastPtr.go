package splits

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
)

func IntoTwoFromLastPtr(s, separator *string, isCaseSensitive bool) (left, right string) {
	splits := LastByLimitPtr(
		s,
		separator,
		isCaseSensitive,
		constants.Two)

	length := len(*splits)
	first := (*splits)[coreindexes.First]

	if length == constants.Two {
		return (*splits)[coreindexes.Second], first
	}

	return constants.EmptyString, first
}
