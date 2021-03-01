package chars

import (
	"gitlab.com/evatix-go/core/constants"
)

func IsUpperCase(c uint8) bool {
	return c >= constants.UpperCaseA &&
		c <= constants.UpperCaseZ
}
