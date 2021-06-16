package stringreplace

import (
	"strings"
)

func ReplaceSplitBy(
	text string,
	search, replace string,
	splitsBy string,
) []string {
	text = strings.ReplaceAll(text, search, replace)

	return strings.Split(text, splitsBy)
}
