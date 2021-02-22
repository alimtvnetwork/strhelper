package whitespace

import (
	"gitlab.com/evatix-go/strhelper/chars"
	"gitlab.com/evatix-go/strhelper/strconst"
)

// assumes input is not nil
// no explicit nil check for input is done
func OnlySpaceList(input *string) *[]int {
	length := len(*input)
	var onlySpaceIndex = make([]int, strconst.Zero, length)
	charsInput := []byte(*input)
	foundAny := false

	for i := 0; i < length; i++ {
		if chars.IsMatch(charsInput[i], strconst.SpaceChar, true) {
			onlySpaceIndex = append(onlySpaceIndex, i)
			foundAny = true
		}
	}

	if foundAny == false {
		return nil
	}

	return &onlySpaceIndex
}
