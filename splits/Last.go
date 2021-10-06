package splits

import (
	"gitlab.com/evatix-go/core/constants"
)

func Last(s, separator string) []string {
	if s == constants.EmptyString {
		return []string{""}
	}

	return LastByLimitPtr(
		s,
		separator,
		true,
		constants.TakeAllMinusOne)
}
