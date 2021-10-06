package splits

import "gitlab.com/evatix-go/core/constants"

func LastByForwardSlash(
	s string,
	limits int,
) []string {
	return LastByRunePtr(
		s,
		constants.ForwardRune,
		limits)
}
