package splits

import "strings"

// TrimSpaceLimits first splits by the separator and then trim each line by space upto the limit given.
func TrimSpaceLimits(
	s, sep string,
	limits int,
) *[]string {
	splits := strings.SplitN(
		s,
		sep,
		limits)

	for i, currentItem := range splits {
		splits[i] = strings.TrimSpace(currentItem)
	}

	return &splits
}
