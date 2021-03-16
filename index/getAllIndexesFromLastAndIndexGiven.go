package index

import "gitlab.com/evatix-go/core/constants"

func getAllIndexesFromLastAndIndexGiven(
	length int,
	rightStartsAtIndex,
	limit int,
) *[]int {
	newArrayLength := length - rightStartsAtIndex

	if newArrayLength <= constants.Zero {
		return nil
	}

	if limit > constants.MinusOne && newArrayLength > limit {
		newArrayLength = limit
	}

	finalIndexes := make(
		[]int,
		newArrayLength)

	if newArrayLength == constants.Zero {
		return &finalIndexes
	}

	newStartAt := newArrayLength - 1
	index := 0
	for i := newStartAt; i >= constants.Zero; i-- {
		finalIndexes[index] = i
		index++
	}

	return &finalIndexes
}
