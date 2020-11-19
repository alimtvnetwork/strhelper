package whitespace

import (
	"unicode"

	"gitlab.com/evatix-go/strhelper/internal/pkg/panichelper"
)

// Returns the whitespace (including unicode whitespaces) counts
func AllWhitespaceCountOfRunes(allRunes *[]rune, startsAt int) int {
	if allRunes == nil || len(*allRunes) == 0 {
		return 0
	}

	length := len(*allRunes)

	if startsAt < 0 || length-1 < startsAt {
		panichelper.StartAtIndexFailed(startsAt, length)
	}

	spacesFound := 0

	var rune rune
	for ; startsAt < length; startsAt++ {
		rune = (*allRunes)[startsAt]
		if (rune <= maxUnit8 && asciiSpaces[rune] == 1) || (rune > maxUnit8 && unicode.IsSpace(rune)) {
			spacesFound++
		}
	}

	return spacesFound
}
