package index

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/panichelper"
	"gitlab.com/evatix-go/strhelper/strto"
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
func OfAllPtr(
	content *string,
	findingString *string,
	startsAtIndex int,
	limits int,
	isCaseSensitive bool,
) *[]int {
	if content == nil || findingString == nil {
		return nil
	}

	length := len(*content)
	searchLength := len(*findingString)

	if startsAtIndex <= constants.InvalidNotFoundCase || startsAtIndex > length-1 {
		panichelper.StartAtIndexFailed(startsAtIndex, length)
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
		sendingContent = strto.LowerStrPtr(sendingContent)
		sendingSearchTerm = strto.LowerStrPtr(findingString)
	}

	indexes := make(
		[]int,
		constants.Zero,
		length)

	lastIndex := length - 1
	foundIndex := ofUsingLengthPtr(
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

		foundIndex = ofUsingLengthPtr(
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

	return &indexes
}

// Returns all indexes where findingString is found.
//
// @limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all
//
// Results:
//  - Invalid result can be nil if any (content == nil || findingString == nil) results nil.
//  - If no indexes found returns nil.
func OfLastAllPtr(
	content *string,
	findingString *string,
	startsAtIndex int,
	limits int,
	isCaseSensitive bool,
) *[]int {
	if content == nil || findingString == nil {
		return nil
	}

	length := len(*content)
	searchLength := len(*findingString)

	if startsAtIndex <= constants.InvalidNotFoundCase || startsAtIndex > length-1 {
		panichelper.StartAtIndexFailed(startsAtIndex, length)
	}

	if length > 0 && searchLength == 0 {
		return getAllIndexesFromLastAndIndexGiven(
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
		sendingContent = strto.LowerStrPtr(sendingContent)
		sendingSearchTerm = strto.LowerStrPtr(findingString)
	}

	indexes := make(
		[]int,
		constants.Zero,
		length)

	lastIndex := length - 1
	foundIndex := ofUsingLengthPtr(
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

		foundIndex = ofUsingLengthPtr(
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

	return &indexes
}
