package splits

import (
	"strings"
)

// TrimSpace first splits by the separator and then trim each line by space
func TrimSpace(s, sep string) *[]string {
	splits := strings.Split(s, sep)

	for i, currentItem := range splits {
		splits[i] = strings.TrimSpace(currentItem)
	}

	return &splits
}
