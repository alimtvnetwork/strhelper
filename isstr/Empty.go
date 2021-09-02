package isstr

import (
	"gitlab.com/evatix-go/core/constants"
)

func Empty(s string) bool {
	return s == constants.EmptyString
}
