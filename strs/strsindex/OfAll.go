package strsindex

import (
	"strings"

	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/internal/isstrsinternal"
	"gitlab.com/auk-go/strhelper/internal/panichelper"
	"gitlab.com/auk-go/strhelper/strs"
)

// Returns all indexes where findingString is found.
//
// @limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all
//
// Results:
//  - Invalid result can be nil if any lines == nil results nil.
//  - If no indexes found returns nil.
func OfAll(
	lines []string,
	findingString string,
	startsAtIndex int,
	limits int,
	isCaseSensitive bool,
) []int {
	if isstrsinternal.Empty(lines) {
		return nil
	}

	length := len(lines)

	if startsAtIndex <= constants.InvalidNotFoundCase || startsAtIndex > length-1 {
		panichelper.StartAtIndexFailed(startsAtIndex, length)
	}

	// making a copy of pointer only, not the object. copy of reference address
	// reference : https://play.golang.org/p/r65MrCg86YH
	sendingLines := lines
	sendingSearchTerm := findingString

	if isCaseSensitive == false {
		// insensitive
		sendingLines = strs.ToLowerStrings(lines)
		sendingSearchTerm = strings.ToLower(findingString)
	}

	indexes := make([]int, constants.Zero, length)
	foundIndex := Of(
		sendingLines,
		sendingSearchTerm,
		startsAtIndex,
		true)

	if foundIndex > constants.InvalidNotFoundCase {
		indexes = append(indexes, foundIndex)
	}

	var nextIndex int
	lastIndex := length - 1

	for foundIndex > constants.InvalidNotFoundCase {
		nextIndex = foundIndex + 1
		if nextIndex > lastIndex || (limits > -1 && len(indexes) >= limits) {
			break
		}

		indexes = append(indexes, foundIndex)
		foundIndex = Of(
			sendingLines,
			sendingSearchTerm,
			nextIndex,
			true)

		if foundIndex > constants.InvalidNotFoundCase {
			indexes = append(indexes, foundIndex)
		} else {
			break
		}
	}

	if len(indexes) == constants.Zero {
		return nil
	}

	return indexes
}
