package hasstr

import (
	"gitlab.com/evatix-go/core/constants"
)

// Has at least one character any, returns true even if a whitespace
func Character(s string) bool {
	return !(s == constants.EmptyString || len(s) == 0)
}
