package splits

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/stringslice"
)

// LinesSimpleProcess split text using constants.NewLineUnix
// then returns lines processed by lineProcessor
func LinesSimpleProcess(
	s string,
	lineProcessor func(lineIn string) (lineOut string),
) []string {
	splitsLines := strings.Split(s, constants.NewLineUnix)

	return stringslice.LinesSimpleProcess(splitsLines, lineProcessor)
}
