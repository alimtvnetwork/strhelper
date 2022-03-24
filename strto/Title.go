package strto

import "strings"

func Title(s string) string {
	if s == "" {
		return ""
	}

	return strings.ToTitle(s)
}
