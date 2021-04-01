package splits

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
)

func IntoTwoTrimSpace(s, separator *string) (left, right string) {
	splits := strings.SplitN(
		*s, *separator,
		constants.Two)

	length := len(splits)
	first := strings.TrimSpace(splits[coreindexes.First])

	if length == constants.Two {
		return first, strings.TrimSpace(splits[coreindexes.Second])
	}

	return first, constants.EmptyString
}
