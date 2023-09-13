package isstr

import "gitlab.com/auk-go/core/constants"

// Defined not empty string but something
func Defined(s string) bool {
	return s != constants.EmptyString
}
