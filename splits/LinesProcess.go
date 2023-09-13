package splits

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/stringslice"
)

// LinesProcess split text using constants.NewLineUnix
// then returns lines processed by lineProcessor
func LinesProcess(
	s string,
	lineProcessor func(index int, lineIn string) (lineOut string, isTake, isBreak bool),
) []string {
	splitsLines := strings.Split(s, constants.NewLineUnix)

	return stringslice.LinesProcess(splitsLines, lineProcessor)
}
