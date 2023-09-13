package splits

import "gitlab.com/auk-go/core/constants"

var (
	bothSlashesMap = map[rune]bool{
		constants.ForwardRune:  true,
		constants.BackwardRune: true,
	}

	ExpectingLengthOfIntoTwoSplits = constants.Two
)
