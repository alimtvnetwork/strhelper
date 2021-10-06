package splits

import "gitlab.com/evatix-go/core/constants"

func LastByBackwardSlash(
	s string,
	limits int,
) []string {
	return LastByRunePtr(
		s,
		constants.BackwardRune,
		limits)
}
