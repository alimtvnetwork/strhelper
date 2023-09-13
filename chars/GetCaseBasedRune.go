package chars

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

func GetCaseBasedRune(str string, r rune, isCaseSensitive bool) *CaseBasedRunes {
	if !isCaseSensitive {
		r = ToLowerRune(r)
	}

	if str == constants.EmptyString || len(str) == 0 {
		return &CaseBasedRunes{
			ToRunes:       nil,
			ComparingRune: r,
		}
	}

	toLowerRunes := []rune(strings.ToLower(str))

	return &CaseBasedRunes{
		ToRunes:       toLowerRunes,
		ComparingRune: r,
	}
}
