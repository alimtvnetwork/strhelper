package splits

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coreindexes"
)

func IntoTwoLeftRight(s, separator string) corestr.LeftRight {
	splits := strings.SplitN(
		s, separator,
		constants.Two)

	length := len(splits)
	first := splits[coreindexes.First]

	if length == constants.Two {
		return corestr.LeftRight{
			Left:    first,
			Right:   splits[coreindexes.Second],
			IsValid: true,
			Message: "",
		}
	}

	return corestr.LeftRight{
		Left:    first,
		Right:   constants.EmptyString,
		IsValid: false,
		Message: "",
	}
}
