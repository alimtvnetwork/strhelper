package strptr

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

func IsNull(s *string) bool {
	return s == nil || &s == nil
}

func IsNullOrEmpty(s *string) bool {
	return IsNull(s) || IsEmpty(s)
}

func IsNullOrWhitespace(s *string) bool {
	return IsNullOrEmpty(s) || len(strings.TrimSpace(*s)) == 0
}

func IsEmpty(s *string) bool {
	return len(*s) == 0 ||  *s == ""
}

func IsBlank(s *string) bool {
	return IsNullOrWhitespace(s)
}

func HasCharacter(s *string) bool {
	return !IsNullOrWhitespace(s)
}

func IsDefined(s *string) bool {
	return !IsNullOrWhitespace(s)
}

func IndexOf(s, findingString *string, isCaseSensitive bool) int {
	if isCaseSensitive {
		return strings.Index(*s, *findingString)
	}

	lowerCase := strings.ToLower(*s)
	findingLower := strings.ToLower(*findingString)

	return strings.Index(lowerCase, findingLower)
}

func IsExists(s, findingString *string, isCaseSensitive bool) bool {
	return IndexOf(s, findingString, isCaseSensitive) > constants.InvalidNotFoundCase
}

func DoesntExist(s, findingString *string, isCaseSensitive bool) bool {
	return !IsExists(s, findingString, isCaseSensitive)
}
