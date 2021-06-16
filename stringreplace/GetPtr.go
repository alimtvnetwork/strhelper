package stringreplace

import (
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

func GetPtr(
	text, search,
	replaceWith *string,
	startsAt,
	howManyReplace int,
	isCaseSensitive bool,
) string {
	request := strhelpercore.ReplaceRequest{
		Text: text,
		ReplaceIndividualRequest: &strhelpercore.ReplaceIndividualRequest{
			Search:          *search,
			ReplaceWith:     *replaceWith,
			StartsAt:        startsAt,
			HowManyReplace:  howManyReplace,
			IsCaseSensitive: isCaseSensitive,
		},
	}

	return UsingReplaceRequest(&request)
}
