package padding

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/isstr"
)

func SafeRepeat(padding *string, width int) string {
	if isstr.EmptyPtr(padding) {
		// nothing to repeat for
		return constants.EmptyString
	}

	if width == 1 {
		return *padding
	}

	return strings.Repeat(*padding, width)
}
