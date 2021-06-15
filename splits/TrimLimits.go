package splits

import "strings"

// TrimLimits first splits by the separator and then trim each line by trimCutter and up to the limit given
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
