package stringindex

import "gitlab.com/evatix-go/core/constants"

func getAllIndexesFromTheStartIndexGiven(
	length int,
	startsAtIndex,
	limits int,
) []int {
	newArrayLength := length - startsAtIndex

	if newArrayLength <= 0 || limits == 0 {
		return nil
	}

	if limits > constants.MinusOne && newArrayLength > limits {
		newArrayLength = limits
	}

	finalIndexes := make(
		[]int,
		newArrayLength)

	if newArrayLength == 0 {
		return finalIndexes
	}

	index := 0
	for ; startsAtIndex < newArrayLength; startsAtIndex++ {
		finalIndexes[index] = startsAtIndex
		index++
	}

	return finalIndexes
}
