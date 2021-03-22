package splits

import "gitlab.com/evatix-go/core/constants"

func LastByBackwardSlash(
	s *string,
	limits int,
) *[]string {
	return LastByRune(
		s,
		constants.BackwardRune,
		limits)
}
