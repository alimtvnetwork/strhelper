package hasstr

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Has at least one character any, returns true even if a whitespace
func Character(s string) bool {
	return !(s == strconst.EmptyString || len(s) == 0)
}
