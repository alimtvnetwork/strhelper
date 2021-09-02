package hasstr

import "gitlab.com/evatix-go/core/constants"

// CharPtr Has at least one character any, returns true even if a whitespace
func CharPtr(s *string) bool {
	return !(s == nil || *s == constants.EmptyString)
}
