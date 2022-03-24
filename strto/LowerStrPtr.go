package strto

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/chars"
)

// Returns
//  - lower string as pointer of string
//  - invalid case: the same pointer back if nil
//
// Warning: It requires more memory to copy and then case string also the order is BigO(n)
func LowerStrPtr(s *string) *string {
	if s == nil {
		// return as is
		return s
	}

	runes := []rune(*s)
	toLowerString := string(*chars.ToLowerRunesInPlace(&runes))

	return &toLowerString
}

func Lower(s string) string {
	if s == "" {
		return ""
	}

	return strings.ToLower(s)
}
