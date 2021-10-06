package splits

import "gitlab.com/evatix-go/core/constants"

func LastByDotPtr(s string) []string {
	if s == constants.EmptyString {
		return defaultResult()
	}

	return LastByLimitPtr(
		s,
		constants.Dot,
		true,
		constants.MinusOne)
}
