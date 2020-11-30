package hasstr

import (
	"gitlab.com/evatix-go/strhelper/strconst"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

// Has at least one character other than space or whitespace
func CharacterExceptWhitespaces(s string) bool {
	return !(s == strconst.EmptyString || len(s) == 0 || whitespace.IsNullOrWhitespacePtr(&s))
}
