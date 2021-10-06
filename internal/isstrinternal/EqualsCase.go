package isstrinternal

import "strings"

func EqualsCase(isCaseSensitive bool, first, second string) bool {
	if first == second && isCaseSensitive {
		return true
	} else if isCaseSensitive {
		return false
	}

	return strings.EqualFold(first, second)
}
