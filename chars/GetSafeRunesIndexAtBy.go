package chars

import (
	"gitlab.com/evatix-go/core/constants"
)

func GetSafeRunesIndexAtBy(runes *[]rune, index int) rune {
	if !(len(*runes)-1 >= index) {
		return constants.InvalidNotFoundCase
	}

	return (*runes)[index]
}
