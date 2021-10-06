package index

import "gitlab.com/evatix-go/core/constants"

// OfCaseSensitiveSimple returns -1 on non found case
func OfCaseSensitiveSimple(s, findingString string) int {
	return OfCaseSensitive(
		s,
		findingString,
		constants.Zero)
}
