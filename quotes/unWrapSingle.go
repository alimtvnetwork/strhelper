package quotes

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Assumption here, s has single quotes and s it not empty
func unWrapSingle(s *string, isLeft bool) string {
	length := len(*s)

	if length == 1 {
		// has quote only
		return strconst.EmptyString
	}

	if isLeft {
		return (*s)[1 : length-1]
	}

	// right
	return (*s)[0 : length-2]
}
