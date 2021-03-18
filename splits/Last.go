package splits

import (
	"gitlab.com/evatix-go/core/constants"
)

func Last(s, separator string) *[]string {
	if s == "" {
		return Default()
	}

	return LastByLimitPtr(
		&s,
		&separator,
		true,
		constants.MinusOne)
}

func Default() *[]string {
	return &[]string{constants.EmptyString}
}

func DefaultWithStr(s *string) *[]string {
	if s == nil {
		return &[]string{constants.EmptyString}
	}

	return &[]string{*s}
}

func ByDot(s string) *[]string {
	if s == "" {
		return Default()
	}

	return LastByLimitPtr(
		&s,
		constants.DotPtr,
		true,
		constants.MinusOne)
}

func ByDotPtr(s *string) *[]string {
	if s == nil || *s == "" {
		return Default()
	}

	return LastByLimitPtr(
		s,
		constants.DotPtr,
		true,
		constants.MinusOne)
}
