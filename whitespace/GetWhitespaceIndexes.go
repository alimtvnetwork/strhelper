package whitespace

import (
	"unicode"

	"gitlab.com/evatix-go/strhelper/internal/panichelper"
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

	var r rune
	for ; startsAt < length; startsAt++ {
		r = (*allRunes)[startsAt]
		if (r <= maxUnit8 && asciiSpaces[r] == 1) || (r > maxUnit8 && unicode.IsSpace(r)) {
			indexes = append(indexes, startsAt)
			hasFoundAny = true
		}
	}

	if !hasFoundAny {
		return nil
	}

	return &indexes
}
