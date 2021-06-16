package splits

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
)

// LinesSimpleProcessNoEmpty split text using constants.NewLineUnix
// then returns lines processed by lineProcessor and don't take any empty string in the output.
// Empty string lineOut will be discarded from the final outputs.
func LinesSimpleProcessNoEmpty(
	s string,
	lineProcessor func(lineIn string) (lineOut string),
) []string {

	splitsLines := strings.Split(s, constants.NewLineUnix)
	slice := stringslice.Make(constants.Zero, len(splitsLines))

	for _, lineIn := range splitsLines {
		lineOut := lineProcessor(lineIn)

		if lineOut == constants.EmptyString {
			continue
		}

		slice = append(slice, lineOut)
	}

	return slice
}
