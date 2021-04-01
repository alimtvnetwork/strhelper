package splits

import "strings"

func TrimSpace(s, sep string) *[]string {
	splits := strings.Split(s, sep)

	for i, currentItem := range splits {
		splits[i] = strings.TrimSpace(currentItem)
	}

	return &splits
}
