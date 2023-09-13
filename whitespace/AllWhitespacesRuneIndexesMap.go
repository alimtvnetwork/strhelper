package whitespace

import (
	"unicode"

	"gitlab.com/auk-go/core/constants"
)

func AllWhitespacesRuneIndexesMap(input *string) *map[rune]*[]int {
	if input == nil || len(*input) == constants.Zero {
		return nil
	}

	inputRunes := []rune(*input)
	length := len(inputRunes)

	var allWhitespaceList = make(map[rune]*[]int, length)
	foundAny := false
	var r rune

	for i := 0; i < length; i++ {
		r = (inputRunes)[i]
		if (r <= maxUnit8 && constants.AsciiSpace[r] == 1) ||
			(r > maxUnit8 && unicode.IsSpace(r)) {
			_, has := allWhitespaceList[r]

			if !has {
				// length/3 is a preliminary assumption for slice capacity
				list := make([]int, 0, length/3)
				allWhitespaceList[r] = &list
				foundAny = true
			}

			*allWhitespaceList[r] = append(*allWhitespaceList[r], i)
		}
	}

	if foundAny == false {
		return nil
	}

	return &allWhitespaceList
}
