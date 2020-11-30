package chars

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

func GetSafeRunesIndexAtBy(runes *[]rune, index int) rune {
	if !(len(*runes)-1 >= index) {
		return strconst.InvalidNotFoundCase
	}

	return (*runes)[index]
}
