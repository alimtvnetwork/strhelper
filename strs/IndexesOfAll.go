package strs

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

// Returns all indexes where findingString is found.
//
// @limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all
//
// Results:
//  - Invalid result can be nil if any (content == nil || findingString == nil) results nil.
//  - If no indexes found returns nil.
func IndexesOfAll(
	lines *[]string,
	findingString *string,
	startsAtIndex int,
	limits int,
	isCaseSensitive bool,
) []int {
	if IsEmpty(lines) || findingString == nil {
		return nil
	}

	length := len(*lines)

	if startsAtIndex <= constants.InvalidNotFoundCase || startsAtIndex > length-1 {
		startAtIndexFailed(startsAtIndex, length)
	}

	// making a copy of pointer only, not the object. copy of reference address
	// reference : https://play.golang.org/p/r65MrCg86YH
	sendingLines := lines
	sendingSearchTerm := findingString

	if isCaseSensitive == false {
		// insensitive
		sendingLines = ToLowerStrings(lines)
		searchTermLowerCase := strings.ToLower(*findingString)
		sendingSearchTerm = &searchTermLowerCase
	}

	indexes := make([]int, constants.Zero, length)
	foundIndex := IndexOf(
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
		foundIndex = IndexOf(
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
