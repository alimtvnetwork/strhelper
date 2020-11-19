package index

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/internal/pkg/constants"
	"gitlab.com/evatix-go/strhelper/internal/pkg/isinternal"
	"gitlab.com/evatix-go/strhelper/strconst"
)

// returns -1 on non found case
// panics if any is nil
func OfCaseInsensitive(s, findingString *string, startAt int) int {
	if s == nil || findingString == nil {
		panic(constants.SearchNullPanicMessage)
	}

	length := len(*s)
	wordLength := len(*findingString)

	if wordLength > length {
		return strconst.InvalidNotFoundCase
	}

	strLower := strings.ToLower(*s)
	wordLower := strings.ToLower(*findingString)

	for i := startAt; i < length; i++ {
		if length-i < wordLength {
			// there is no need to check anymore
			// exceeded word length and not found case
			break
		}

		if isinternal.IsStartsWithInternal(&strLower, &wordLower, i) {
			return i
		}
	}

	return strconst.InvalidNotFoundCase
}
