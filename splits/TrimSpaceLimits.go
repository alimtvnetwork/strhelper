package splits

import "strings"

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
