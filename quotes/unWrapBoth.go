package quotes

import (
	"gitlab.com/evatix-go/core/constants"
)

// Assumption here, both quotations exist and s it not empty
func unWrapBoth(s *string) string {
	length := len(*s)

	if length == 2 {
		// both are quotes only
		return constants.EmptyString
	}

	return (*s)[1 : length-2]
}
