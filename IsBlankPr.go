package strhelper

import (
	"gitlab.com/evatix-go/strhelper/constants"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

// returns true if IsNullOrWhitespace(s)
func IsBlankPr(s *string) bool {
	return s == nil || *s == constants.EmptyString || len(*s) == 0 || whitespace.IsWhitespaceOnly(s)
}
