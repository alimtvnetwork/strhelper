package splits

import "strings"

func BySpace(s string) []string {
	return strings.Fields(s)
}
