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
//  - `-1` means find all, 0 => nil
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
	if content == nil || findingString == nil || limits == 0 {
		return nil
	}

	if content == findingString && startsAtIndex == 0 {
		return &[]int{0}
	}

	wholeTextLength := len(*content)
	searchLength := len(*findingString)

	if searchLength > wholeTextLength-startsAtIndex {
		return nil
	}

	if startsAtIndex <= constants.InvalidNotFoundCase || startsAtIndex > wholeTextLength-1 {
		panichelper.StartAtIndexFailed(startsAtIndex, wholeTextLength)
	}

	if wholeTextLength > 0 && searchLength == 0 {
		return getAllIndexesFromLastAndIndexGiven(
			wholeTextLength,
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

	// keep the default as best so that doesn't resize.
	defaultCapacity := wholeTextLength

	if wholeTextLength > constants.ArbitraryCapacity1000 {
		defaultCapacity = constants.ArbitraryCapacity100
	}

	if wholeTextLength > constants.ArbitraryCapacity250 {
		defaultCapacity = wholeTextLength / constants.ArbitraryCapacity10
	}

	indexes := make(
		[]int,
		constants.Zero,
		defaultCapacity)

	lastIndex := wholeTextLength - 1
	foundIndex := ofLastCaseSensitiveUsingLength(
		sendingContent,
		sendingSearchTerm,
		startsAtIndex,
		wholeTextLength,
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

		foundIndex = ofLastCaseSensitiveUsingLength(
			sendingContent,
			sendingSearchTerm,
			nextIndex,
			wholeTextLength,
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
