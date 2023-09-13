package chars

import (
	"gitlab.com/auk-go/core/constants"
)

func IsUpperCase(c uint8) bool {
	return c >= constants.UpperCaseA &&
		c <= constants.UpperCaseZ
}
