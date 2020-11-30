package index

func getAllIndexesFromTheStartIndexGiven(
	length int,
	startsAtIndex int,
) *[]int {
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

	return &finalIndexes
}
