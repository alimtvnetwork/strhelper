package isstrsinternal

func caseSensitiveEqual(
	leftLines *[]string,
	rightLines *[]string,
	startsAt int,
) bool {
	leftLength := len(*leftLines)

	for ; startsAt < leftLength; startsAt++ {
		left := (*leftLines)[startsAt]
		right := (*rightLines)[startsAt]

		if left != right {
			return false
		}
	}

	return true
}
