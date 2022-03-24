package strs

import (
	"gitlab.com/evatix-go/core/defaultcapacity"
)

// OfAllFunc
//
//  returns all the indexes where the regexes are found.
func OfAllFunc(
	lines []string,
	f func(index int, in string) (isFound, isBreak bool),
) *[]int {
	if len(lines) == 0 {
		return &[]int{}
	}

	length := len(lines)
	capacity := defaultcapacity.OfSearch(length)
	list := make([]int, 0, capacity)

	for i, s := range lines {
		isFound, isBreak := f(i, s)
		if isFound {
			list = append(list, i)
		}

		if isBreak {
			break
		}
	}

	return &list
}
