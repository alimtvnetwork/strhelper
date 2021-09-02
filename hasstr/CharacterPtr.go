package hasstr

import (
	"gitlab.com/evatix-go/core/constants"
)

// CharacterPtr Has at least one character any, returns true even if a whitespace
func CharacterPtr(s *string) bool {
	return !(s == nil || *s == constants.EmptyString)
}
