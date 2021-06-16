package splits

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
)

// LinesNonWhiteSpacesPtr split text using constants.NewLineUnix
// then returns lines without any whitespaces
func LinesNonWhiteSpacesPtr(s string) *[]string {
	splits := strings.Split(s, constants.NewLineUnix)

	return stringslice.NonWhitespaceSlicePtr(&splits)
}
