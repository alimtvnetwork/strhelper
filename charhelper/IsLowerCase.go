package charhelper

import "gitlab.com/evatix-go/strhelper/constants"

func IsLowerCase(c uint8) bool {
	return c >= constants.LowerCaseA &&
		c <= constants.LowerCaseZ
}
