package strhelper

import (
	"gitlab.com/evatix-go/strhelper/constants"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

// returns true if IsNullOrWhitespace(s)
func IsBlank(s string) bool {
	return s == constants.EmptyString || len(s) == 0 || whitespace.IsWhitespaceOnly(&s)
}

// returns true if IsNullOrWhitespace(s)
func IsBlankPtr(s *string) bool {
	return s == nil || *s == constants.EmptyString || len(*s) == 0 || whitespace.IsWhitespaceOnly(s)
}
