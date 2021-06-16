package splits

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
)

// LinesSimpleProcess split text using constants.NewLineUnix
// then returns lines processed by lineProcessor
func LinesSimpleProcess(
	s string,
	lineProcessor func(lineIn string) (lineOut string),
) []string {

	splitsLines := strings.Split(s, constants.NewLineUnix)
	length := len(splitsLines)
	slice := stringslice.Make(length, length)

	for i, lineIn := range splitsLines {
		lineOut := lineProcessor(lineIn)

		slice[i] = lineOut
	}

	return slice
}

