package isstr

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/index"
)

func Exists(
	s, findingString string,
	isCaseSensitive bool,
) bool {
	return index.Of(
		s,
		findingString,
		constants.Zero,
		isCaseSensitive) > constants.InvalidNotFoundCase
}

func ExistsPtr(
	s, findingString *string,
	isCaseSensitive bool,
) bool {
	return index.OfPtr(
		s,
		findingString,
		constants.Zero,
		isCaseSensitive) > constants.InvalidNotFoundCase
}
