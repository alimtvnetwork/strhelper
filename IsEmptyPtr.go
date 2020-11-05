package strhelper

import "gitlab.com/evatix-go/strhelper/constants"

// returns len(s) == 0 || s == constants.EmptyString
func IsEmptyPtr(s *string) bool {
	return s == nil || *s == constants.EmptyString || len(*s) == 0
}
