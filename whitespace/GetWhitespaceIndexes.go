package whitespace

import (
	"unicode"

	"gitlab.com/evatix-go/strhelper/internal/pkg/panichelper"
)

// Returns the whitespace (including unicode whitespaces) indexes
//
// Returns nil if s is nil or empty
func GetWhitespaceIndexes(allRunes *[]rune, startsAt int) *[]int {
	if allRunes == nil || len(*allRunes) == 0 {
		return nil
	}

	length := len(*allRunes)

	if startsAt < 0 || length-1 < startsAt {
		panichelper.StartAtIndexFailed(startsAt, length)
	}

	indexes := make([]int, 0, length/2)
	hasFoundAny := false

	var rune rune
	for ; startsAt < length; startsAt++ {
		rune = (*allRunes)[startsAt]
		if (rune <= maxUnit8 && asciiSpaces[rune] == 1) || (rune > maxUnit8 && unicode.IsSpace(rune)) {
			indexes = append(indexes, startsAt)
			hasFoundAny = true
		}
	}

	if !hasFoundAny {
		return nil
	}

	return &indexes
}
