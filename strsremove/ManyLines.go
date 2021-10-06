package strsremove

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/ds/strhashset"
	"gitlab.com/evatix-go/strhelper/internal/misc"
	"gitlab.com/evatix-go/strhelper/internal/panichelper"
)

// ManyLines
//
// Creates new lines where removeLinesHashSet items will not appear.
//
// @count:
// - `-1` meaning remove all.
// - or else remove up to the given number only.
func ManyLines(
	lines []string,
	removeLinesHashset *strhashset.Hashset,
	startsAt,
	count int,
	isCaseSensitive bool,
) []string {
	if removeLinesHashset == nil || removeLinesHashset.IsEmptySet() {
		return lines
	}

	length := len(lines)
	lengthMinusOne := length - 1

	if startsAt < 0 || lengthMinusOne < startsAt {
		panichelper.StartAtIndexFailed(startsAt, length)
	}

	if startsAt == lengthMinusOne {
		// nothing to strsremove
		return lines
	}

	newLines := make([]string, startsAt, length)

	for i := 0; i < startsAt; i++ {
		newLines[i] = lines[i]
	}

	if isCaseSensitive {
		return finalRemoveResultsCaseSensitive(
			lines,
			newLines,
			removeLinesHashset,
			startsAt,
			count,
			length,
		)
	}

	// insensitive
	lowerRemoveMap := removeLinesHashset.ToLowerSet()
	// Reference : https://blog.golang.org/slices-intro | https://i.imgur.com/O3Hlmac.png
	// no copy just points
	linesRemainingParts := lines[startsAt:]
	linesRemainingToLower := misc.ToLowerStrings(&linesRemainingParts)
	isCountUnset := count == constants.InvalidNotFoundCase
	var line string
	index := 0
	for ; startsAt < length; startsAt++ {
		lowerLine := (*linesRemainingToLower)[index]

		if lowerRemoveMap.Has(lowerLine) && (isCountUnset || count > 0) {
			count--
			continue
		}

		line = lines[startsAt]

		newLines = append(newLines, line)
	}

	return newLines
}
