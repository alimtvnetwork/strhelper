package splits

import (
	"gitlab.com/evatix-go/core/constants"
)

func LastByDot(s string) *[]string {
	if s == "" {
		return defaultResult()
	}

	return LastByLimitPtr(
		&s,
		constants.DotPtr,
		true,
		constants.MinusOne)
}

func LastByDotPtr(s *string) *[]string {
	if s == nil || *s == "" {
		return defaultResult()
	}

	return LastByLimitPtr(
		s,
		constants.DotPtr,
		true,
		constants.MinusOne)
}
