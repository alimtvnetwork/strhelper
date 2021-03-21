package padding

import "gitlab.com/evatix-go/core/constants"

func RepeatPtr(s *string, width int) string {
	if s == nil {
		return constants.EmptyString
	}

	length := len(*s)

	if length == 0 {
		return constants.EmptyString
	}

	allChars := make([]byte, length*width)
	// keeping same as index
	width--
	for width >= 0 {
		for i := 0; i < length; i++ {
			allChars[width+i] = (*s)[i]
		}

		width--
	}

	return string(allChars)
}
