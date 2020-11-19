package whitespace

import "gitlab.com/evatix-go/strhelper/internal/pkg/panichelper"

func AllNewLinesCountOfRunes(allRunes *[]rune, startsAt int) int {
	if allRunes == nil || len(*allRunes) == 0 {
		return 0
	}

	length := len(*allRunes)

	if startsAt < 0 || length-1 < startsAt {
		panichelper.StartAtIndexFailed(startsAt, length)
	}

	newLineFound := 0
	var rune rune

	for ; startsAt < length; startsAt++ {
		rune = (*allRunes)[startsAt]
		if rune <= maxUnit8 && asciiNewLinesChars[rune] == 1 {
			newLineFound++
		}
	}

	return newLineFound
}
