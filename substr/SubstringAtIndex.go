package substr

import (
	"gitlab.com/evatix-go/strhelper/tostr"
)

// AtIndex
//
// language integrated ones will be faster str[startAtIndex:endsAtIndex]
// Under the hood this method usages that functionality from language
// panics if startsAtIndex < 0
func AtIndex(
	str string,
	startsAtIndex, endsAtIndex int,
) string {
	if startsAtIndex < 0 {
		message := "Substring Index cannot have negative startsAtIndex : " + tostr.FromInt(startsAtIndex)

		panic(message)
	}

	return str[startsAtIndex:endsAtIndex]
}
