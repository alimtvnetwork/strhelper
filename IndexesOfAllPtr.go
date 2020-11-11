package strhelper

import (
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
func IndexesOfAllPtr(
	content *string,
	findingString *string,
	startsAtIndex int,
	limits int,
	isCaseSensitive bool,
) []int {
	if content == nil || findingString == nil {
		return nil
	}

	length := len(*content)

	if startsAtIndex <= constants.InvalidNotFoundCase || startsAtIndex > length-1 {
		startAtIndexFailed(startsAtIndex, length)
	}

	if length > 0 && *findingString == "" {
		return getAllIndexesFromTheStartIndexGiven(length, startsAtIndex)
	}

	indexes := make([]int, constants.Zero, length)

	lastIndex := length - 1
	foundIndex := IndexOfPtr(
		content,
		findingString,
		startsAtIndex,
		isCaseSensitive)

	if foundIndex > constants.InvalidNotFoundCase {
		indexes = append(indexes, foundIndex)
	}

	var nextIndex int

	for foundIndex > constants.InvalidNotFoundCase {
		nextIndex = foundIndex + 1
		if nextIndex > lastIndex || (limits > -1 && len(indexes) >= limits) {
			break
		}

		foundIndex = IndexOfPtr(
			content,
			findingString,
			nextIndex,
			isCaseSensitive)

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

func getAllIndexesFromTheStartIndexGiven(
	length int,
	startsAtIndex int,
) []int {
	newArrayLength := length - startsAtIndex

	if newArrayLength <= 0 {
		return nil
	}

	finalIndexes := make([]int, newArrayLength)
	index := 0
	for ; startsAtIndex < newArrayLength; startsAtIndex++ {
		finalIndexes[index] = startsAtIndex
		index++
	}

	return finalIndexes
}
