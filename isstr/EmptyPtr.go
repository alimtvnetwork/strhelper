package isstr

import (
	"gitlab.com/evatix-go/core/constants"
)

// returns len(s) == 0 || s == constants.EmptyString
func EmptyPtr(s *string) bool {
	return s == nil || *s == constants.EmptyString || len(*s) == 0
}
