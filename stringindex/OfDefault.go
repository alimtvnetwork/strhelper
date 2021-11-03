package stringindex

import "strings"

func OfDefault(s, findingString string) int {
	return strings.Index(s, findingString)
}
