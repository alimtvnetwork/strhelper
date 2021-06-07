package strsindex

import "gitlab.com/evatix-go/core/constants"

// OfCaseSensitive Returns the index where the string first found, rest don't care
func OfCaseSensitive(
	lines *[]string,
	findingString string,
) int {
	return Of(
		lines,
		findingString,
		constants.Zero,
		true)
}
