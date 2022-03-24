package hasstr

import "strings"

func WordOption(
	isCaseSensitive bool,
	input, search string,
) bool {
	if isCaseSensitive {
		return strings.Contains(
			input, search)
	}

	return strings.Contains(
		strings.ToLower(input),
		strings.ToLower(search))
}
