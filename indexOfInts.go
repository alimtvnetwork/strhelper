package strhelper

import "gitlab.com/evatix-go/strhelper/constants"

func indexOfInts(integers *[]int, findingInt int) int {
	for index, current := range *integers {
		if current == findingInt {
			return index
		}
	}

	return constants.InvalidNotFoundCase
}
