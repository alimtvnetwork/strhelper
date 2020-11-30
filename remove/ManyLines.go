package remove

import (
	"gitlab.com/evatix-go/strhelper/ds/strhashset"
	"gitlab.com/evatix-go/strhelper/internal/pkg/misc"
	"gitlab.com/evatix-go/strhelper/internal/pkg/panichelper"
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Creates new lines where removeLinesHashSet items will not appear.
//
// @count:
// - `-1` meaning remove all.
// - or else remove up to the given number only.
func ManyLines(
	lines *[]string,
	removeLinesHashSet *strhashset.Hashset,
	startsAt,
	count int,
	isCaseSensitive bool,
) *[]string {
	if removeLinesHashSet == nil || removeLinesHashSet.IsEmptySet() {
		return lines
	}

	length := len(*lines)
	lengthMinusOne := length - 1

	if startsAt < 0 || lengthMinusOne < startsAt {
		panichelper.StartAtIndexFailed(startsAt, length)
	}

	if startsAt == lengthMinusOne {
		// nothing to remove
		return lines
	}

	newLines := make([]string, startsAt, length)

	for i := 0; i < startsAt; i++ {
		newLines[i] = (*lines)[i]
	}

	if isCaseSensitive {
		return finalRemoveResultsCaseSensitive(
			lines,
			&newLines,
			removeLinesHashSet,
			startsAt,
			count,
			length,
		)
	}

	// insensitive
	lowerRemoveMap := removeLinesHashSet.ToLowerSet()
	// Reference : https://blog.golang.org/slices-intro | https://i.imgur.com/O3Hlmac.png
	// no copy just points
	linesRemainingParts := (*lines)[startsAt:]
	linesRemainingToLower := misc.ToLowerStrings(&linesRemainingParts)
	isCountUnset := count == strconst.InvalidNotFoundCase
	var line string
	index := 0
	for ; startsAt < length; startsAt++ {
		lowerLine := (*linesRemainingToLower)[index]

		if lowerRemoveMap.Has(lowerLine) && (isCountUnset || count > 0) {
			count--
			continue
		}

		line = (*lines)[startsAt]

		newLines = append(newLines, line)
	}

	return &newLines
}

func finalRemoveResultsCaseSensitive(
	lines *[]string,
	newLines *[]string,
	removeLinesHashSet *strhashset.Hashset,
	startsAt int,
	count int,
	length int,
) *[]string {
	var line string

	isCountUnset := count == strconst.InvalidNotFoundCase

	for i := startsAt; i < length; i++ {
		line = (*lines)[i]

		if removeLinesHashSet.Has(line) && (isCountUnset || count > 0) {
			count--
			continue
		}

		*newLines = append(*newLines, line)
	}

	return newLines
}
