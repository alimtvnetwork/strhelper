package splits

import "gitlab.com/auk-go/core/constants"

func LastByBackwardSlashPtr(
	s string,
	limits int,
) []string {
	return LastByRunePtr(
		s,
		constants.BackwardRune,
		limits)
}
