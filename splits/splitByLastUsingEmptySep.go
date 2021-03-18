package splits

func splitByLastUsingEmptySep(
	s *string,
	limits int,
) *[]string {
	runes := []rune(*s)
	runesLength := len(runes)
	newLength := runesLength
	var list []string = nil

	if newLength > limits && limits > -1 {
		newLength = limits
		list = make([]string, newLength+1)
	} else {
		list = make([]string, runesLength)
	}

	index := 0
	runesIndex := runesLength - 1
	for ; index < newLength; index++ {
		list[index] = string(runes[runesIndex])
		runesIndex--
	}

	if runesIndex > -1 {
		list[index] = string(runes[0 : runesIndex+1])
	}

	return &list
}
