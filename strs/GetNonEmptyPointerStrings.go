package strs

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/internal/isstrsinternal"
	"gitlab.com/evatix-go/strhelper/strconst"
)

// GetNonEmptyPointerStrings returns new array without empty strings, skip whitespaces if isTrimSpace true
//
// Warning: For nil pointer / string "" will be skipped
func GetNonEmptyPointerStrings(lines *[]*string, isTrimSpace bool) *[]string {
	newLines := make([]string, 0, len(*lines))

	if isstrsinternal.EmptyPtrStr(lines) {
		return &newLines
	}

	for _, line := range *lines {
		if line == nil {
			continue
		}

		line2 := *line

		if isTrimSpace {
			line2 = strings.TrimSpace(line2)
		}

		if line2 == strconst.EmptyString || len(line2) == 0 {
			continue
		}

		newLines = append(newLines, line2)
	}

	return &newLines
}
