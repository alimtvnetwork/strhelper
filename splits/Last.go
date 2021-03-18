package splits

import (
	"gitlab.com/evatix-go/core/constants"
)

func Last(s, separator string) *[]string {
	if s == "" {
		return defaultResult()
	}

	return LastByLimitPtr(
		&s,
		&separator,
		true,
		constants.MinusOne)
}
