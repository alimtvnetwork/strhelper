package strlines

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// GetBySpace split by ` ` space
func GetBySpace(content string) []string {
	allLines := strings.Split(content, constants.Space)

	return allLines
}
