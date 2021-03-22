package splits

import "gitlab.com/evatix-go/core/constants"

func LastByForwardSlash(
	s *string,
	limits int,
) *[]string {
	return LastByRune(
		s,
		constants.ForwardRune,
		limits)
}
