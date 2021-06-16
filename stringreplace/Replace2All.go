package stringreplace

import "strings"

func Replace2All(
	text string,
	search1, replace1 string,
	search2, replace2 string,
) string {
	text = strings.ReplaceAll(text, search1, replace1)
	text = strings.ReplaceAll(text, search2, replace2)

	return text
}
