package splits

import (
	"gitlab.com/evatix-go/core/constants"
)

func LastByDot(s string) []string {
	if s == "" {
		return defaultResult()
	}

	return LastByLimitPtr(
		s,
		constants.Dot,
		true,
		constants.MinusOne)
}
