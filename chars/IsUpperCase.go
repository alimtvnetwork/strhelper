package chars

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

func IsUpperCase(c uint8) bool {
	return c >= strconst.UpperCaseA &&
		c <= strconst.UpperCaseZ
}
