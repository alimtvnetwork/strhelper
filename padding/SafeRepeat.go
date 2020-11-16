package padding

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/strconst"
)

func SafeRepeat(padding *string, width int) string {
	if isstr.EmptyPtr(padding) {
		// nothing to repeat for
		return strconst.EmptyString
	}

	return strings.Repeat(*padding, width)
}
