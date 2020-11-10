package strs

import "gitlab.com/evatix-go/strhelper/constants"

// Returns all indexes where findingString is found.
//
// Results:
//  - Invalid result can be nil if any (content == nil || findingString == nil) results nil.
//  - If no indexes found returns nil.
func IndexesOfAll(
	lines *[]string,
	findingString *string,
	startsAtIndex int,
	isCaseSensitive bool,
) []int {
	if IsEmpty(lines) || findingString == nil {
		return nil
	}

	length := len(*lines)

	if startsAtIndex <= constants.InvalidNotFoundCase || startsAtIndex > length-1 {
		startAtIndexFailed(startsAtIndex)
	}

	indexes := make([]int, constants.Zero, length)
	foundIndex := IndexOf(
		lines,
		findingString,
		startsAtIndex,
		isCaseSensitive)

	if foundIndex > constants.InvalidNotFoundCase {
		indexes = append(indexes, foundIndex)
	}

	for foundIndex > constants.InvalidNotFoundCase {
		indexes = append(indexes, foundIndex)
		foundIndex = IndexOf(
			lines,
			findingString,
			foundIndex+1,
			isCaseSensitive)

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
