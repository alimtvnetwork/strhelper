package panichelper

import (
	"strings"

	"gitlab.com/evatix-go/strhelper"
	"gitlab.com/evatix-go/strhelper/constants"
)

func SafeRepeat(padding *string, width int) string {
	if strhelper.IsEmptyPtr(padding) {
		// nothing to repeat for
		return constants.EmptyString
	}

	return strings.Repeat(*padding, width)
}
