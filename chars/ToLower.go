package chars

import (
	"gitlab.com/evatix-go/core/constants"
)

func ToLower(c uint8) uint8 {
	if c >= constants.UpperCaseA &&
		c <= constants.UpperCaseZ {
		return c + constants.LowerCase
	}

	return c
}
