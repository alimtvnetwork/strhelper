package chars

import (
	"gitlab.com/evatix-go/core/constants"
)

func IsLowerCase(c uint8) bool {
	return c >= constants.LowerCaseA &&
		c <= constants.LowerCaseZ
}
