package whitespace

import "gitlab.com/evatix-go/core/constants"

func IsWhitespacesPtr(s *string) bool {
	if s == nil {
		return true
	}

	sC := *s
	if sC == constants.EmptyString || sC == doubleSpace || sC == tripleSpace {
		return true
	}

	runes := []rune(sC)

	return IsRunesWhitespaces(runes)
}
