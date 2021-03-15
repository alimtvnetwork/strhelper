package index

func getAllIndexesFromLastAndIndexGiven(
	length int,
	rightStartsAtIndex,
	limit int,
) *[]int {
	newArrayLength := length - rightStartsAtIndex

	if newArrayLength <= 0 {
		return nil
	}

	if limit > -1 && newArrayLength > limit {
		newArrayLength = limit
	}

	finalIndexes := make(
		[]int,
		newArrayLength)

	if newArrayLength == 0 {
		return &finalIndexes
	}

	newStartAt := length - 1 - rightStartsAtIndex
	index := 0
	for i := newStartAt; newArrayLength >= 0; newStartAt-- {
		finalIndexes[index] = i
		index++
	}

	return &finalIndexes
}
