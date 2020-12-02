package index

import (
	"gitlab.com/evatix-go/strhelper/internal/panichelper"
	"gitlab.com/evatix-go/strhelper/strconst"
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

	if startsAtIndex <= strconst.InvalidNotFoundCase || startsAtIndex > length-1 {
		panichelper.StartAtIndexFailed(startsAtIndex, length)
	}

	if length > 0 && *findingString == "" {
		return getAllIndexesFromTheStartIndexGiven(length, startsAtIndex)
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

	indexes := make([]int, strconst.Zero, length)

	lastIndex := length - 1
	foundIndex := OfPtr(
		sendingContent,
		sendingSearchTerm,
		startsAtIndex,
		true)

	if foundIndex > strconst.InvalidNotFoundCase {
		indexes = append(indexes, foundIndex)
	}

	var nextIndex int

	for foundIndex > strconst.InvalidNotFoundCase {
		nextIndex = foundIndex + 1
		if nextIndex > lastIndex || (limits > -1 && len(indexes) >= limits) {
			break
		}

		foundIndex = OfPtr(
			sendingContent,
			sendingSearchTerm,
			nextIndex,
			true)

		if foundIndex > strconst.InvalidNotFoundCase {
			indexes = append(indexes, foundIndex)
		} else {
			// not found at any, will not continue
			break
		}
	}

	if len(indexes) == strconst.Zero {
		return nil
	}

	return &indexes
}
