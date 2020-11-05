package strs

import "gitlab.com/evatix-go/strhelper/constants"

// Returns all indexes where the string is found
// if empty lines given then returns nil
// Invalid result can be nil
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
		message := "startsAtIndex cannot be negative or more than length. startsAtIndex:" + string(startsAtIndex)

		panic(message)
	}

	indexes := make([]int, 0, length)
	foundIndex := IndexOf(
		lines,
		findingString,
		startsAtIndex,
		isCaseSensitive)

	for foundIndex > constants.InvalidNotFoundCase {
		indexes = append(indexes, foundIndex)
		foundIndex = IndexOf(
			lines,
			findingString,
			foundIndex+1,
			isCaseSensitive)
	}

	if len(indexes) == 0 {
		return nil
	}

	return indexes
}
