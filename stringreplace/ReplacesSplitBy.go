package stringreplace

import (
	"strings"

	"gitlab.com/evatix-go/core/coredata/corestr"
)

func ReplacesSplitBy(
	text string,
	splitsBy string,
	replaceRequests ...corestr.KeyValuePair,
) []string {
	text = ReplacesAll(text, replaceRequests...)

	return strings.Split(text, splitsBy)
}
