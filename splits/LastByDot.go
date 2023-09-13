package splits

import (
	"gitlab.com/auk-go/core/constants"
)

func LastByDot(s string) []string {
	if s == "" {
		return defaultResult()
	}

	return LastByLimit(
		s,
		constants.Dot,
		true,
		constants.MinusOne)
}
