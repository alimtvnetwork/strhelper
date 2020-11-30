package whitespace

import (
	"unicode"

	"gitlab.com/evatix-go/strhelper/internal/pkg/panichelper"
)

// AllWhitespacesCountOfRunes returns the whitespace (including unicode whitespaces) counts
func AllWhitespacesCountOfRunes(allRunes *[]rune, startsAt int) int {
	if allRunes == nil || len(*allRunes) == 0 {
		return 0
	}

	length := len(*allRunes)

	if startsAt < 0 || length-1 < startsAt {
		panichelper.StartAtIndexFailed(startsAt, length)
	}

	spacesFound := 0

	var r rune
	for ; startsAt < length; startsAt++ {
		r = (*allRunes)[startsAt]
		if (r <= maxUnit8 && asciiSpaces[r] == 1) || (r > maxUnit8 && unicode.IsSpace(r)) {
			spacesFound++
		}
	}

	return spacesFound
}
