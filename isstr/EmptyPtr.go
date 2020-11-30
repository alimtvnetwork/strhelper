package isstr

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// returns len(s) == 0 || s == constants.EmptyString
func EmptyPtr(s *string) bool {
	return s == nil || *s == strconst.EmptyString || len(*s) == 0
}
