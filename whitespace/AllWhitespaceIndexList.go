package whitespace

import (
	"gitlab.com/evatix-go/strhelper/strconst"
	"unicode"
)

func AllWhitespaceIndexList(input *string) *map[rune]*[]int {
	if input == nil || len(*input) == strconst.Zero {
		return nil
	}

	inputRunes := []rune(*input)
	length := len(inputRunes)

	var allWhitespaceList = make(map[rune]*[]int, length)
	foundAny := false
	var r rune

	for i := 0; i < length; i++ {
		r = (inputRunes)[i]
		if (r <= maxUnit8 && strconst.AsciiSpace[r] == 1) ||
			(r > maxUnit8 && unicode.IsSpace(r)) {
			listPtr, has := allWhitespaceList[r]

			if !has {
				// length/3 is a preliminary assumption for slice capacity
				list := make([]int, 0, length/3)
				list = append(list, i)
				allWhitespaceList[r] = &list
				foundAny = true
			}

			*listPtr = append(*listPtr, i)
		}
	}

	if foundAny == false {
		return nil
	}

	return &allWhitespaceList
}
