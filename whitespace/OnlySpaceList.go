package whitespace

import (
	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/chars"
)

// assumes input is not nil
// no explicit nil check for input is done
func OnlySpaceList(input *string) *[]int {
	length := len(*input)
	var onlySpaceIndex = make([]int, constants.Zero, length)
	charsInput := []byte(*input)
	foundAny := false

	for i := 0; i < length; i++ {
		if chars.IsMatch(charsInput[i], constants.SpaceChar, true) {
			onlySpaceIndex = append(onlySpaceIndex, i)
			foundAny = true
		}
	}

	if foundAny == false {
		return nil
	}

	return &onlySpaceIndex
}
