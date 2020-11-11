package strhelper

import "gitlab.com/evatix-go/strhelper/constants"

// returns len(s) == 0 || s == constants.EmptyString
// Better to use IsEmptyPtr
func IsEmpty(s string) bool {
	return s == constants.EmptyString || len(s) == 0
}
