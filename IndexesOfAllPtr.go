package strhelper

import "gitlab.com/evatix-go/strhelper/constants"

// Returns all indexes where findingString is found.
// if empty content given then returns nil
// Invalid result can be nil if any (content == nil || findingString == nil) then returns nil
func IndexesOfAllPtr(
	content *string,
	findingString *string,
	startsAtIndex int,
	isCaseSensitive bool,
) []int {
	if content == nil || findingString == nil {
		return nil
	}

	length := len(*content)

	if startsAtIndex <= constants.InvalidNotFoundCase || startsAtIndex > length-1 {
		startAtIndexFailed(startsAtIndex)
	}

	indexes := make([]int, constants.Zero, length)
	foundIndex := IndexOfPtr(
		content,
		findingString,
		startsAtIndex,
		isCaseSensitive)

	if foundIndex > constants.InvalidNotFoundCase {
		indexes = append(indexes, foundIndex)
	}

	for foundIndex > constants.InvalidNotFoundCase {
		foundIndex = IndexOfPtr(
			content,
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
