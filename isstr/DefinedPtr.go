package isstr

import (
	"gitlab.com/evatix-go/core/constants"
)

// DefinedPtr not empty string but something
func DefinedPtr(s *string) bool {
	return !(s == nil || *s == constants.EmptyString)
}
