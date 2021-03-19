package splits

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
)

func IntoTwoFromLastCaseSensitive(s, separator *string) (left, right string) {
	splits := LastByLimitPtr(
		s,
		separator,
		true,
		constants.Two)

	length := len(*splits)

	if length == 2 {
		return (*splits)[coreindexes.First], (*splits)[coreindexes.Second]
	}

	return (*splits)[coreindexes.First], constants.EmptyString
}
