package splits

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
)

// LinesProcess split text using constants.NewLineUnix
// then returns lines processed by lineProcessor
func LinesProcess(
	s string,
	lineProcessor func(index int, lineIn string) (lineOut string, isTake, isBreak bool),
) []string {
	splitsLines := strings.Split(s, constants.NewLineUnix)
	slice := stringslice.Make(constants.Zero, len(splitsLines))

	for i, lineIn := range splitsLines {
		lineOut, isTake, isBreak := lineProcessor(i, lineIn)

		if isTake {
			slice = append(slice, lineOut)
		}

		if isBreak {
			break
		}
	}

	return slice
}
