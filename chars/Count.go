package chars

import "unicode/utf8"

// Count
//
//	The best way to count length for
//	string is to convert to runes and then do len([]runes)
//	Or use utf8.RuneCountInString(input)
//
//	Ref: https://stackoverflow.com/a/55209017
func Count(input string) int {
	return utf8.RuneCountInString(input)
}
