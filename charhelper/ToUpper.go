package charhelper

import "gitlab.com/evatix-go/strhelper/constants"

func ToUpper(c uint8) uint8 {
	if c >= constants.LowerCaseA &&
		c <= constants.LowerCaseZ {
		return c + constants.UpperCaseA - constants.LowerCaseA
	}

	return c
}
