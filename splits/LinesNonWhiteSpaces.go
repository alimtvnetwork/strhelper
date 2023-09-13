package splits

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/stringslice"
)

// LinesNonWhiteSpaces split text using constants.NewLineUnix
// then returns lines without any whitespaces
func LinesNonWhiteSpaces(s string) []string {
	splits := strings.Split(s, constants.NewLineUnix)

	return *stringslice.NonWhitespaceSlicePtr(&splits)
}
