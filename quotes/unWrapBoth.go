package quotes

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Assumption here, both quotations exist and s it not empty
func unWrapBoth(s *string) string {
	length := len(*s)

	if length == 2 {
		// both are quotes only
		return strconst.EmptyString
	}

	return (*s)[1 : length-2]
}
