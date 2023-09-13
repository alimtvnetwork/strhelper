package chars

import (
	"gitlab.com/auk-go/core/constants"
)

// runes nil results -1 regardless
// Or else returns the index where the rune exist
func IndexOfRune(runes *[]rune, searchingFor rune) int {
	if runes == nil || len(*runes) == 0 {
		return constants.InvalidNotFoundCase
	}

	for index, currentRune := range *runes {
		if currentRune == searchingFor {
			return index
		}
	}

	return constants.InvalidNotFoundCase
}
