package splits

import "gitlab.com/evatix-go/core/constants"

func splitByLastUsingEmptySep(
	s *string,
	limits int,
) *[]string {
	runes := []rune(*s)
	runesLength := len(runes)
	newLength := runesLength

	if newLength > limits && limits > -1 {
		newLength = limits
	}

	list := make([]string, newLength+1)

	index := 0
	runesIndex := runesLength - 1
	for ; runesIndex >= constants.Zero; runesIndex-- {
		list[index] = string(runes[runesIndex])
		index++
	}

	list[index+1] = string(runes[0:runesIndex])

	return &list
}
