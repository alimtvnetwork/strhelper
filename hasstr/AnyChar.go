package hasstr

import "gitlab.com/auk-go/core/constants"

// AnyChar Has at least one character any, returns true even if a whitespace
func AnyChar(s string) bool {
	return s != constants.EmptyString
}
