package substr

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/tostr"
)

// language integrated ones will be faster str[startAtIndex:endsAtIndex]
// Under the hood this method usages that functionality from language
// In case length is larger than given index, then it will fix to the length end
// panics if startsAtIndex < 0
// if endsAtIndex-startsAtIndex <= 0 then returns EmptyString(constants.EmptyString)
func SafeSubstringAtIndex(
	str string,
	startsAtIndex, endsAtIndex int,
) string {
	if startsAtIndex < 0 {
		message := "Substring Index cannot have negative startsAtIndex : " + tostr.FromInt(startsAtIndex)

		panic(message)
	}

	if len(str)-1 < endsAtIndex {
		endsAtIndex = len(str) - 1
	}

	if endsAtIndex-startsAtIndex <= 0 {
		return constants.EmptyString
	}

	return str[startsAtIndex:endsAtIndex]
}
