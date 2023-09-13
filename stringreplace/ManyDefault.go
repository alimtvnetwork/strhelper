package stringreplace

import "gitlab.com/auk-go/core/constants"

func ManyDefault(
	text string,
	searchReplaceMap map[string]string,
) string {
	return ManyPtr(
		text,
		searchReplaceMap,
		constants.Zero,
		constants.TakeAllMinusOne,
		true)
}
