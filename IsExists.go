package strhelper

import "gitlab.com/evatix-go/strhelper/constants"

func IsExists(
	s, findingString string,
	isCaseSensitive bool,
) bool {
	return IndexOf(
		s,
		findingString,
		constants.Zero,
		isCaseSensitive) > constants.InvalidNotFoundCase
}
