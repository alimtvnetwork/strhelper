package strs

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/isstrsinternal"
)

// GetNonEmptyStrings returns new array without empty strings, skip whitespaces if isTrimSpace true
func GetNonEmptyStrings(lines []string, isTrimSpace bool) []string {
	newLines := make([]string, 0, len(lines))

	if isstrsinternal.Empty(lines) {
		return newLines
	}

	for _, line := range lines {
		line2 := line

		if isTrimSpace {
			line2 = strings.TrimSpace(line2)
		}

		if line == constants.EmptyString || len(line) == 0 {
			continue
		}

		newLines = append(newLines, line2)
	}

	return newLines
}
