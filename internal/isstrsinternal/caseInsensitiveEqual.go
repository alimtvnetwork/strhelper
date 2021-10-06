package isstrsinternal

import "strings"

func caseInsensitiveEqual(
	leftLines *[]string,
	rightLines *[]string,
	startsAt int,
) bool {
	leftLength := len(*leftLines)

	for ; startsAt < leftLength; startsAt++ {
		left := strings.ToLower((*leftLines)[startsAt])
		right := strings.ToLower((*rightLines)[startsAt])

		if left != right {
			return false
		}
	}

	return true
}
