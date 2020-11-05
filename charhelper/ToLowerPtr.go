package charhelper

import "gitlab.com/evatix-go/strhelper/constants"

func ToLowerPtr(c *uint8) uint8 {
	if *c >= constants.UpperCaseA &&
		*c <= constants.UpperCaseZ {
		return *c + constants.LowerCase
	}

	return *c
}
