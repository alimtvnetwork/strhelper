package strhelper

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

func SafeRepeat(padding *string, width int) string {
	if IsEmptyPtr(padding) {
		// nothing to repeat for
		return constants.EmptyString
	}

	return strings.Repeat(*padding, width)
}
