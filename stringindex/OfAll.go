package stringindex

import (
	"strings"

	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/internal/panichelper"
)

// Returns all indexes where findingString is found.
//
// @limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all, 0 => nil
//
// Results:
//  - Invalid result can be nil if any (content == nil || findingString == nil) results nil.
//  - If no indexes found returns nil.
func OfAll(
	content string,
	findingString string,
	startsAtIndex int,
	limits int,
	isCaseSensitive bool,
) []int {
	if limits == 0 {
		return nil
	}

	if content == findingString && startsAtIndex == 0 {
		return []int{constants.Zero}
	}

	length := len(content)
	searchLength := len(findingString)

	if startsAtIndex <= constants.InvalidNotFoundCase || startsAtIndex > length-1 {
		panichelper.StartAtIndexFailed(startsAtIndex, length)
	}

	if searchLength > length-startsAtIndex {
		return nil
	}

	if length > 0 && searchLength == 0 {
		return getAllIndexesFromTheStartIndexGiven(
			length,
			startsAtIndex,
			limits)
	}

	if searchLength == 0 {
		return nil
	}

	// making a copy of pointer only, not the object. copy of reference address
	// reference : https://play.golang.org/p/r65MrCg86YH
	sendingContent := content
	sendingSearchTerm := findingString

	if isCaseSensitive == false {
		// insensitive
		sendingContent = strings.ToLower(sendingContent)
		sendingSearchTerm = strings.ToLower(findingString)
	}

	indexes := make(
		[]int,
		constants.Zero,
		length)

	lastIndex := length - 1
	foundIndex := ofCaseSensitiveUsingLengthPtr(
		sendingContent,
		sendingSearchTerm,
		startsAtIndex,
		length,
		searchLength)

	if foundIndex > constants.InvalidNotFoundCase {
		indexes = append(indexes, foundIndex)
	}

	var nextIndex int

	for foundIndex > constants.InvalidNotFoundCase {
		nextIndex = foundIndex + 1
		if nextIndex > lastIndex || (limits > -1 && len(indexes) >= limits) {
			break
		}

		foundIndex = ofCaseSensitiveUsingLengthPtr(
			sendingContent,
			sendingSearchTerm,
			nextIndex,
			length,
			searchLength)

		if foundIndex > constants.InvalidNotFoundCase {
			indexes = append(indexes, foundIndex)
		} else {
			// not found at any, will not continue
			break
		}
	}

	if len(indexes) == constants.Zero {
		return nil
	}

	return indexes
}
