package isstr

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// returns len(s) == 0 || s == constants.EmptyString
// Better to use IsEmptyPtr
func Empty(s string) bool {
	return s == strconst.EmptyString || len(s) == 0
}
