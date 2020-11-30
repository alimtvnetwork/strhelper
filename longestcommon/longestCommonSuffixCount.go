package longestcommon

func longestCommonSuffixCount(a *string, b *string, bothLastIndexReduceBy int) int {
	lenA := len(*a)
	lenB := len(*b)
	incrementing := 0

	for ; bothLastIndexReduceBy < lenA && bothLastIndexReduceBy < lenB; bothLastIndexReduceBy++ {
		if (*a)[lenA-1-bothLastIndexReduceBy] != (*b)[lenB-1-bothLastIndexReduceBy] {
			return incrementing
		}

		incrementing++
	}

	return incrementing
}
