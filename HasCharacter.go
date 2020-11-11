package strhelper

import "gitlab.com/evatix-go/strhelper/constants"

// Has at least one character any, returns true even if a whitespace
func HasCharacter(s string) bool {
	return !(s == constants.EmptyString || len(s) == 0)
}
