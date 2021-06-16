package stringreplace

import "strings"

func Replace3All(
	text string,
	search1, replace1 string,
	search2, replace2 string,
	search3, replace3 string,
) string {
	text = strings.ReplaceAll(text, search1, replace1)
	text = strings.ReplaceAll(text, search2, replace2)
	text = strings.ReplaceAll(text, search3, replace3)

	return text
}
