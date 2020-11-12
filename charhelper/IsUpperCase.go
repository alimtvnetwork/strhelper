package charhelper

import "gitlab.com/evatix-go/strhelper/constants"

func IsUpperCase(c uint8) bool {
	return c >= constants.UpperCaseA &&
		c <= constants.UpperCaseZ
}
