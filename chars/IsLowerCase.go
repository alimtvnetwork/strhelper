package chars

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

func IsLowerCase(c uint8) bool {
	return c >= strconst.LowerCaseA &&
		c <= strconst.LowerCaseZ
}
