package hasstr

import "strings"

func Word(input, search string) bool {
	return strings.Contains(input, search)
}
