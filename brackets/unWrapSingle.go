package brackets

import (
	"gitlab.com/evatix-go/core/constants"
)

// Assumption here, s has single bracket and s it is not empty
func unWrapSingle(s *string, isLeft bool) string {
	length := len(*s)

	if length == 1 {
		// it has brackets only
		return constants.EmptyString
	}

	if isLeft {
		return (*s)[1 : length-1]
	}

	// right
	return (*s)[0 : length-2]
}
