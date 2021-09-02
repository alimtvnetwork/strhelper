package isstr

import (
	"gitlab.com/evatix-go/core/constants"
)

// EmptyPtr  s == constants.EmptyString
func EmptyPtr(s *string) bool {
	return s == nil || *s == constants.EmptyString
}
