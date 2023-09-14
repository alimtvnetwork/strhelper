package strto

import (
	"strings"

	"gitlab.com/auk-go/strhelper/chars"
)

// Returns
//   - lower string as pointer of string
//   - invalid case: the same pointer back if nil
//
// Warning: It requires more memory to copy and then case string also the order is BigO(n)
func LowerString(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)
	toLowerString := string(chars.ToLowerRunesInPlace(runes))

	return toLowerString
}

// Returns
//   - lower string as pointer of string
//   - invalid case: the same pointer back if nil
//
// Warning: It requires more memory to copy and then case string also the order is BigO(n)
func LowerStringPtr(s *string) *string {
	if s == nil {
		return s
	}

	toLower := LowerString(*s)

	return &toLower
}

func Lower(s string) string {
	if s == "" {
		return ""
	}

	return strings.ToLower(s)
}
