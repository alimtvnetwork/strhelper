package chars

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// runes nil results -1 regardless
// Or else returns the index where the rune exist
func IndexOfRune(runes *[]rune, searchingFor rune) int {
	if runes == nil || len(*runes) == 0 {
		return strconst.InvalidNotFoundCase
	}

	for index, currentRune := range *runes {
		if currentRune == searchingFor {
			return index
		}
	}

	return strconst.InvalidNotFoundCase
}
