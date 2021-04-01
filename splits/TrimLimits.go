package splits

import "strings"

func TrimLimits(
	s, sep string,
	trimCutter string,
	limits int,
) *[]string {
	splits := strings.SplitN(s, sep, limits)

	for i, currentItem := range splits {
		splits[i] = strings.Trim(currentItem, trimCutter)
	}

	return &splits
}
