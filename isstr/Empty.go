package isstr

import (
	"gitlab.com/evatix-go/core/constants"
)

// returns len(s) == 0 || s == constants.EmptyString
// Better to use IsEmptyPtr
func Empty(s string) bool {
	return s == constants.EmptyString || len(s) == 0
}
