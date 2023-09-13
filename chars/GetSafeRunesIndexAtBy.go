package chars

import (
	"gitlab.com/auk-go/core/constants"
)

func GetSafeRunesIndexAtBy(runes *[]rune, index int) rune {
	if !(len(*runes)-1 >= index) {
		return constants.InvalidNotFoundCase
	}

	return (*runes)[index]
}
