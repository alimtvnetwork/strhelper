package isstr

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/stringindex"
)

func Exists(
	s, findingString string,
	isCaseSensitive bool,
) bool {
	return stringindex.Of(
		s,
		findingString,
		constants.Zero,
		isCaseSensitive) > constants.InvalidNotFoundCase
}
