package isstr

import (
	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/stringindex"
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
