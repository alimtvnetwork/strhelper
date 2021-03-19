package splits

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
)

func IntoTwo(s, separator *string) (left, right string) {
	splits := strings.SplitN(
		*s, *separator,
		constants.Two)

	length := len(splits)

	if length == 2 {
		return splits[coreindexes.First], splits[coreindexes.Second]
	}

	return splits[coreindexes.First], constants.EmptyString
}

func IntoTwoFromLast(s, separator *string, isCaseSensitive bool) (left, right string) {
	splits := LastByLimitPtr(
		s,
		separator,
		isCaseSensitive,
		constants.Two)

	length := len(*splits)

	if length == 2 {
		return (*splits)[coreindexes.First], (*splits)[coreindexes.Second]
	}

	return (*splits)[coreindexes.First], constants.EmptyString
}
