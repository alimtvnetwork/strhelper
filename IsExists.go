package strhelper

import "gitlab.com/evatix-go/strhelper/constants"

func IsExists(
	s, findingString string,
	isCaseSensitive bool,
) bool {
	return IndexOf(s, findingString, 0, isCaseSensitive) > constants.InvalidNotFoundCase
}
