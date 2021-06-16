package splits

import "gitlab.com/evatix-go/core/constants"

func LastByDotPtr(s *string) *[]string {
	if s == nil || *s == constants.EmptyString {
		return defaultResult()
	}

	return LastByLimitPtr(
		s,
		constants.DotPtr,
		true,
		constants.MinusOne)
}
